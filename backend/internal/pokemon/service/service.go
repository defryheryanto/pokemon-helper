package service

import (
	"context"

	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
)

type Service struct {
	repository pokemon.Repository
}

func NewService(repository pokemon.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Upsert(ctx context.Context, data *pokemon.PokemonData) error {
	return s.repository.Upsert(ctx, data)
}

func (s *Service) List(ctx context.Context, filter pokemon.ListFilter) ([]pokemon.PokemonData, error) {
	return s.repository.List(ctx, filter)
}

func (s *Service) GetByID(ctx context.Context, id int) (*pokemon.PokemonData, error) {
	return s.repository.Get(ctx, pokemon.GetFilter{ID: id})
}

func (s *Service) GetByName(ctx context.Context, name string) (*pokemon.PokemonData, error) {
	return s.repository.Get(ctx, pokemon.GetFilter{Name: name})
}
