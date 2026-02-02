package storage

import (
	"errors"
	"strconv"

	"github.com/Ahmed-Armaan/FileNest/storage/helper"
	"github.com/gin-gonic/gin"
)

func (s *StorageHolder) GetNewUploadUrl(c *gin.Context) {
	uploadId, objectKey, err := helper.CreateNewUpload(s.ctx, s.s3Client, s.bucketName, "data/")
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"uploadId":  uploadId,
		"objectKey": objectKey,
	})
}

func (s *StorageHolder) GetUploadUrl(c *gin.Context) {
	uploadId := c.Query("uploadId")
	objectKey := c.Query("objectKey")
	partStr := c.Query("partNumber")

	partNumber, err := strconv.Atoi(partStr)
	if err != nil || partNumber < 1 {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid partNumber"})
		return
	}

	url, err := helper.PresignPart(s.ctx, s.s3Client, s.bucketName, objectKey, uploadId, int32(partNumber))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"url": url,
	})
}

func (s *StorageHolder) CompleteUploadS3(objectKey string, uploadId string, completdPartsData []helper.CompetedPartsData) error {
	err := helper.CompleteMultipartUpload(s.ctx, s.s3Client, s.bucketName, objectKey, uploadId, completdPartsData)
	if err != nil {
		return errors.New("Failed to complete upload")
	}
	return nil
}

//func (s *StorageHolder) CompleteUpload(c *gin.Context) {
//	fileName := c.Query("name")
//	objectKey := c.Query("objectKey")
//	uploadId := c.Query("uploadId")
//
//	size, err := strconv.ParseInt(c.Query("size"), 10, 64)
//	if err != nil {
//		c.AbortWithStatusJSON(500, gin.H{
//			"error": "Invalid size provided",
//		})
//		return
//	}
//
//	if uploadId == "" || objectKey == "" {
//		c.AbortWithStatusJSON(400, gin.H{
//			"error": "uploadId or objectKey missing",
//		})
//		return
//	}
//
//	parentId, err := uuid.Parse(c.Query("parentId"))
//	if err != nil {
//		c.AbortWithStatusJSON(500, gin.H{
//			"error": "Invalid parentId provided",
//		})
//		return
//	}
//
//	// googleId type assertion
//	googleId, err := utils.GoogleIdstring(c)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//			"error": err,
//		})
//		return
//	}
//
//	// get userId
//	user, err := db.GetUserDataByGoogleId(googleId, database.UserDbColums.ID)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//			"error": err,
//		})
//		return
//	}
//
//	req := c.Request.Body
//	defer c.Request.Body.Close()
//	reqData, err := io.ReadAll(req)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//			"error": "Failed to read request body",
//		})
//		return
//	}
//
//	completdPartsData := make([]helper.CompetedPartsData, 0)
//	if err := json.Unmarshal(reqData, &completdPartsData); err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
//			"error": "Invalid request body",
//		})
//		return
//	}
//
//	err = helper.CompleteMultipartUpload(s.ctx, s.s3Client, s.bucketName, objectKey, uploadId, completdPartsData)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//			"error": "Failed to complete upload",
//		})
//		return
//	}
//
//	if err := db.CreateNode(fileName, database.NodeTypeFile, &parentId, user.ID, &size, objectKey); err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//			"error": "Database insert failed",
//		})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"msg": "File upload successfull",
//	})
//}
