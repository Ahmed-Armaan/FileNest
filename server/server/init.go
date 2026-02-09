package server

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/handlers"
	"github.com/Ahmed-Armaan/FileNest/handlers/auth"
	"github.com/Ahmed-Armaan/FileNest/handlers/files"
	"github.com/Ahmed-Armaan/FileNest/handlers/middleware"
	"github.com/Ahmed-Armaan/FileNest/storage"
)

func Run(db database.DatabaseStore, s storage.StorageStore) error {
	r := gin.Default()

	// cors.DefaultConfig()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URI")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/auth/callback", auth.GetCredentials)
	r.GET("/test_user", handlers.Test_user)

	api := r.Group("/api")
	api.Use(middleware.VerifyJwt())
	api.GET("/me", handlers.Me(db)) // provide user their data

	// returns root node id and updated at
	api.GET("/root_node", files.GetRootDirId(db))

	// GET /get_elements?parentId=...
	// Returns children of a directory
	api.GET("/get_elements", files.GetCurrDirElements(db))

	// PUT /create_directory?parentId=...&dirName=...
	// Creates a new directory under parent
	api.PUT("/create_directory", files.CreateDirectory(db))

	// POST used to prevent caching of time-limited upload URLs
	api.POST("/get_upload_url", s.GetNewUploadUrl)

	// POST /get_upload_url/parts?uploadId=...&objectKey=...&partNumber=...
	// Returns a presigned URL for a single multipart upload part
	api.POST("/get_upload_url/parts", s.GetUploadUrl)

	// POST /complete_upload?name=...&objectKey=...&uploadId=...
	// Body: []CompletedPartsData
	api.POST("/complete_upload", files.CompleteUpload(db, s))

	api.POST("/share", files.ShareNode(db))

	api.GET("/all_shared", files.GetSharedFilesList(db))

	// This route is on the r group because share is free from user constraint;
	// Shared files can be accecssed by anyone irrespective of logon status. Thus jwt middleware is not used here
	r.POST("/get_share", files.GetSharedRootNode(db))

	r.GET("/shared_password_status", files.GetSharedPasswordStatus(db))

	// DELETE /deelete?nodeId=...
	// deleted a node in file tree
	api.DELETE("/delete", files.DeleteNode(db))

	// POST /start_download?fileId=...
	// Resolves fileId to objectKey, size and name
	api.POST("/start_download", s.DownloadInit(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := r.Run(":" + port)
	if err != nil {
		return err
	}
	return nil
}
