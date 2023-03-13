package main

import (
	"github.com/defry256/pokemon-helper/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
)

func setupTracer() *trace.TracerProvider {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(config.JaegerCollectorURL()),
	))
	if err != nil {
		panic(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("pokemon-helper"),
			attribute.String("environment", config.Environment()),
		)),
	)
	return tp
}
