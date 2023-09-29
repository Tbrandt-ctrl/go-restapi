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
	 {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
    router.GET("/albums/:id", getAlbumByID)


	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context){
	    c.IndentedJSON(http.StatusOK, albums)

}

func postAlbum(c *gin.Context){

	var newAlbum album


	// Call BindJSON to bind the received JSON to
    // newAlbum.
	if err:= c.BindJSON(&newAlbum); err!=nil{
		return 
	}

	albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum) //Adding 201

}

func getAlbumByID(c * gin.Context){

	id:= c.Param("id")

	for _, a:= range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}