package traced

import (
	"context"
	"encoding/json"

	"github.com/defry256/pokemon-helper/internal/pokemontype"
	"github.com/defry256/pokemon-helper/internal/teambuilder"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type tracedService struct {
	teambuilder.IService
	tracer trace.Tracer
}

func NewTracedService(baseService teambuilder.IService, tracer trace.Tracer) *tracedService {
	return &tracedService{baseService, tracer}
}

func (s *tracedService) CalculateTypeCoverage(ctx context.Context, pokemonNames []string) (coveredTypes, uncoveredTypes []pokemontype.IType, err error) {
	ctx, span := s.tracer.Start(ctx, "TeamBuilderService:CalculateTypeCoverage")
	defer span.End()

	span.AddEvent(
		"parameters",
		trace.WithAttributes(
			attribute.StringSlice("pokemonNames", pokemonNames),
		),
	)

	coveredTypes, uncoveredTypes, err = s.IService.CalculateTypeCoverage(ctx, pokemonNames)
	if err != nil {
		span.RecordError(err)
		return nil, nil, err
	}

	attributes := []attribute.KeyValue{}
	b, err := json.Marshal(coveredTypes)
	if err == nil {
		attributes = append(attributes, attribute.String("coveredTypes", string(b)))
	}
	b, err = json.Marshal(uncoveredTypes)
	if err == nil {
		attributes = append(attributes, attribute.String("uncoveredTypes", string(b)))
	}

	span.AddEvent(
		"result",
		trace.WithAttributes(attributes...),
	)

	return coveredTypes, uncoveredTypes, nil
}

func (s *tracedService) CalculateSuggestedType(ctx context.Context, uncoveredTypes []pokemontype.IType, suggestLength int) []pokemontype.IType {
	ctx, span := s.tracer.Start(ctx, "TeamBuilderService:CalculateSuggestedType")
	defer span.End()

	paramsAttributes := []attribute.KeyValue{
		attribute.Int("suggestLength", suggestLength),
	}
	b, err := json.Marshal(uncoveredTypes)
	if err == nil {
		paramsAttributes = append(paramsAttributes, attribute.String("uncoveredTypes", string(b)))
	}

	span.AddEvent(
		"parameters",
		trace.WithAttributes(paramsAttributes...),
	)

	pokeTypes := s.IService.CalculateSuggestedType(ctx, uncoveredTypes, suggestLength)

	attributes := []attribute.KeyValue{}
	b, err = json.Marshal(pokeTypes)
	if err == nil {
		attributes = append(attributes, attribute.String("result", string(b)))
	}

	span.AddEvent(
		"result",
		trace.WithAttributes(attributes...),
	)

	return pokeTypes
}
