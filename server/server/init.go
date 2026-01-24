package server

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Ahmed-Armaan/FileNest/handlers"
	"github.com/Ahmed-Armaan/FileNest/handlers/auth"
	"github.com/Ahmed-Armaan/FileNest/handlers/files"
	"github.com/Ahmed-Armaan/FileNest/handlers/middleware"
	"github.com/Ahmed-Armaan/FileNest/storage"
)

func Run() error {
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

	api := r.Group("/api")
	api.Use(middleware.VerifyJwt())
	api.GET("/me", handlers.Me) // provide user their data

	// GET /get_elements?parentId=...
	// Returns children of a directory
	api.GET("/get_elements", files.GetCurrDirElements)

	// PUT /create_directory?parentId=...&dirName=...
	// Creates a new directory under parent
	api.PUT("/create_directory", files.CreateDirectory)

	// POST /complete_upload?name=...&objectKey=...&uploadId=...
	// Body: []CompletedPartsData
	api.POST("/complete_upload", storage.CompleteUpload)

	// POST used to prevent caching of time-limited upload URLs
	api.POST("/get_upload_url", storage.GetNewUploadUrl)

	// POST /get_upload_url/parts?uploadId=...&objectKey=...&partNumber=...
	// Returns a presigned URL for a single multipart upload part
	api.POST("/get_upload_url/parts", storage.GetUploadUrl)

	// POST /start_download?fileId=...
	// Resolves fileId to objectKey, size and name
	api.POST("/start_download", storage.DownloadInit)

	// POST /get_download_url/parts?objectKey=...&partNumber=...
	// Returns a presigned URL for downloading a file part
	//api.POST("/get_download_url/parts", storage.GetDownloadUrl)

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
