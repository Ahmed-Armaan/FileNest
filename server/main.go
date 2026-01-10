package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/handlers"
	"github.com/Ahmed-Armaan/FileNest/handlers/auth"
	"github.com/Ahmed-Armaan/FileNest/handlers/files"
	"github.com/Ahmed-Armaan/FileNest/handlers/middleware"
	"github.com/Ahmed-Armaan/FileNest/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error: Cant load env file\n%s", err)
	}

	if err := database.DbInit(); err != nil {
		log.Fatalln(err)
	}
	if err := storage.S3Init(); err != nil {
		log.Fatalln(err)
	}
	runServer()
}

func runServer() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URI")},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/auth/callback", auth.GetCredentials)

	api := r.Group("/api")
	api.Use(middleware.VerifyJwt())
	api.GET("/me", handlers.Me)
	api.GET("/get_elements", files.GetCurrDirElements)
	api.GET("/upload", storage.GetUploadUrl)

	err := r.Run()
	if err != nil {
		log.Fatalf("Error: Cant run server\n%s", err)
	}
}
