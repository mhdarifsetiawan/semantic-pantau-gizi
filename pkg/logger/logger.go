package logger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type contextKey struct{}

// Logger wraps zerolog.Logger.
type Logger struct {
	zl zerolog.Logger
}

// Config holds logger configuration.
type Config struct {
	// Level: debug, info, warn, error. Default: info.
	Level string
	// Pretty enables human-readable output. Use only in development.
	Pretty bool
	// ServiceName is added to every log entry.
	ServiceName string
}

// New creates a new Logger based on the given Config.
func New(cfg Config) *Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}

	var w io.Writer = os.Stdout
	if cfg.Pretty {
		w = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	}

	zl := zerolog.New(w).
		Level(level).
		With().
		Timestamp().
		Str("service", cfg.ServiceName).
		Logger()

	return &Logger{zl: zl}
}

// WithContext returns a new context with the logger embedded.
func (l *Logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey{}, l)
}

// FromContext extracts the logger from context.
// Falls back to a default stdout logger if none found.
func FromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(contextKey{}).(*Logger); ok {
		return l
	}
	return New(Config{Level: "info"})
}

// With returns a new Logger with additional fixed fields.
func (l *Logger) With(fields map[string]any) *Logger {
	ctx := l.zl.With()
	for k, v := range fields {
		ctx = ctx.Interface(k, v)
	}
	return &Logger{zl: ctx.Logger()}
}

// WithRequestID returns a new Logger with request_id field.
func (l *Logger) WithRequestID(id string) *Logger {
	return &Logger{zl: l.zl.With().Str("request_id", id).Logger()}
}

// WithUserID returns a new Logger with user_id field.
func (l *Logger) WithUserID(id string) *Logger {
	return &Logger{zl: l.zl.With().Str("user_id", id).Logger()}
}

func (l *Logger) Debug(msg string, fields ...map[string]any) {
	e := l.zl.Debug()
	l.applyFields(e, fields...).Msg(msg)
}

func (l *Logger) Info(msg string, fields ...map[string]any) {
	e := l.zl.Info()
	l.applyFields(e, fields...).Msg(msg)
}

func (l *Logger) Warn(msg string, fields ...map[string]any) {
	e := l.zl.Warn()
	l.applyFields(e, fields...).Msg(msg)
}

func (l *Logger) Error(msg string, err error, fields ...map[string]any) {
	e := l.zl.Error().Err(err)
	l.applyFields(e, fields...).Msg(msg)
}

func (l *Logger) Fatal(msg string, err error, fields ...map[string]any) {
	e := l.zl.Fatal().Err(err)
	l.applyFields(e, fields...).Msg(msg)
}

func (l *Logger) applyFields(e *zerolog.Event, fields ...map[string]any) *zerolog.Event {
	for _, f := range fields {
		for k, v := range f {
			e = e.Interface(k, v)
		}
	}
	return e
}
