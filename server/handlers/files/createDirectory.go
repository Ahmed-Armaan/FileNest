package files

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateDirectory(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		dirName := c.Query("dirName")
		parentIdstr := c.Query("parentId")
		parentIdUUID, err := uuid.Parse(parentIdstr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid parentId",
			})
			return
		}

		googleId, err := utils.GoogleIdstring(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		user, err := db.GetUserDataByGoogleId(googleId, database.UserDbColums.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		if err := db.CreateNode(dirName, database.NodeTypeDirectory, &parentIdUUID, user.ID, nil); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
			})
			return
		}
	}
}
