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
	"github.com/Ahmed-Armaan/FileNest/handlers/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error: Cant load env file\n%s", err)
	}

	if err := database.DbInit(); err != nil {
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

	r.GET("/auth/callback", auth.GetCredentials)
	r.GET("/me", middleware.VerifyJwt(), handlers.Me)

	err := r.Run()
	if err != nil {
		log.Fatalf("Error: Cant run server\n%s", err)
	}
}
