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

// people slice to seed record person data.
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
	// insert request handlers
	// Listen at http://localhost:8080
	router.Run(":8080")
}

func getPeople(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, people)
}
