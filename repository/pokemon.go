package repository

import (
	"rest-api-golang/config"
	"rest-api-golang/entity"
	"rest-api-golang/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PokemonRepositoryInterface interface {
	GetAllPokemons() (*[]entity.Pokemon, error)
	GetPokemon(id string) (*[]entity.Pokemon, error)
	CreatePokemon(pokemon *entity.Pokemon) error
	DeletePokemon(id string) error
	UpdatePokemon(pokemon *entity.Pokemon, id string) error
}

type PokemonRepository struct {
	DB *gorm.DB
}

func NewPokemonRepository(db config.Database) PokemonRepositoryInterface {
	return &PokemonRepository{
		DB: db.DB,
	}
}

func (p *PokemonRepository) GetAllPokemons() (*[]entity.Pokemon, error) {
	pokemon := &[]entity.Pokemon{}
	if err := p.DB.Preload(clause.Associations).Find(pokemon).Error; err != nil {
		logger.Log.Error("Error GetAllPokemons")
		return nil, err
	}
	return pokemon, nil
}

func (p *PokemonRepository) GetPokemon(id string) (*[]entity.Pokemon, error) {
	pokemon := &[]entity.Pokemon{}
	if err := p.DB.Preload(clause.Associations).Where("id = ?", id).Find(pokemon).Error; err != nil {
		logger.Log.Error("Error GetPokemon")
		return nil, err
	}
	return pokemon, nil
}

func (p *PokemonRepository) CreatePokemon(pokemon *entity.Pokemon) error {
	if err := p.DB.Create(&pokemon).Error; err != nil {
		logger.Log.Error("Error CreatePokemon")
		return err
	}
	return nil
}

func (p *PokemonRepository) DeletePokemon(id string) error {
	var pokemon *entity.Pokemon
	if err := p.DB.Delete(&pokemon, id).Error; err != nil {
		logger.Log.Error("Error DeletePokemon")
		return err
	}
	return nil
}

func (p *PokemonRepository) UpdatePokemon(pokemon *entity.Pokemon, id string) error {
	if err := p.DB.Model(&pokemon).Where("id = ?", id).Updates(pokemon).Error; err != nil {
		logger.Log.Error("Error UpdatePokemon")
		return err
	}
	return nil
}
