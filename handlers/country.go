package handlers

import (
	"net/http"

	"zones/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCountries(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetCountries())
}

func GetCountry(c *gin.Context) {
	countryId, err := primitive.ObjectIDFromHex(c.Query("countryId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, services.GetCountry(countryId))
}
