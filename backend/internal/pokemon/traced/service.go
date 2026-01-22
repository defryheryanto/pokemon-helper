package traced

import (
	"context"
	"encoding/json"

	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type TracedService struct {
	pokemon.Service
	tracer trace.Tracer
}

func NewTracedService(baseService pokemon.Service, tracer trace.Tracer) *TracedService {
	return &TracedService{baseService, tracer}
}

func (s *TracedService) Upsert(ctx context.Context, data *pokemon.PokemonData) error {
	ctx, span := s.tracer.Start(ctx, "service-UpsertPokemon")
	defer span.End()

	if data != nil {
		span.AddEvent(
			"parameters",
			trace.WithAttributes(
				attribute.Int("id", data.ID),
				attribute.String("name", data.Name),
			),
		)
	}

	err := s.Service.Upsert(ctx, data)
	s.logResult(span, err)

	return err
}

func (s *TracedService) List(ctx context.Context, filter pokemon.ListFilter) ([]pokemon.PokemonData, error) {
	ctx, span := s.tracer.Start(ctx, "service-ListPokemons")
	defer span.End()

	span.AddEvent(
		"parameters",
		trace.WithAttributes(
			attribute.Int("limit", filter.Limit),
			attribute.Int("offset", filter.Offset),
		),
	)

	result, err := s.Service.List(ctx, filter)
	s.logResult(span, result)

	return result, err
}

func (s *TracedService) GetByID(ctx context.Context, id int) (*pokemon.PokemonData, error) {
	ctx, span := s.tracer.Start(ctx, "service-GetPokemonByID")
	defer span.End()

	span.AddEvent(
		"parameters",
		trace.WithAttributes(
			attribute.Int("id", id),
		),
	)

	result, err := s.Service.GetByID(ctx, id)
	s.logResult(span, result)

	return result, err
}

func (s *TracedService) GetByName(ctx context.Context, name string) (*pokemon.PokemonData, error) {
	ctx, span := s.tracer.Start(ctx, "service-GetPokemonByName")
	defer span.End()

	span.AddEvent(
		"parameters",
		trace.WithAttributes(
			attribute.String("name", name),
		),
	)

	result, err := s.Service.GetByName(ctx, name)
	s.logResult(span, result)

	return result, err
}

func (s *TracedService) logResult(span trace.Span, result interface{}) {
	returnAttributes := []attribute.KeyValue{}
	b, err := json.Marshal(result)
	if err == nil {
		returnAttributes = append(returnAttributes, attribute.String("return", string(b)))
	}

	span.AddEvent(
		"result",
		trace.WithAttributes(returnAttributes...),
	)
}
