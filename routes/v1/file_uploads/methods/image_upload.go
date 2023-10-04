package file_uploads

import (
	"fmt"
	common_utils "go-gin-api-boilerplate/utils"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// SendImageRequest represents the request for image upload.
type SendImageRequest struct {
	StorageType string                `json:"storageType"` // "aws" or "firebase"
	File        *multipart.FileHeader `json:"file"`
}

// SendImageResponse represents the response for image upload.
type SendImageResponse struct {
	URL string `json:"url"`
}

// @Summary Upload an image
// @Tags Image Upload
// @Description Upload an image to either AWS S3 or Firebase Storage
// @Accept multipart/form-data
// @Produce json
// @Param storageType formData string true "Storage Type ('aws' or 'firebase')"
// @Param file formData file true "Image File"
// @Success 200 {object} SendImageResponse
// @Router /api/v1/files/upload-image [post]
func UploadImage(c *gin.Context) {
	var request SendImageRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storageType := c.DefaultPostForm("storageType", "aws")
	// Generate a unique object key (filename)
	objectKey := "1696420660094_k0tjRe.png"
	// Upload the image based on the specified storage type
	var publicURL string
	var err error

	// Retrieve the file parameter
	file, err := c.FormFile("file")
	if err != nil {
		// Handle the error if the file is not included in the request
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	if storageType == "aws" {
		publicURL, err = common_utils.UploadToS3(file, os.Getenv("AWS_S3_BUCKET"), objectKey)
		if err != nil {
			log.Printf("Error uploading to AWS S3: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	} else if storageType == "firebase" {
		// You can call the UploadToFirebase function like this:
		publicURL, err := common_utils.UploadToFirebase(common_utils.FirebaseConnection, file, objectKey)
		fmt.Println(publicURL)
		if err != nil {
			log.Printf("Error uploading to Firebase: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	} else {
		log.Printf("Invalid storage type: %s", storageType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid storage type"})
		return
	}

	response := SendImageResponse{URL: publicURL}
	c.JSON(http.StatusOK, response)
}
