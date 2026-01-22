package files

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCurrDirElements(c *gin.Context) {
	userId, exist := c.Get("userId")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "userId not present",
		})
		return
	}
	ownerId, ok := userId.(uuid.UUID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get userId",
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

	children, err := database.GetAllChild(&parentIdUUID, ownerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch directory contents",
		})
		return
	}

	c.JSON(http.StatusOK, children)
}

//func stringsToUUID(ids []string) ([]uuid.UUID, error) {
//	uuids := make([]uuid.UUID, 0, len(ids))
//
//	for _, id := range ids {
//		if id == "" {
//			uuids = append(uuids, uuid.Nil)
//			continue
//		}
//
//		parsed, err := uuid.Parse(id)
//		if err != nil {
//			return nil, err
//		}
//		uuids = append(uuids, parsed)
//	}
//
//	return uuids, nil
//}
