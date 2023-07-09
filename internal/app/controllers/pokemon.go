package controllers

import (
	"encoding/json"
	"io/ioutil"
	"rest-api-golang/internal/app/entity"
	"rest-api-golang/internal/app/service"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

// PokemonController - ...
type PokemonController struct {
	PokemonService service.PokemonServiceInterface
}

// NewPokemonController - ...
func NewPokemonController(servicePokemon service.PokemonServiceInterface) PokemonController {
	return PokemonController{
		PokemonService: servicePokemon,
	}
}

// GetPokemons - Get all pokemons
func (p PokemonController) GetPokemons(c *gin.Context) {
	id := c.Param("id")
	var err error

	var pokemon *[]entity.Pokemon

	if id == "" {
		pokemon, err = p.PokemonService.GetAllPokemons()
	} else {
		pokemon, err = p.PokemonService.GetPokemon(id)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": pokemon,
	})
}

// PostPokemons - Create new pokemon
func (p PokemonController) PostPokemons(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var pokemon entity.Pokemon

	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "request body not valid",
			"data":    pokemon,
		})
	}

	result := p.PokemonService.CreatePokemon(&pokemon)
	if result != nil {
		// Failed insert pokemon to database
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "failed insert data to database",
			"data":    pokemon,
		})
	}
	// success created new pokemon
	c.JSON(http.StatusCreated, gin.H{
		"code":       http.StatusCreated,
		"message":    "created",
		"created_at": time.Now().Format(time.RFC850),
		"data":       pokemon,
	})
}

// PatchPokemons - Create new pokemon
func (p PokemonController) PatchPokemons(c *gin.Context) {
	id := c.Param("id")
	body, _ := ioutil.ReadAll(c.Request.Body)
	var pokemon entity.Pokemon

	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "request body not valid",
			"data":    pokemon,
		})
	}

	result := p.PokemonService.UpdatePokemon(&pokemon, id)
	if result != nil {
		// Failed update pokemon to database
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "failed update data to database",
			"data":    pokemon,
		})
	}
	// success updated pokemon
	c.JSON(http.StatusCreated, gin.H{
		"code":       http.StatusCreated,
		"message":    "updated",
		"updated_at": time.Now().Format(time.RFC850),
		"data":       pokemon,
	})
}

// DeletePokemon delete pokemon
func (p PokemonController) DeletePokemon(c *gin.Context) {
	id := c.Param("id")
	result := p.PokemonService.DeletePokemon(id)
	if result != nil {
		// Failed insert pokemon to database
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "failed insert data to database",
		})
	}
	// success deleted new pokemon
	c.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"message":    "deleted",
		"deleted_at": time.Now().Format(time.RFC850),
	})
}
