package files

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShareReq struct {
	NodeId   string `json:"nodeId"`
	Password string `json:"password"`
}

type ShareFetchReq struct {
	Code     string `json:"code"`
	Password string `json:"password"`
}

func ShareNode(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request.Body
		defer c.Request.Body.Close()

		reqdata, err := io.ReadAll(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		shareReq := ShareReq{}
		if err := json.Unmarshal(reqdata, &shareReq); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		nodeIdUUID, err := uuid.Parse(shareReq.NodeId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid nodeId",
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

		code, err := db.ShareNode(nodeIdUUID, shareReq.Password, googleId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "database error",
			})
			return
		}

		c.JSON(200, gin.H{
			"code": code,
		})
	}
}

func GetSharedPasswordStatus(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		status, err := db.GetSharedPasswordStatus(code)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get the node's status",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": status,
		})
	}
}

func GetSharedRootNode(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "failed to read request body",
			})
			return
		}

		var fetchReq ShareFetchReq
		if err := json.Unmarshal(data, &fetchReq); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "failed to read request body",
			})
			return
		}

		sharedFiles, err := db.GetSharedNode(fetchReq.Code, fetchReq.Password)
		if err != nil {
			if err == database.ErrNoPasswordProvided || err == database.ErrWrongPassword {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to fetch requested file",
				})
			}
			return
		}

		c.JSON(200, sharedFiles)
	}
}

func GetSharedFilesList(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		googleId, err := utils.GoogleIdstring(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to resolve user",
			})
			return
		}

		var sharedNodes []database.SharedNode

		sharedNodes, err = db.GetAllSharedNodes(googleId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "database error",
			})
			return
		}

		c.JSON(http.StatusOK, sharedNodes)
	}
}

func StopSharing(db database.DatabaseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		googleId, err := utils.GoogleIdstring(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "no googleId provided",
			})
			return
		}
		nodeId := c.Query("nodeId")

		if err := db.RemoveSharedNode(googleId, nodeId); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "database error",
			})
			return
		}

		c.Status(200)
	}
}
