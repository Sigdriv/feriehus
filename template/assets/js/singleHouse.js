const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get("id");

fetch(`/api/apartments/${id}`)
  .then((res) => {
    if (!res.ok) throw new Error("HTTP error " + res.status);

    return res.json();
  })
  .then((data) => {
    const {
      amenities = [],
      baths,
      beds,
      description,
      images = [],
      location,
      maps,
      name,
      shortDescription,
      size,
      type,
    } = data;

    const slidesWrapper = document.getElementById("apartment-slides");
    images.forEach((src) => {
      const slide = document.createElement("div");
      slide.className = "swiper-slide";

      const img = document.createElement("img");
      img.src = `/images/${src}.png`;
      img.alt = data.name;
      slide.appendChild(img);
      slidesWrapper.appendChild(slide);
    });

    const swiperEl = document.querySelector(".portfolio-details-slider");
    if (swiperEl && swiperEl.swiper) swiperEl.swiper.update();

    document.title = `${data.name} - Feriehus`;
    replaceTextContentById("apartment-name", name);
    replaceTextContentById("apartment-description-short", shortDescription);

    replaceTextContentById("apartment-price", `${data.price} NOK`);
    replaceTextContentById("apartment-id", id);
    replaceTextContentById("apartment-location", location);
    replaceTextContentById("apartment-type", type);
    replaceTextContentById("apartment-size", size);
    replaceTextContentById("apartment-beds", beds);
    replaceTextContentById("apartment-baths", baths);
    replaceTextContentById("apartment-amenities", amenities.join(", "));

    replaceTextContentById("apartment-description", description);
    document.getElementById("apartment-location-map").src =
      `https://www.google.com/maps/embed?pb=${maps}`;
    document.getElementById("apartment-floor-plan").src =
      `/images/${data.floorPlan}.png`;
  })
  .catch((error) => console.error("Error fetching apartment data:", error));

function replaceTextContentById(id, text) {
  const element = document.getElementById(id);

  if (element) element.textContent = text;
}
