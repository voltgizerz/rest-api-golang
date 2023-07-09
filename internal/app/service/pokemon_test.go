package service

import (
	"rest-api-golang/internal/app/entity"
	"rest-api-golang/internal/app/errors"
	"rest-api-golang/internal/app/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPokemonService_GetPokemon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPokeRepo := mocks.NewMockPokemonRepositoryInterface(ctrl)
	var pokemon *[]entity.Pokemon

	ps := NewPokemonService(mockPokeRepo)

	t.Run("Success", func(t *testing.T) {
		mockPokeRepo.EXPECT().GetPokemon("1", "Types").Return(pokemon, nil)

		poke, err := ps.GetPokemon("1")

		assert.Nil(t, err)
		assert.Nil(t, poke)
	})

	t.Run("Error", func(t *testing.T) {
		mockPokeRepo.EXPECT().GetPokemon("1", "Types").Return(pokemon, errors.ErrorOccur)

		poke, err := ps.GetPokemon("1")

		assert.NotNil(t, err)
		assert.Nil(t, poke)
	})
}

func TestPokemonService_GetAllPokemons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPokeRepo := mocks.NewMockPokemonRepositoryInterface(ctrl)
	var pokemon *[]entity.Pokemon

	ps := NewPokemonService(mockPokeRepo)

	t.Run("Success", func(t *testing.T) {
		mockPokeRepo.EXPECT().GetAllPokemons("Types").Return(pokemon, nil)

		poke, err := ps.GetAllPokemons()

		assert.Nil(t, err)
		assert.Nil(t, poke)
	})

	t.Run("Error", func(t *testing.T) {
		mockPokeRepo.EXPECT().GetAllPokemons("Types").Return(pokemon, errors.ErrorOccur)

		poke, err := ps.GetAllPokemons()

		assert.NotNil(t, err)
		assert.Nil(t, poke)
	})
}

func TestPokemonService_CreatePokemon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPokeRepo := mocks.NewMockPokemonRepositoryInterface(ctrl)
	pokemon := entity.Pokemon{}

	ps := NewPokemonService(mockPokeRepo)

	t.Run("Success", func(t *testing.T) {
		mockPokeRepo.EXPECT().CreatePokemon(&pokemon).Return(nil)

		err := ps.CreatePokemon(&pokemon)

		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		mockPokeRepo.EXPECT().CreatePokemon(&pokemon).Return(errors.ErrorOccur)

		err := ps.CreatePokemon(&pokemon)

		assert.NotNil(t, err)
	})
}

func TestPokemonService_DeletePokemon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPokeRepo := mocks.NewMockPokemonRepositoryInterface(ctrl)

	ps := NewPokemonService(mockPokeRepo)

	t.Run("Success", func(t *testing.T) {
		mockPokeRepo.EXPECT().DeletePokemon("1").Return(nil)

		err := ps.DeletePokemon("1")

		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		mockPokeRepo.EXPECT().DeletePokemon("1").Return(errors.ErrorOccur)

		err := ps.DeletePokemon("1")

		assert.NotNil(t, err)
	})
}

func TestPokemonService_UpdatePokemon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPokeRepo := mocks.NewMockPokemonRepositoryInterface(ctrl)
	pokemon := entity.Pokemon{}

	ps := NewPokemonService(mockPokeRepo)

	t.Run("Success", func(t *testing.T) {
		mockPokeRepo.EXPECT().UpdatePokemon(&pokemon, "1").Return(nil)

		err := ps.UpdatePokemon(&pokemon, "1")

		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		mockPokeRepo.EXPECT().UpdatePokemon(&pokemon, "1").Return(errors.ErrorOccur)

		err := ps.UpdatePokemon(&pokemon, "1")

		assert.NotNil(t, err)
	})
}
