package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func postAlbums(c *gin.Context) {
    var newAlbum album
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context){
	id:= c.Param("id")

	for _,alb := range albums{
		if alb.ID == id{
			c.IndentedJSON(http.StatusOK, alb)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Album not found"})


}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
	router.POST("/albums",postAlbums)
	router.GET("/albums/:id", getAlbumByID)
    router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,albums)
}
