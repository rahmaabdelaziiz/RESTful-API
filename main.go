package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var users = []user{
	{ID: "1", FirstName: "Rahma", LastName: "Abdelaziz"},
	{ID: "2", FirstName: "Rahmaa", LastName: "Abdelaziiz"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
}
