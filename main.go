package main

import (
	"net/http"

	"io"

	"log"

	"api/resizer"

	"github.com/gin-gonic/gin"

	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/upload", HandleUpload)

	router.Run("localhost:8080")
}

// This function handles the upload of an image
func HandleUpload(c *gin.Context) {

	// Parses the file
	c.Request.ParseMultipartForm(20 << 30)

	// Retrieves the file
	file, err := c.FormFile("image")
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "couldn't read the file",
		})
		return
	}
	
	// Log of the file
	fmt.Printf("Uploaded file: %v\n", file.Filename)

	// Opens the file to geth the multipart file
	multiFile, err := file.Open()
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "The file is corrupted",
		})
		return
	}

	// Converts the multipart file to a []byte
	fileBytes, err := io.ReadAll(multiFile)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	// Logs the success
	fmt.Println("Successfully uploaded the file")

	// Calls the downscaler
	imgResized, err := resizer.Downscale(fileBytes)
	if err != nil {
		log.Fatalf(err.Error())
	}

	UploadImages(fileBytes, imgResized)

}

// This function uploads images to the cloud service
func UploadImages(image []byte, resized []byte) {

}

