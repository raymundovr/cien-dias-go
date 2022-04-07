package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID 		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

var albums = []album{
	{ ID: "1", Title: "Siembra", Artist: "Willie Colón & Rubén Blades", Price: 29.99 },
	{ ID: "2", Title: "Princesa Donashii", Artist: "Princesa Donashii", Price: 25.99 },
	{ ID: "3", Title: "El Silencio", Artist: "Caifanes", Price: 24.99 },
	{ ID: "4", Title: "Latin-Rock-Soul", Artist: "Fania All Stars", Price: 29.99 },
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
