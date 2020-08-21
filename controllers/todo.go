package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usman2661/Todo-Crud-Go/models"
)

//GET all todos
func GetTodos(c *gin.Context) {

	var todos []models.Todo
	models.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}
