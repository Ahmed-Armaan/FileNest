package storage

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/storage/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *StorageHolder) DownloadInit(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileIdstr := c.Query("fileId")
		fileId, err := uuid.Parse(fileIdstr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid fileId",
			})
			return
		}

		downloadMetaData, err := db.GetNodeObjectInfo(fileId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get objectKey",
			})
			return
		}

		url, err := helper.GetPresignDownloadUrl(s.ctx, s.s3Client, s.bucketName, *downloadMetaData.ObjectKey, *downloadMetaData.SizeBytes)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get download url",
			})
			return
		}

		c.JSON(200, gin.H{
			"objectKey": downloadMetaData.ObjectKey,
			"size":      downloadMetaData.SizeBytes,
			"name":      downloadMetaData.Name,
			"url":       url,
		})
	}
}
