package main

import (
	"net/http"

	"github.com/ae-lexs/vinyl_store/database"
	"github.com/ae-lexs/vinyl_store/model"
	"github.com/gin-gonic/gin"
)

var albums = []model.Album{
	{Title: "Revolver", Artist: "The Beatles", Price: 100},
	{Title: "Dummy", Artist: "Portished", Price: 200},
	{Title: "In Rainbows", Artist: "Radiohead", Price: 400},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	database.SetUpDatabase()

	for _, album := range albums {
		database.DBInstance.Create(&album)
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run(":8080")
}
