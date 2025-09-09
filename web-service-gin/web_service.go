package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Dookie", Artist: "Green Day", Price: 24.99},
	{ID: "2", Title: "Rocket to Russia", Artist: "Ramones", Price: 14.99},
	{ID: "3", Title: "So no forevis", Artist: "Raimundos", Price: 19.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
