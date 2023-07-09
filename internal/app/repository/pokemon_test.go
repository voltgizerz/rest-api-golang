package repository

import (
	"regexp"
	"rest-api-golang/configs"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockDBConn(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	var mockORM *gorm.DB

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if mock == nil {
		t.Error("sqlmock is null")
	}

	mockORM, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return mock, mockORM
}

func TestPokemonRepo_GetAllPokemons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Parallel()

	mock, mockORM := mockDBConn(t)
	dbConfig := config.Database{
		ORM: mockORM,
	}

	pr := NewPokemonRepository(dbConfig)

	pokemonMockRows := sqlmock.NewRows([]string{"id", "name", "height", "weight", "base_experience", "type"}).
		AddRow(nil, "bulbasaur", "123", "123", "123", nil)

	t.Run("Success return pokemons rows", func(t *testing.T) {
		mock.MatchExpectationsInOrder(false)
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `pokemons`")).
			WillReturnRows(pokemonMockRows)
		mock.ExpectCommit()

		pokemons, err := pr.GetAllPokemons()

		assert.Nil(t, err)
		assert.NotNil(t, pokemons)
		assert.Equal(t, (*pokemons)[0].Name, "bulbasaur")
	})

	t.Run("Error", func(t *testing.T) {
		mock.MatchExpectationsInOrder(false)
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `digimons`")).
			WillReturnRows(pokemonMockRows)
		mock.ExpectCommit()

		pokemons, err := pr.GetAllPokemons()

		assert.NotNil(t, err)
		assert.Nil(t, pokemons)
	})
}
