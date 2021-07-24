package controllers

import (
	"rest-api-golang/models"

	"net/http"
  "github.com/gin-gonic/gin"
 )

//GetUsers ... Get all users
func GetPokemons(c *gin.Context) {
	var pokemon []models.Pokemon
	err := models.GetAllPokemons(&pokemon)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, gin.H{
			"data": pokemon, 
			"status": http.StatusOK,
		})
	}
 }