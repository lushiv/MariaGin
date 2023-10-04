package file_uploads

import (
	file_uploads_methods "go-gin-api-boilerplate/routes/v1/file_uploads/methods"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/upload-image", file_uploads_methods.UploadImage)
}
