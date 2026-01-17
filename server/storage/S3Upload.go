package storage

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/storage/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type test struct {
	Name      string
	Age       int
	Something uuid.UUID
}

func GetUploadUrl(c *gin.Context) {
	objectKey := "data/" + uuid.NewString()

	uploadUrl, err := helper.GetPresignedUrl(http.MethodPut, objectKey, s3Client, ctx, bucketName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "S3 url fetch failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"uploadUrl": uploadUrl,
		"objectKey": objectKey,
	})
}
