package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Sigdriv/feriehus/model"
	"github.com/gin-gonic/gin"
)

func (srv *Handler) HandleGetApartment(c *gin.Context) {
	log := srv.getLog(c)

	id := c.Param("id")

	bytes, err := os.ReadFile("./data/apartments.json")
	if err != nil {
		log.Errorf("Failed reading apartments data >> %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var apartments []model.Apartment
	err = json.Unmarshal(bytes, &apartments)
	if err != nil {
		log.Errorf("Failed parsing apartments data >> %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	for _, apartment := range apartments {
		if id == apartment.ID {
			c.JSON(http.StatusOK, apartment)
			return
		}
	}
}
