package kontext

import (
	"context"

	eventV1 "github.com/AlpacaLabs/protorepo-event-go/alpacalabs/event/v1"
	"go.opentelemetry.io/otel/api/trace"
)

func GetTraceInfo(ctx context.Context) eventV1.TraceInfo {
	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	return eventV1.TraceInfo{
		TraceId: sc.TraceID.String(),
		Sampled: sc.IsSampled(),
	}
}
