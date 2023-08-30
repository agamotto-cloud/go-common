package logger

import (
	"context"
	"github.com/agamotto-cloud/go-common/common/config"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	trace2 "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func init() {

	// 创建全局的 Tracer 实例
	// 创建一个新的 TraceProvider，并使用 Noop 实现

	traceProvider := trace2.NewTracerProvider()
	//trace.WithNewRoot()
	// 设置全局的 TraceProvider
	otel.SetTracerProvider(traceProvider)

	// 设置全局的 TraceContext 解析器，这里使用的是 W3C TraceContext 格式
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))

	tracer = otel.GetTracerProvider().Tracer(config.GetServerConfig().Name)

	log.Info().Msg("tracer init success")
}

func CreateSpan(ctx context.Context, operationName string) (context.Context, trace.Span) {
	//logg := log.Ctx(ctx)
	//	log.Logger.Info().Any("trace", ctx.Value("trace_id")).Msg("CreateSpan")
	// 创建一个 Span
	c, span := tracer.Start(ctx, operationName)
	//从c里面取出traceId和spanId
	logger := log.Logger.With().
		Str("span_id", span.SpanContext().SpanID().String()).
		Str("trace_id", span.SpanContext().TraceID().String()).
		Logger()
	return logger.WithContext(c), span
}
