package main

import (
	"vasquezruiz.com/conversions"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

func parseValue(c *gin.Context) (float64, error) {
	value := c.Param("value")
	return strconv.ParseFloat(value, 64)
}

func getFarenheit(c *gin.Context) {
	celsius, err := parseValue(c)
	if err != nil {
		fmt.Println("Errrrr", err)
		c.String(400, "Incorrect value")
		return 
	}

	farenheit := conversions.IntoFarenheit(conversions.Celsius(celsius))
	c.String(200, "Value %s", farenheit)
}

func getCelsius(c *gin.Context) {
	farenheit, err := parseValue(c)
	if err != nil {
		fmt.Println("Errrrr", err)
		c.String(400, "Incorrect value")
		return 
	}

	celsius := conversions.IntoCelsius(conversions.Farenheit(farenheit))
	c.String(200, "Value %s", celsius)
}

func main() {
	server := gin.Default()
	server.GET("/farenheit/:value", getFarenheit)
	server.GET("/celsius/:value", getCelsius)
	server.Run("localhost:8081")
}