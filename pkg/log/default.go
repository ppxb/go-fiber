package log

import "context"

var DefaultWrapper *Wrapper

func init() {
	DefaultWrapper = &Wrapper{
		log:    New(),
		fields: map[string]interface{}{},
	}
}

func Trace(args ...interface{}) {
	DefaultWrapper.Trace(args...)
}

func Info(args ...interface{}) {
	DefaultWrapper.Info(args...)
}

func WithError(err error) *Wrapper {
	return DefaultWrapper.WithError(err)
}

func WithContext(ctx context.Context) *Wrapper {
	return DefaultWrapper.WithContext(ctx)
}
