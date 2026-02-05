package files

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteNode(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		googleId, err := utils.GoogleIdstring(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		nodeIdStr, exist := c.GetQuery("nodeId")
		if !exist {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "parentId not sent",
			})
			return
		}
		nodeIdUUID, err := uuid.Parse(nodeIdStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Could not read prentId",
			})
			return
		}

		if err := db.MarkNodeDeleted(googleId, nodeIdUUID); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "database error",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "deleted",
		})
	}
}
