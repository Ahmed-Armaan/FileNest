package files

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateDirectory(c *gin.Context) {
	dirName := c.Query("dirName")
	parentIdstr := c.Query("parentId")
	parentIdUUID, err := uuid.Parse(parentIdstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid parentId",
		})
		return
	}

	userId, exist := c.Get("userId")
	if !exist {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get userId",
		})
		return
	}
	userIdUUID, ok := userId.(uuid.UUID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get userId",
		})
		return
	}

	if err := database.InsertNode(dirName, database.NodeTypeDirectory, &parentIdUUID, userIdUUID, nil); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}
}
