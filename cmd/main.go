package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/grosch-capital/api-ripe-gateway/internal/handlers"
	"github.com/newrelic/go-agent/v3/newrelic"
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

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("The Grosch RIPE"),
		newrelic.ConfigLicense("eu01xx8abfdac3510c18d83bcae625b4FFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		logger.Fatalln(err)
	}

	r.HandleFunc(newrelic.WrapHandleFunc(app, "/", handlers.LookupRAWHandler))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/ip/raw", handlers.LookupRAWHandler))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/ip/json", handlers.LookupJSONHandler))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/geo", handlers.LookupClientGeoHandler))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/geo/{ip}", handlers.LookupSpecGeoHandler))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/route", handlers.HealthCheckHandler))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/route/{host}", handlers.HealthCheckHandler))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/healthz", handlers.HealthCheckHandler))

	logMiddleware := NewLogMiddleware(logger)
	r.Use(logMiddleware.Func())

	logger.Fatalln(http.ListenAndServe(":8080", r))
}
