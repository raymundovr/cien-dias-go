package main

import (
	"vasquezruiz.com/conversions"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

func getFarenheit(c *gin.Context) {
	value := c.Param("value")
	celsius, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Errrrr", err)
		c.String(400, "Incorrect value")
		return 
	}

	farenheit := conversions.IntoFarenheit(conversions.Celsius(celsius))
	c.String(200, "Value %s", farenheit)
}

func main() {
	server := gin.Default()
	server.GET("/farenheit/:value", getFarenheit)

	server.Run("localhost:8081")
}