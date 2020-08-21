package models

import (
	"fmt" // new

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Instatiate the GORM DB
var DB *gorm.DB

//Connect to the Databse
func ConnectDataBase(host string, port string, user string, dbname string, password string, sslmode string) {


	//Connect to the Databse
	database, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password+" sslmode="+sslmode)

	fmt.Println(err)

	if err != nil {
		panic("Failed to connect to database!")
	}
	if err == nil {
		fmt.Println("Connected to the Database!!!")
	}

	database.AutoMigrate(&Todo{})

	DB = database
}
