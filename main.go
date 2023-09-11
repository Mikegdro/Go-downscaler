package main

import (
	"net/http"
	"os"

	"io"

	"log"

	"api/resizer"

	"github.com/gin-gonic/gin"

	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/test", TestResizer)

	router.Run("localhost:8080")
}

func TestResizer(c *gin.Context) {

	c.Request.ParseMultipartForm(20 << 30)

	file, err := c.FormFile("image")
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "couldn't read the file",
		})
		return
	}
	
	fmt.Printf("Uploaded file: %v\n", file.Filename)
	
	tempFile, err := os.CreateTemp("/home/mike/Documentos/temp-images", "upload-*.png")
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "couldn't create a temp file",
		})
	}
	defer tempFile.Close()

	bytes, err := file.Open()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	fileBytes, err := io.ReadAll(bytes)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	tempFile.Write(fileBytes)

	fmt.Println("Successfully uploaded the file")

	resizer.Downscale(fileBytes)
}

