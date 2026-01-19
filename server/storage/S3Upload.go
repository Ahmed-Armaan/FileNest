package storage

import (
	"strconv"

	"github.com/Ahmed-Armaan/FileNest/storage/helper"
	"github.com/gin-gonic/gin"
)

func GetNewUploadUrl(c *gin.Context) {
	uploadId, objectKey, err := helper.CreateNewUpload(ctx, s3Client, bucketName)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"uploadId":  uploadId,
		"objectKey": objectKey,
	})
}

func GetUploadUrl(c *gin.Context) {
	uploadId := c.Query("uploadId")
	objectKey := c.Query("objectKey")
	partStr := c.Query("partNumber")

	partNumber, err := strconv.Atoi(partStr)
	if err != nil || partNumber < 1 {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid partNumber"})
		return
	}

	url, err := helper.PresignPart(ctx, s3Client, bucketName, objectKey, uploadId, int32(partNumber))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"url": url,
	})
}
