package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/server"
	"github.com/Ahmed-Armaan/FileNest/storage"
	"github.com/Ahmed-Armaan/FileNest/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env found, relying on enviornment variables") // in production .env is not present
	}

	if err := database.DbInit(); err != nil {
		log.Fatalln(err)
	}

	if err := storage.S3Init(); err != nil {
		log.Fatalln(err)
	}

	if err := utils.JWTinit(); err != nil {
		log.Fatalln(err)
	}

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
