package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type (
	Logger = logrus.Logger
	Entry  = logrus.Entry
	Level  = logrus.Level
)

const (
	PanicLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func SetLevel(l Level) {
	logrus.SetLevel(l)
}

func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	default:
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}
}

const (
	TraceIdKey  = "trace_id"
	UserIdKey   = "user_id"
	UserNameKey = "user_name"
	TagKey      = "tag"
	StackKey    = "stack"
)

type (
	traceIdKey  = struct{}
	userIdKey   = struct{}
	userNameKey = struct{}
	tagKey      = struct{}
	stackKey    = struct{}
)

func NewTraceIdContext(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, traceIdKey{}, traceId)
}

func FromTraceIdContext(ctx context.Context) string {
	v := ctx.Value(traceIdKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewUserIdContext(ctx context.Context, userId uint64) context.Context {
	return context.WithValue(ctx, userIdKey{}, userId)
}

func FromUserIdContext(ctx context.Context) uint64 {
	v := ctx.Value(userIdKey{})
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
	return 0
}

func NewUserNameContext(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey{}, userName)
}

func FromUserNameContext(ctx context.Context) string {
	v := ctx.Value(userNameKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

func FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func NewStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

func FromStackContext(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}
	return nil
}

func WithContext(ctx context.Context) *Entry {
	fields := logrus.Fields{}

	if v := FromTraceIdContext(ctx); v != "" {
		fields[TraceIdKey] = v
	}

	if v := FromUserIdContext(ctx); v != 0 {
		fields[UserIdKey] = v
	}

	if v := FromUserNameContext(ctx); v != "" {
		fields[UserNameKey] = v
	}

	if v := FromTagContext(ctx); v != "" {
		fields[TagKey] = v
	}

	if v := FromStackContext(ctx); v != nil {
		fields[StackKey] = fmt.Sprintf("%+v", v)
	}

	return logrus.WithContext(ctx).WithFields(fields)
}

var (
	Tracef          = logrus.Tracef
	Debugf          = logrus.Debugf
	Infof           = logrus.Infof
	Warnf           = logrus.Warnf
	Errorf          = logrus.Errorf
	Fatalf          = logrus.Fatalf
	Panicf          = logrus.Panicf
	Printf          = logrus.Printf
	SetOutput       = logrus.SetOutput
	SetReportCaller = logrus.SetReportCaller
	StandardLogger  = logrus.StandardLogger
	ParseLevel      = logrus.ParseLevel
)
