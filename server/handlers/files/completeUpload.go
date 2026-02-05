package files

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/storage"
	"github.com/Ahmed-Armaan/FileNest/storage/helper"
	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CompleteUpload(db database.DatabaseStore, s storage.StorageStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileName := c.Query("name")
		objectKey := c.Query("objectKey")
		uploadId := c.Query("uploadId")

		size, err := strconv.ParseInt(c.Query("size"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": "Invalid size provided",
			})
			return
		}

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

		// googleId type assertion
		googleId, err := utils.GoogleIdstring(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		// get userId
		user, err := db.GetUserDataByGoogleId(googleId, database.UserDbColums.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
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

		err = s.CompleteUploadS3(objectKey, uploadId, completdPartsData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		// if file is uploaded to S3 but database insertion fails, delete the uploaded file
		if err := db.CreateNode(fileName, database.NodeTypeFile, &parentId, user.ID, &size, objectKey); err != nil {
			s.DeleteFileByObjectKey(&objectKey)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Database insert failed",
			})
			return
		}

		c.JSON(200, gin.H{
			"msg": "File upload successfull",
		})
	}
}
