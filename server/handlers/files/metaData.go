package files

import (
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCurrDirElements(c *gin.Context) {
	userIdAny, exist := c.Get("userId")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "userId not present",
		})
		return
	}

	userIdStr, ok := userIdAny.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid userId type",
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

	ids, err := stringsToUUID([]string{parentIdStr, userIdStr})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid uuid format",
		})
		return
	}

	parentId := &ids[0]
	ownerId := ids[1]

	children, err := database.GetAllChild(parentId, ownerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch directory contents",
		})
		return
	}

	c.JSON(http.StatusOK, children)
}

func stringsToUUID(ids []string) ([]uuid.UUID, error) {
	uuids := make([]uuid.UUID, 0, len(ids))

	for _, id := range ids {
		if id == "" {
			uuids = append(uuids, uuid.Nil)
			continue
		}

		parsed, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, parsed)
	}

	return uuids, nil
}
