package traced

import (
	"context"
	"encoding/json"

	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type TracedService struct {
	pokedex.IService
	tracer trace.Tracer
}

func NewTracedService(baseService pokedex.IService, tracer trace.Tracer) *TracedService {
	return &TracedService{baseService, tracer}
}

func (s *TracedService) GetAllPokedex(ctx context.Context, filter pokedex.GetAllPokedexFilter) ([]pokemon.PokemonData, error) {
	ctx, span := s.tracer.Start(ctx, "service-GetPokedex")
	defer span.End()

	span.AddEvent(
		"parameters",
		trace.WithAttributes(
			attribute.Int("page", filter.Page),
			attribute.Int("pageSize", filter.PageSize),
		),
	)

	pokemons, err := s.IService.GetAllPokedex(ctx, filter)
	s.logResult(span, pokemons)

	return pokemons, err
}

func (s *TracedService) GetPokedex(ctx context.Context, pokemonName string) (*pokemon.PokemonData, error) {
	ctx, span := s.tracer.Start(ctx, "service-GetPokedex")
	defer span.End()

	span.AddEvent(
		"parameters",
		trace.WithAttributes(
			attribute.String("pokemonName", pokemonName),
		),
	)

	poke, err := s.IService.GetPokedex(ctx, pokemonName)
	s.logResult(span, poke)

	return poke, err
}

func (s *TracedService) logResult(span trace.Span, result interface{}) {
	returnAttribute := []attribute.KeyValue{}
	b, err := json.Marshal(result)
	if err == nil {
		returnAttribute = append(returnAttribute, attribute.String("return", string(b)))
	}

	span.AddEvent(
		"result",
		trace.WithAttributes(returnAttribute...),
	)
}
