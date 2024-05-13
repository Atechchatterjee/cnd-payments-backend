package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// data definitions
type person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type pingResponse struct {
	Ping string
}

var people = []person{
	{
		ID:   "1",
		Name: "ABC",
	},
	{
		ID:   "2",
		Name: "DEF",
	},
	{
		ID:   "3",
		Name: "GHI",
	},
}

func main() {
	// a gin router to handle requests
	var router *gin.Engine = gin.Default()
	router.GET("/people", getPeople)
	router.GET("/ping", ping)
	// insert request handlers
	// Listen at http://localhost:8080
	router.Run(":8080")
}

// a route used by the cron-schedule node.js server to ping the server
func ping(context *gin.Context) {
	var pingMsg = pingResponse{
		Ping: "success",
	}
	context.IndentedJSON(http.StatusOK, pingMsg)
}

func getPeople(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, people)
}
