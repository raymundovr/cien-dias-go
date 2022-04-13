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

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return 
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumById(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	for i, album := range albums {
		if album.ID == id {
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, newAlbum)
			return
		}
	}
	c.Status(404)	
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.Status(200)
			return
		}
	}
	c.Status(404)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)
	router.PATCH("/albums/:id", updateAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.Run("localhost:8080")
}
