package service

import (
	"context"
	"errors"
	"testing"

	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	pokemonmock "github.com/defryheryanto/pokemon-helper/internal/pokemon/mock"
	"github.com/golang/mock/gomock"
)

func TestServiceUpsert(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := pokemonmock.NewMockRepository(ctrl)
	service := NewService(repo)

	data := &pokemon.PokemonData{ID: 25, Name: "pikachu"}
	repo.EXPECT().Upsert(gomock.Any(), data).Return(nil)

	if err := service.Upsert(context.Background(), data); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestServiceList(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := pokemonmock.NewMockRepository(ctrl)
	service := NewService(repo)

	filter := pokemon.ListFilter{Limit: 10, Offset: 5}
	expected := []pokemon.PokemonData{{ID: 1, Name: "bulbasaur"}}
	repo.EXPECT().List(gomock.Any(), filter).Return(expected, nil)

	result, err := service.List(context.Background(), filter)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result) != len(expected) || result[0].ID != expected[0].ID {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestServiceGetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := pokemonmock.NewMockRepository(ctrl)
	service := NewService(repo)

	expected := &pokemon.PokemonData{ID: 7, Name: "squirtle"}
	repo.EXPECT().Get(gomock.Any(), pokemon.GetFilter{ID: 7}).Return(expected, nil)

	result, err := service.GetByID(context.Background(), 7)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.ID != expected.ID {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestServiceGetByNamePropagatesError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := pokemonmock.NewMockRepository(ctrl)
	service := NewService(repo)

	repoErr := errors.New("read failure")
	repo.EXPECT().Get(gomock.Any(), pokemon.GetFilter{Name: "eevee"}).Return(nil, repoErr)

	if _, err := service.GetByName(context.Background(), "eevee"); !errors.Is(err, repoErr) {
		t.Fatalf("expected error %v, got %v", repoErr, err)
	}
}
