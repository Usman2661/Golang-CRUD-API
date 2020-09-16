package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usman2661/Todo-Crud-Go/models"
)

type CreateTodoInput struct {
	Title     string `json:"title" binding="required"`
	Completed bool   `json:"completed"`
	Catagory  string `json:"catagory" binding="required"`
	UserName  string `json:"username" binding="required"`
}

type UpdateTodoInput struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Catagory  string `json:"catagory"`
	UserName  string `json:"username"`
}

//GET all todos
func GetTodos(c *gin.Context) {

	var todos []models.Todo
	models.DB.Find(&todos)

	c.JSON(http.StatusOK, todos)
}

//
func CreateTodo(c *gin.Context) {

	// Validate input
	var todoInput CreateTodoInput
	if err := c.ShouldBindJSON(&todoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	todo := models.Todo{Title: todoInput.Title, Completed: todoInput.Completed, Catagory: todoInput.Catagory, UserName: todoInput.UserName}
	models.DB.Create(&todo)

	c.JSON(http.StatusCreated, todo)
}

func GetTodo(c *gin.Context) {

	var todo models.Todo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found!"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Update the todo
func UpdateTodo(c *gin.Context) {

	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found!"})
		return
	}

	// Validate input
	var todoInput UpdateTodoInput
	if err := c.ShouldBindJSON(&todoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Updating with the Struct does not update boolean true back to false hence using string interface to update

	// models.DB.Model(&todo).Updates(todoInput)

	models.DB.Model(&todo).Updates(map[string]interface{}{
		"Title":     todoInput.Title,
		"Completed": todoInput.Completed,
		"Catagory":  todoInput.Catagory,
		"UserName":  todoInput.UserName,
	})

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found!"})
		return
	}

	models.DB.Delete(&todo)

	c.JSON(http.StatusOK, todo)
}
