package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usman2661/Todo-Crud-Go/models" // new
)

func main() {
	router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	models.ConnectDataBase()

	router.Run()
}
