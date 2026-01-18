package tracer

import (
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.8.0"
	"go.opentelemetry.io/otel/trace"
)

func logRequest(span trace.Span, r *http.Request) {
	b, err := httputil.DumpRequest(r, r.Body != nil)
	if err != nil {
		span.RecordError(err)
		return
	}

	span.AddEvent("request", trace.WithAttributes(attribute.String("http.request", string(b))))
}

func logResponse(span trace.Span, resp *http.Response) {
	b, err := httputil.DumpResponse(resp, resp.Body != nil)
	if err != nil {
		span.RecordError(err)
		return
	}
	attr := trace.WithAttributes(attribute.String("http.response", string(b)))
	span.AddEvent("response", attr)
}

func TracerMiddleware(tracer trace.Tracer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if tracer == nil {
				h.ServeHTTP(w, r)
				return
			}

			routeName, err := mux.CurrentRoute(r).GetPathTemplate()
			if err != nil {
				routeName = r.URL.Path
			}
			ctx, span := tracer.Start(
				r.Context(),
				routeName,
				trace.WithSpanKind(trace.SpanKindServer),
				trace.WithAttributes(semconv.NetAttributesFromHTTPRequest("tcp", r)...),
				trace.WithAttributes(semconv.EndUserAttributesFromHTTPRequest(r)...),
				trace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest(routeName, "", r)...),
			)
			defer span.End()

			logRequest(span, r)
			responseWriterWrapper := newResponseWriter(w)
			h.ServeHTTP(responseWriterWrapper, r.WithContext(ctx))
			logResponse(span, responseWriterWrapper.Response())
		})
	}
}
