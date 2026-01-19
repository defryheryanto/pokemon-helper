package tracer

import (
	"net/http"
	"net/http/httptest"
)

type responseWriter struct {
	http.ResponseWriter
	*httptest.ResponseRecorder
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter:   w,
		ResponseRecorder: httptest.NewRecorder(),
	}
}

func (rw *responseWriter) Header() http.Header {
	return rw.ResponseWriter.Header()
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.ResponseRecorder.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.ResponseRecorder.WriteHeader(statusCode)
	rw.ResponseRecorder.WriteHeader(statusCode)
}

func (rw *responseWriter) Response() *http.Response {
	rw.ResponseRecorder.Result().Header = rw.ResponseWriter.Header()
	return rw.ResponseRecorder.Result()
}
