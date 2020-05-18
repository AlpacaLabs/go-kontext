package kontext

import (
	"context"

	"go.opentelemetry.io/otel/api/trace"
)

func GetTraceID(ctx context.Context) string {
	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	return sc.TraceID.String()
}
