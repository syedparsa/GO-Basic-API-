package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed string `json:"completed"`
}

var todos = []todo{

	{ID: "1", Item: "Complete your Assignment", Completed: "false"},
	{ID: "2", Item: "Finsish your Readings", Completed: "false"},
	{ID: "3", Item: "Talk to your family ", Completed: "false"},
}

// this context has bunch of information about the  incoming http request
func getTodosData(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)
}

func PostTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, todos)
}

func getTodo(context *gin.Context) {

	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {

		context.Copy().IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}
func getTodoByID(id string) (*todo, error) {
	for i, T := range todos {

		if T.ID == id {

			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {

	server := gin.Default()
	server.GET("/todos", getTodosData)
	server.GET("/todos/:id", getTodosData)
	server.POST("/todos", PostTodo)
	server.Run("localhost:9090")
}
