package main

import (
	"net/http"
	
	"api/resizer"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/test", TestResizer)

	router.Run("localhost:8080")
}

func TestResizer(c *gin.Context) {
	resizer.Downscale()
	c.JSON(http.StatusOK, gin.H{
		"Work": "Done",
	})
}