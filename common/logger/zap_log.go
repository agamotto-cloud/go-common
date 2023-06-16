package logger

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
	"os"
	"strings"
	"time"
)

//还需要给这个log集成分布式日志追踪，用最流行的那个开源的opentracing

func init() {
	output := zerolog.ConsoleWriter{
		Out: os.Stdout,
		FormatTimestamp: func(i interface{}) string {
			parse, _ := time.Parse(time.RFC3339Nano, i.(string))
			return parse.Format(time.RFC3339Nano)
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf(" %-6s ", i))
		},
	}
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = zerolog.New(output).With().
		Timestamp().CallerWithSkipFrameCount(2).Logger()

	log.Info().Msg("logger init success")
}

type C struct {
}

func (l C) Info(ctx context.Context, s string, i ...interface{}) {
	log.Ctx(ctx).Info().Msg(s)
}

func (l C) Warn(ctx context.Context, s string, i ...interface{}) {
	log.Ctx(ctx).Info().Msg(s)
}

func (l C) Error(ctx context.Context, s string, i ...interface{}) {
	log.Ctx(ctx).Info().Msg(s)
}

func (l C) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	log.Ctx(ctx).Info().Msgf("spanTime %s sql:%s,rowsAffected:%d", time.Since(begin), sql, rowsAffected)
}

type Interface interface {
	LogMode(int) Interface
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
}

type Log struct {
	ctx    context.Context
	span   trace.Span
	logger *zerolog.Logger
}

func (l Log) Debug(msg string, data ...interface{}) {

	l.logger.Debug().Msg(msg)
}

func (l Log) Info(msg string, data ...interface{}) {

	l.logger.Info().Msg(msg)
}

func (l Log) Warn(msg string, data ...interface{}) {
	l.logger.Warn().Msg(msg)
}

func (l Log) Error(msg string, data ...interface{}) {
	l.logger.Error().Msg(msg)
}

func Debug(ctx context.Context, msg string, data ...interface{}) {
	log.Debug().Msg(msg)
}

func Info(ctx context.Context, msg string, data ...interface{}) {
	log.Info().Msg(msg)
}

func Warn(ctx context.Context, msg string, data ...interface{}) {
	log.Warn().Msg(msg)
}

func Error(ctx context.Context, msg string, data ...interface{}) {
	log.Error().Msg(msg)
}
