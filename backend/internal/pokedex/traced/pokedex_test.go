package traced

import (
	"context"
	"testing"

	"github.com/defryheryanto/pokemon-helper/internal/pokedex"
	pokedexmock "github.com/defryheryanto/pokemon-helper/internal/pokedex/mock"
	"github.com/defryheryanto/pokemon-helper/internal/pokemon"
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

func TestTracedServiceGetAllPokedexCreatesSpan(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	service := pokedexmock.NewMockIService(ctrl)

	expected := []pokemon.PokemonData{{ID: 1, Name: "bulbasaur"}}
	service.EXPECT().GetAllPokedex(gomock.Any(), pokedex.GetAllPokedexFilter{Page: 1, PageSize: 10}).Return(expected, nil)

	tracer, recorder := newTestTracer()
	traced := NewTracedService(service, tracer)

	result, err := traced.GetAllPokedex(context.Background(), pokedex.GetAllPokedexFilter{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(result) != 1 || result[0].ID != expected[0].ID {
		t.Fatalf("expected %v, got %v", expected, result)
	}

	spans := recorder.Ended()
	if len(spans) != 1 {
		t.Fatalf("expected 1 span, got %d", len(spans))
	}
	if spans[0].Name() != "service-GetPokedex" {
		t.Fatalf("expected span name service-GetPokedex, got %s", spans[0].Name())
	}
	if !hasEvent(spans[0], "parameters") || !hasEvent(spans[0], "result") {
		t.Fatalf("expected parameters and result events, got %+v", spans[0].Events())
	}
}

func TestTracedServiceGetPokedexCreatesSpan(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	service := pokedexmock.NewMockIService(ctrl)

	expected := &pokemon.PokemonData{ID: 25, Name: "pikachu"}
	service.EXPECT().GetPokedex(gomock.Any(), "pikachu").Return(expected, nil)

	tracer, recorder := newTestTracer()
	traced := NewTracedService(service, tracer)

	result, err := traced.GetPokedex(context.Background(), "pikachu")
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
