package traced

import (
	"context"
	"testing"

	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
	pokemonmock "github.com/defryheryanto/pokemon-helper/internal/pokemon/mock"
	"github.com/golang/mock/gomock"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
)

func newTestTracer() (trace.Tracer, *tracetest.SpanRecorder) {
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	return provider.Tracer("test"), recorder
}

func hasEvent(span sdktrace.ReadOnlySpan, name string) bool {
	for _, event := range span.Events() {
		if event.Name == name {
			return true
		}
	}
	return false
}

func TestTracedServiceUpsertCreatesSpan(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	service := pokemonmock.NewMockService(ctrl)

	data := &pokemon.PokemonData{ID: 1, Name: "bulbasaur"}
	service.EXPECT().Upsert(gomock.Any(), data).Return(nil)

	tracer, recorder := newTestTracer()
	traced := NewTracedService(service, tracer)

	if err := traced.Upsert(context.Background(), data); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	spans := recorder.Ended()
	if len(spans) != 1 {
		t.Fatalf("expected 1 span, got %d", len(spans))
	}
	span := spans[0]
	if span.Name() != "service-UpsertPokemon" {
		t.Fatalf("expected span name service-UpsertPokemon, got %s", span.Name())
	}
	if !hasEvent(span, "parameters") || !hasEvent(span, "result") {
		t.Fatalf("expected parameters and result events, got %+v", span.Events())
	}
}

func TestTracedServiceGetByNameRecordsParameters(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	service := pokemonmock.NewMockService(ctrl)

	expected := &pokemon.PokemonData{ID: 4, Name: "charmander"}
	service.EXPECT().GetByName(gomock.Any(), "charmander").Return(expected, nil)

	tracer, recorder := newTestTracer()
	traced := NewTracedService(service, tracer)

	result, err := traced.GetByName(context.Background(), "charmander")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.ID != expected.ID {
		t.Fatalf("expected %v, got %v", expected, result)
	}

	spans := recorder.Ended()
	if len(spans) != 1 {
		t.Fatalf("expected 1 span, got %d", len(spans))
	}
	if !hasEvent(spans[0], "parameters") || !hasEvent(spans[0], "result") {
		t.Fatalf("expected parameters and result events, got %+v", spans[0].Events())
	}
}
