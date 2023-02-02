package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"iD"`
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

func main() {

	server := gin.Default()
	server.GET("/todos", getTodosData)
	server.Run("localhost:9090")
}
