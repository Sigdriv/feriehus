package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Sigdriv/feriehus/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var cacheApartments = make(map[string]model.Apartment)

func (srv *Handler) InitApartments() {
	bytes, err := os.ReadFile("./data/apartments.json")
	if err != nil {
		logrus.Fatalf("Failed reading apartments data >> %v", err)
	}

	var apartments []model.Apartment
	err = json.Unmarshal(bytes, &apartments)
	if err != nil {
		logrus.Fatalf("Failed parsing apartments data >> %v", err)
	}

	for _, apartment := range apartments {
		cacheApartments[apartment.ID] = apartment
	}
}

func (srv *Handler) HandleGetApartment(c *gin.Context) {
	log := srv.getLog(c)

	id := c.Param("id")
	if id == "" {
		log.Error("Apartment ID is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Apartment ID is required"})
		return
	}

	apartment := cacheApartments[id]

	c.JSON(http.StatusOK, apartment)
}

func (srv *Handler) HandleGetApartments(c *gin.Context) {
	log := srv.getLog(c)

	if len(cacheApartments) == 0 {
		log.Warn("No apartments found")
		c.JSON(http.StatusNotFound, gin.H{"error": "No apartments found"})
		return
	}

	var apartments = make([]model.Apartments, 0, len(cacheApartments))
	for _, apartment := range cacheApartments {
		apartments = append(apartments, model.Apartments{
			ID:     apartment.ID,
			Name:   apartment.Name,
			Price:  apartment.Price,
			Images: apartment.Images,
			Size:   apartment.Size,
			Beds:   apartment.Beds,
			Baths:  apartment.Baths,
		})
	}

	c.JSON(http.StatusOK, gin.H{"apartments": apartments})
}
