package pokemon

//go:generate go run github.com/golang/mock/mockgen -source=interface.go -package=mock -destination=mock/mock.go

import "context"

type Service interface {
	Upsert(ctx context.Context, data *PokemonData) error
	List(ctx context.Context, filter ListFilter) ([]PokemonData, error)
	GetByID(ctx context.Context, id int) (*PokemonData, error)
	GetByName(ctx context.Context, name string) (*PokemonData, error)
}

type Repository interface {
	Upsert(ctx context.Context, data *PokemonData) error
	Get(ctx context.Context, filter GetFilter) (*PokemonData, error)
	List(ctx context.Context, filter ListFilter) ([]PokemonData, error)
}
