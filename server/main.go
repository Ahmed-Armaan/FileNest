package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Ahmed-Armaan/FileNest/handlers/auth"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error: Cant load env file\n%s", err)
	}

	runServer()
}

func runServer() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/auth/callback", auth.GetCredentials)

	err := r.Run()
	if err != nil {
		log.Fatalf("Error: Cant run server\n%s", err)
	}
}
