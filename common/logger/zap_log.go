package logger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.CallerMarshalFunc = setLogMarshalFunc
	zerolog.CallerSkipFrameCount = 3
	log.Info().Interface("test", "test").Msg("test")
}

func setLogMarshalFunc(pc uintptr, file string, line int) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file + ":" + strconv.Itoa(line)
}

type C struct {
}

func (l C) Info(ctx context.Context, s string, i ...interface{}) {
	log.Info().Msg(s)
}

func (l C) Warn(ctx context.Context, s string, i ...interface{}) {
	log.Info().Msg(s)
}

func (l C) Error(ctx context.Context, s string, i ...interface{}) {
	log.Info().Msg(s)
}

func (l C) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
}

type Interface interface {
	LogMode(int) Interface
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
}

type Log struct {
	ctx context.Context
}

func GetLog(ctx context.Context) Log {
	return Log{ctx: ctx}
}

func (l Log) Debug(msg string, data ...interface{}) {
	log.Debug().Msg(msg)
}

func (l Log) Info(msg string, data ...interface{}) {
	log.Info().Msg(msg)
}

func (l Log) Warn(msg string, data ...interface{}) {
	log.Warn().Msg(msg)
}

func (l Log) Error(msg string, data ...interface{}) {
	log.Error().Msg(msg)
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
