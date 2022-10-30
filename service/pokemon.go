package service

import (
	"rest-api-golang/entity"
	"rest-api-golang/repository"
)

// PokemonServiceInterface - ...
type PokemonServiceInterface interface {
	GetAllPokemons() (*[]entity.Pokemon, error)
	GetPokemon(id string) (*[]entity.Pokemon, error)
	CreatePokemon(pokemon *entity.Pokemon) error
	DeletePokemon(id string) error
	UpdatePokemon(pokemon *entity.Pokemon, id string) error
}

// PokemonService - ...
type PokemonService struct {
	PokemonRepository repository.PokemonRepositoryInterface
}

// NewPokemonService - ...
func NewPokemonService(repoPokemon repository.PokemonRepositoryInterface) PokemonServiceInterface {
	return &PokemonService{
		PokemonRepository: repoPokemon,
	}
}

// GetAllPokemons - ...
func (p *PokemonService) GetAllPokemons() (*[]entity.Pokemon, error) {
	return p.PokemonRepository.GetAllPokemons()
}

// GetPokemon - ...
func (p *PokemonService) GetPokemon(id string) (*[]entity.Pokemon, error) {
	return p.PokemonRepository.GetPokemon(id)
}

// CreatePokemon - ...
func (p *PokemonService) CreatePokemon(pokemon *entity.Pokemon) error {
	return p.PokemonRepository.CreatePokemon(pokemon)
}
// DeletePokemon - ...
func (p *PokemonService) DeletePokemon(id string) error {
	return p.PokemonRepository.DeletePokemon(id)
}

// UpdatePokemon - ...
func (p *PokemonService) UpdatePokemon(pokemon *entity.Pokemon, id string) error {
	return p.PokemonRepository.UpdatePokemon(pokemon, id)
}
