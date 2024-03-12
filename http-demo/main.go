package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Year     int    `json:"year"`
	Duration int    `json:"duration"`
}

var albums = []album{
	{
		ID:       "1",
		Title:    "Thriller",
		Artist:   "Michael Jackson",
		Year:     1982,
		Duration: 42,
	},
	{
		ID:       "2",
		Title:    "Back in Black",
		Artist:   "AC/DC",
		Year:     1980,
		Duration: 42,
	},
	{
		ID:       "3",
		Title:    "Revolver",
		Artist:   "The Beatles",
		Year:     1966,
		Duration: 42,
	},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
