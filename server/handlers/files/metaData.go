package files

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRootDirId(c *gin.Context) {
	user, err := utils.GetUserFromGoogleId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	rootNode, err := database.GetRootNodeId(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "database Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"rootNodeId":        rootNode.ID,
		"rootNodeUpdatedAt": rootNode.UpdatedAt,
	})
}

func GetCurrDirElements(c *gin.Context) {
	user, err := utils.GetUserFromGoogleId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	parentIdStr, exist := c.GetQuery("parentId")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "parentId not sent",
		})
		return
	}
	parentIdUUID, err := uuid.Parse(parentIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Could not read prentId",
		})
		return
	}

	children, err := database.GetAllChild(&parentIdUUID, user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch directory contents",
		})
		return
	}

	c.JSON(http.StatusOK, children)
}
