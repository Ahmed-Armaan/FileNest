package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/storage/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetNewUploadUrl(c *gin.Context) {
	fmt.Printf("\nGetting URL\n")
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

func CompleteUpload(c *gin.Context) {
	fileName := c.Query("name")
	objectKey := c.Query("objectKey")
	uploadId := c.Query("uploadId")

	if uploadId == "" || objectKey == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "uploadId or objectKey missing",
		})
		return
	}

	parentId, err := uuid.Parse(c.Query("parentId"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Invalid parentId provided",
		})
		return
	}

	userId, exist := c.Get("userId")
	if !exist {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Cant get userId",
		})
		return
	}

	userIdUUID, ok := userId.(uuid.UUID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cant read userId",
		})
		return
	}

	req := c.Request.Body
	defer c.Request.Body.Close()
	reqData, err := io.ReadAll(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	completdPartsData := make([]helper.CompetedPartsData, 0)
	if err := json.Unmarshal(reqData, &completdPartsData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err = helper.CompleteMultipartUpload(ctx, s3Client, bucketName, objectKey, uploadId, completdPartsData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to complete upload",
		})
		return
	}

	if err := database.InsertNode(fileName, database.NodeTypeFile, &parentId, userIdUUID, objectKey); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Database insert failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "File upload successfull",
	})
}
