const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get("id");

fetch(`/api/apartments/${id}`)
  .then((res) => {
    if (!res.ok) throw new Error("HTTP error " + res.status);

    return res.json();
  })
  .then((data) => {
    const apartment = data.apartment;
    document.getElementById("apartment-name").textContent = apartment.name;
    document.getElementById("apartment-description").textContent =
      apartment.description;
    document.getElementById("apartment-price").textContent =
      `Price: ${apartment.price} NOK per night`;
  })
  .catch((error) => console.error("Error fetching apartment data:", error));
