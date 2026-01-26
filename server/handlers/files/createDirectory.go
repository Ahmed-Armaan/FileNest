package files

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/utils"
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

	user, err := utils.GetUserFromGoogleId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if err := database.InsertNode(dirName, database.NodeTypeDirectory, &parentIdUUID, user.ID, nil); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}
}
