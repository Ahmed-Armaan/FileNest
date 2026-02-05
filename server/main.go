package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/database/cleanupjobs"
	"github.com/Ahmed-Armaan/FileNest/server"
	"github.com/Ahmed-Armaan/FileNest/storage"
	"github.com/Ahmed-Armaan/FileNest/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env found, relying on enviornment variables") // in production .env is not present
	}

	databaseStore, err := database.DbInit()
	if err != nil {
		log.Fatalln(err)
	}

	storageStore, err := storage.S3Init()
	if err != nil {
		log.Fatalln(err)
	}

	if err := utils.JWTinit(); err != nil {
		log.Fatalln(err)
	}

	cleanupjobs.CronInit(databaseStore, storageStore)

	if err := server.Run(databaseStore, storageStore); err != nil {
		log.Fatalln(err)
	}
}
