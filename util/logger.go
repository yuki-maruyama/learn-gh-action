package util

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
)

type LogHandler struct {
	slog.Handler
}

func NewLogHandler() LogHandler {
	return LogHandler{
		Handler: slog.NewJSONHandler(os.Stdout, nil),
	}
}

func (h LogHandler) Handle(ctx context.Context, r slog.Record) error {
	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	if sc.IsValid() {
		r.AddAttrs(
			slog.String("trace_id", sc.TraceID().String()),
			slog.String("span_id", sc.SpanID().String()),
		)
	}
	bg := baggage.FromContext(ctx)
	members := bg.Members()
	for _, member := range members {
		r.AddAttrs(slog.String(member.Key(), member.Value()))
	}
	return h.Handler.Handle(ctx, r)
}
