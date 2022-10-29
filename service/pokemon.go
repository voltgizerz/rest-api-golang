package service

import (
	"rest-api-golang/entity"
	"rest-api-golang/repository"
)

type PokemonServiceInterface interface {
	GetAllPokemons() (*[]entity.Pokemon, error)
	GetPokemon(id string) (*[]entity.Pokemon, error)
	CreatePokemon(pokemon *entity.Pokemon) error
	DeletePokemon(id string) error
	UpdatePokemon(pokemon *entity.Pokemon, id string) error
}

type PokemonService struct {
	PokemonRepository repository.PokemonRepositoryInterface
}

func NewPokemonService(repoPokemon repository.PokemonRepositoryInterface) PokemonServiceInterface {
	return &PokemonService{
		PokemonRepository: repoPokemon,
	}
}

func (p *PokemonService) GetAllPokemons() (*[]entity.Pokemon, error) {
	return p.PokemonRepository.GetAllPokemons()
}

func (p *PokemonService) GetPokemon(id string) (*[]entity.Pokemon, error) {
	return p.PokemonRepository.GetPokemon(id)
}

func (p *PokemonService) CreatePokemon(pokemon *entity.Pokemon) error {
	return p.PokemonRepository.CreatePokemon(pokemon)
}

func (p *PokemonService) DeletePokemon(id string) error {
	return p.PokemonRepository.DeletePokemon(id)
}

func (p *PokemonService) UpdatePokemon(pokemon *entity.Pokemon, id string) error {
	return p.PokemonRepository.UpdatePokemon(pokemon, id)
}
