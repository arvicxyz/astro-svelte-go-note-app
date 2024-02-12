package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodos)
	router.POST("/todo", postTodo)
	router.PUT("/todo/:id", putTodo)
	router.PATCH("/todo/:id", patchTodo)
	router.DELETE("/todo/:id", deleteTodo)
	router.Run("localhost:8080")
}

type Todo struct {
	Id string `json:id`
	Todo string `json:todo`
	Completed bool `json:completed`
}

var todos = []Todo {
	{Id: "1", Todo: "Create an app", Completed: false},
	{Id: "2", Todo: "Finish template", Completed: false},
	{Id: "3", Todo: "Learn go", Completed: true},
	{Id: "4", Todo: "Learn svelte", Completed: false},
	{Id: "5", Todo: "Learn astro", Completed: false},
	{Id: "6", Todo: "Review HTML, CSS and JavaScript", Completed: false},
}

func getTodos(c *gin.Context) {
	fmt.Println("Get todos")
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodo(c *gin.Context) {
	fmt.Println("Add todo")
}

func putTodo(c *gin.Context) {
	fmt.Println("Edit todo")
}

func patchTodo(c *gin.Context) {
	fmt.Println("Update todo")
}

func deleteTodo(c *gin.Context) {
	fmt.Println("Delete todo")
}