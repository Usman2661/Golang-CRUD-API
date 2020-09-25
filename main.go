package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/usman2661/Todo-Crud-Go/controllers"
	"github.com/usman2661/Todo-Crud-Go/models"

	// "github.com/gin-contrib/cors"
	"net/http"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

// CORS Middleware
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
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

	router.Use(CORS)

	router.GET("/todos", controllers.GetTodos)          //  Get all the todos
	router.POST("/todos", controllers.CreateTodo)       // Create a new todo
	router.GET("/todos/:id", controllers.GetTodo)       // Fetch a single todo using id
	router.PUT("/todos/:id", controllers.UpdateTodo)    // Update and existing todo id has to be passed
	router.DELETE("/todos/:id", controllers.DeleteTodo) // Delete a todo using the id

	router.GET("/catagory/:username", controllers.GetCatagories) //  Get the Different catagories

	router.Run()
}
