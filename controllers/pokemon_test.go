package controllers

import (
	"net/http/httptest"
	"rest-api-golang/entity"
	"rest-api-golang/errors"
	"rest-api-golang/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestPokemonController_GetPokemons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPokeService := mocks.NewMockPokemonServiceInterface(ctrl)
	var pokemon *[]entity.Pokemon

	pc := NewPokemonController(mockPokeService)

	t.Run("Success GetAllPokemons 200 OK", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		mockPokeService.EXPECT().GetAllPokemons().Return(pokemon, nil)

		pc.GetPokemons(c)

		assert.Equal(t, w.Code, 200)
	})

	t.Run("Success GetPokemon 200 OK", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Params = []gin.Param{{Key: "id", Value: "123"}}

		mockPokeService.EXPECT().GetPokemon("123").Return(pokemon, nil)

		pc.GetPokemons(c)

		assert.Equal(t, w.Code, 200)
	})

	t.Run("Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Params = []gin.Param{{Key: "id", Value: "123"}}

		mockPokeService.EXPECT().GetPokemon("123").Return(nil, errors.ErrorOccur)

		pc.GetPokemons(c)

		assert.Equal(t, w.Code, 404)
	})
}
