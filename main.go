package main

import (
	"github.com/gin-gonic/gin"
	"rest-api-golang/database"
	"rest-api-golang/models"
)

func main() {
	db := database.Database()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/members", func(c *gin.Context) {
		var members []models.Member
		dataMembers :=  db.Find(&members)
		c.JSON(200, gin.H {
			"data": dataMembers,
		})
	})


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")


}