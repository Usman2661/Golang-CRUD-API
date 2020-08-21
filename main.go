package main

import (
	"fmt"
	"os"
	"github.com/usman2661/Todo-Crud-Go/controllers"
	"github.com/usman2661/Todo-Crud-Go/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	router := gin.Default()

	host := GoDotEnvVariable("host")
	port := GoDotEnvVariable("dbport")
	user := GoDotEnvVariable("user")
	dbname := GoDotEnvVariable("dbname")
	password := GoDotEnvVariable("password")
	sslmode := GoDotEnvVariable("sslmode")

	models.ConnectDataBase(host, port, user, dbname, password, sslmode)

	router.GET("/todos", controllers.GetTodos) //  Get all the todos
	router.POST("/todos", controllers.CreateTodo) // Create a new todo
	router.GET("/todos/:id", controllers.GetTodo) // Fetch a single todo using id
	router.PUT("/todos/:id", controllers.UpdateTodo) // Update and existing todo id has to be passed
	router.DELETE("/todos/:id", controllers.DeleteTodo) // Delete a todo using the id

	router.Run()
}
