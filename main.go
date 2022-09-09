package main

import (
	"zones/env"
	"zones/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET(baseUrl+"/countries", handlers.GetCountries)
	r.GET(baseUrl+"/countryById", handlers.GetCountry)

	r.GET(baseUrl+"/regions", handlers.GetRegions)
	r.GET(baseUrl+"/regionById", handlers.GetRegion)

	r.Run(env.GetInstance().Port)
}
