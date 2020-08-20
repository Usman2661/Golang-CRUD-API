package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/usman2661/Todo-Crud-Go/models"
	// new
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	router := gin.Default()

	host := goDotEnvVariable("host")
	port := goDotEnvVariable("dbport")
	user := goDotEnvVariable("user")
	dbname := goDotEnvVariable("dbname")
	password := goDotEnvVariable("password")
	sslmode := goDotEnvVariable("sslmode")

	models.ConnectDataBase(host, port, user, dbname, password, sslmode)

	router.Run()
}
