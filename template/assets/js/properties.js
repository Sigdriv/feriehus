fetch("/api/apartments")
  .then((res) => {
    if (!res.ok) throw new Error("HTTP error " + res.status);

    console.log("Response received:", res.json);

    return res.json();
  })
  .then((data) => {
    const propertiesWrapper = document.getElementById("properties");

    const { apartments = [] } = data;

    apartments.forEach((apartment, index) => {
      const { id, name, price, images = [], size, beds, baths } = apartment;

      const propertyEl = document.createElement("div");
      propertyEl.className = "col-xl-4 col-md-6";
      propertyEl.dataset.aos = "fade-up";
      propertyEl.dataset.aosDelay = String((index + 1) * 100);

      console.log("Apartment data:", apartment);

      propertyEl.innerHTML = `
        <div class="card">
          <img src="${`/images/${images[0]}.png`}" alt="${name}" class="img-fluid" />
          
          <div class="card-body">
            <span class="sale-rent">${price} NOK</span>
            <h3>
              <a href="property-single.html?id=${id}" class="stretched-link">${name}</a>
            </h3>

            <div
              class="card-content d-flex flex-column justify-content-center text-center"
            >
              <div class="row propery-info">
                <div class="col">Areal</div>
                <div class="col">Senger</div>
                <div class="col">Baderom</div>
              </div>
              <div class="row">
                <div class="col">${size}m<sup>2</sup></div>
                <div class="col">${beds}</div>
                <div class="col">${baths}</div>
              </div>
            </div>
          </div>
        </div>`;

      propertiesWrapper.appendChild(propertyEl);
    });

    if (window.AOS) AOS.refresh();
  })
  .catch((error) => console.error("Error fetching apartments data:", error));
