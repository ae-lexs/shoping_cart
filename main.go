package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Album represents data about a record album.
type Album struct {
	gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{Title: "Revolver", Artist: "The Beatles", Price: 100},
	{Title: "Dummy", Artist: "Portished", Price: 200},
	{Title: "In Rainbows", Artist: "Radiohead", Price: 400},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the database.")
	}

	db.AutoMigrate(&Album{})

	for _, album := range albums {
		db.Create(album)
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
