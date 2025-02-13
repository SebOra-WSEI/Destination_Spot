package swo

import (
	"github.com/solarwinds/apm-go/swo"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func Start() (func(), error) {
	cb, err := swo.Start(
		semconv.ServiceVersion("v0.0.1"),
		attribute.String("environment", "testing"),
	)

	return cb, err
}
