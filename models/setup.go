package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//Instatiate the GORM DB
var DB *gorm.DB

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

//Connect to the Databse
func ConnectDataBase() {

	host := goDotEnvVariable("host")
	port := goDotEnvVariable("port")
	user := goDotEnvVariable("user")
	dbname := goDotEnvVariable("dbname")
	password := goDotEnvVariable("password")
	sslmode := goDotEnvVariable("sslmode")

	//Connect to the Databse
	database, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password+" sslmode="+sslmode)

	fmt.Println(err)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Todo{})

	DB = database
}
