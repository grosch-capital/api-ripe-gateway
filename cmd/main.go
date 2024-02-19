package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/grosch-capital/api-ripe-gateway/internal/handlers"
)

type LogMiddleware struct {
	logger *log.Logger
}

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
	buf        bytes.Buffer
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	w.buf.Write(body)
	return w.ResponseWriter.Write(body)
}

func NewLogMiddleware(logger *log.Logger) *LogMiddleware {
	return &LogMiddleware{logger: logger}
}

func (m *LogMiddleware) Func() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			logRespWriter := NewLogResponseWriter(w)
			next.ServeHTTP(logRespWriter, r)

			m.logger.Printf(
				"duration=%s status=%d body=%s",
				time.Since(startTime).String(),
				logRespWriter.statusCode,
				logRespWriter.buf.String())
		})
	}
}

func main() {
	r := mux.NewRouter()

	logger := log.New(os.Stdout, "", log.LstdFlags)

	r.HandleFunc("/", handlers.LookupRAWHandler)
	r.HandleFunc("/ip/raw", handlers.LookupRAWHandler)
	r.HandleFunc("/ip/json", handlers.LookupJSONHandler)
	r.HandleFunc("/geo", handlers.LookupClientGeoHandler)
	r.HandleFunc("/geo/{ip}", handlers.LookupSpecGeoHandler)
	r.HandleFunc("/route", handlers.HealthCheckHandler)
	r.HandleFunc("/route/{host}", handlers.HealthCheckHandler)
	r.HandleFunc("/healthz", handlers.HealthCheckHandler)

	logMiddleware := NewLogMiddleware(logger)
	r.Use(logMiddleware.Func())

	logger.Fatalln(http.ListenAndServe(":8080", r))
}
