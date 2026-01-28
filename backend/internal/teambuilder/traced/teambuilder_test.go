package traced

import (
	"context"
	"testing"

	"github.com/defryheryanto/pokemon-helper/internal/pokemontype"
	teambuildermock "github.com/defryheryanto/pokemon-helper/internal/teambuilder/mock"
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

func TestTracedServiceCalculateTypeCoverage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	base := teambuildermock.NewMockIService(ctrl)

	covered := []pokemontype.IType{pokemontype.FireType}
	uncovered := []pokemontype.IType{pokemontype.WaterType}
	base.EXPECT().CalculateTypeCoverage(gomock.Any(), []string{"pikachu"}).Return(covered, uncovered, nil)

	tracer, recorder := newTestTracer()
	traced := NewTracedService(base, tracer)

	gotCovered, gotUncovered, err := traced.CalculateTypeCoverage(context.Background(), []string{"pikachu"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(gotCovered) != 1 || gotCovered[0] != pokemontype.FireType {
		t.Fatalf("expected covered %v, got %v", covered, gotCovered)
	}
	if len(gotUncovered) != 1 || gotUncovered[0] != pokemontype.WaterType {
		t.Fatalf("expected uncovered %v, got %v", uncovered, gotUncovered)
	}

	spans := recorder.Ended()
	if len(spans) != 1 {
		t.Fatalf("expected 1 span, got %d", len(spans))
	}
	if spans[0].Name() != "TeamBuilderService:CalculateTypeCoverage" {
		t.Fatalf("expected span name TeamBuilderService:CalculateTypeCoverage, got %s", spans[0].Name())
	}
	if !hasEvent(spans[0], "parameters") || !hasEvent(spans[0], "result") {
		t.Fatalf("expected parameters and result events, got %+v", spans[0].Events())
	}
}

func TestTracedServiceCalculateSuggestedType(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	base := teambuildermock.NewMockIService(ctrl)

	uncovered := []pokemontype.IType{pokemontype.FireType}
	suggested := []pokemontype.IType{pokemontype.WaterType}
	base.EXPECT().CalculateSuggestedType(gomock.Any(), uncovered, 1).Return(suggested)

	tracer, recorder := newTestTracer()
	traced := NewTracedService(base, tracer)

	result := traced.CalculateSuggestedType(context.Background(), uncovered, 1)
	if len(result) != 1 || result[0] != pokemontype.WaterType {
		t.Fatalf("expected %v, got %v", suggested, result)
	}

	spans := recorder.Ended()
	if len(spans) != 1 {
		t.Fatalf("expected 1 span, got %d", len(spans))
	}
	if spans[0].Name() != "TeamBuilderService:CalculateSuggestedType" {
		t.Fatalf("expected span name TeamBuilderService:CalculateSuggestedType, got %s", spans[0].Name())
	}
	if !hasEvent(spans[0], "parameters") || !hasEvent(spans[0], "result") {
		t.Fatalf("expected parameters and result events, got %+v", spans[0].Events())
	}
}
