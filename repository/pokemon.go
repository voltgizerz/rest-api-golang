package repository

import (
	"rest-api-golang/config"
	"rest-api-golang/entity"
	"rest-api-golang/logger"

	"gorm.io/gorm"
)

// PokemonRepositoryInterface - ,
type PokemonRepositoryInterface interface {
	GetAllPokemons(preloads ...string) (*[]entity.Pokemon, error)
	GetPokemon(id string, preloads ...string) (*[]entity.Pokemon, error)
	CreatePokemon(pokemon *entity.Pokemon) error
	DeletePokemon(id string) error
	UpdatePokemon(pokemon *entity.Pokemon, id string) error
	CreateTx() *gorm.DB
	RollbackTx() *gorm.DB
	CommitTx() error
}

// PokemonRepository - .
type PokemonRepository struct {
	DB *gorm.DB
}

// NewPokemonRepository - .
func NewPokemonRepository(dbConfig config.Database) PokemonRepositoryInterface {
	return &PokemonRepository{
		DB: dbConfig.ORM,
	}
}

// GetAllPokemons - .
func (p *PokemonRepository) GetAllPokemons(preloads ...string) (*[]entity.Pokemon, error) {
	pokemon := &[]entity.Pokemon{}
	if err := p.DBWithPreloads(preloads).Find(pokemon).Error; err != nil {
		logger.Log.Error("Error GetAllPokemons")
		return nil, err
	}
	return pokemon, nil
}

// GetPokemon - .
func (p *PokemonRepository) GetPokemon(id string, preloads ...string) (*[]entity.Pokemon, error) {
	pokemon := &[]entity.Pokemon{}
	if err := p.DBWithPreloads(preloads).Where("id = ?", id).Find(pokemon).Error; err != nil {
		logger.Log.Error("Error GetPokemon")
		return nil, err
	}
	return pokemon, nil
}

// CreatePokemon - .
func (p *PokemonRepository) CreatePokemon(pokemon *entity.Pokemon) error {
	if err := p.DB.Create(&pokemon).Error; err != nil {
		logger.Log.Error("Error CreatePokemon")
		return err
	}
	return nil
}

// DeletePokemon - .
func (p *PokemonRepository) DeletePokemon(id string) error {
	var pokemon *entity.Pokemon
	if err := p.DB.Delete(&pokemon, id).Error; err != nil {
		logger.Log.Error("Error DeletePokemon")
		return err
	}
	return nil
}

// UpdatePokemon - .
func (p *PokemonRepository) UpdatePokemon(pokemon *entity.Pokemon, id string) error {
	if err := p.DB.Model(&pokemon).Where("id = ?", id).Updates(pokemon).Error; err != nil {
		logger.Log.Error("Error UpdatePokemon")
		return err
	}
	return nil
}

// DBWithPreloads - preload data.
func (p *PokemonRepository) DBWithPreloads(preloads []string) *gorm.DB {
	dbConn := p.DB

	for _, preload := range preloads {
		dbConn = dbConn.Preload(preload)
	}

	return dbConn
}

// CreateTx - create database transaction
func (p *PokemonRepository) CreateTx() *gorm.DB {
	return p.DB.Begin()
}

// RollbackTx - rollback transaction
func (p *PokemonRepository) RollbackTx() *gorm.DB {
	return p.DB.Rollback()
}

// CommitTx - commit transaction
func (p *PokemonRepository) CommitTx() error {
	return p.DB.Commit().Error
}
