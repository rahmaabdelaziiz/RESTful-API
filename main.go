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

func postUsers(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUsers)
	router.Run("localhost:8080")
}
