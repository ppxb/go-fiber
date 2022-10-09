package log

import (
	"context"
	"github.com/ppxb/go-fiber/pkg/constant"
)

type Wrapper struct {
	log    Interface
	fields map[string]interface{}
}

func (w *Wrapper) Trace(args ...interface{}) {
	if !w.log.Options().level.Enabled(TraceLevel) {
		return
	}
	//ns := copyFields(w.fields)
	//if w.log.Options().lineNum {
	//	ns[constant.LogLineNumKey] = fileWithLineNum(w.log.Options())
	//}
}

func (w *Wrapper) Info(args ...interface{}) {
	if !w.log.Options().level.Enabled(InfoLevel) {
		return
	}
	ns := copyFields(w.fields)
	if w.log.Options().lineNum {

	}
	if len(args) > 1 {
		if format, ok := args[0].(string); ok {
			w.log.WithFields(ns).LogF(InfoLevel, format, args[1:]...)
			return
		}
	}
	w.log.WithFields(ns).Log(InfoLevel, args...)
}

func (w *Wrapper) WithContext(ctx context.Context) *Wrapper {
	ns := copyFields(w.fields)

	return &Wrapper{
		log:    w.log,
		fields: ns,
	}
}

func (w *Wrapper) WithError(err error) *Wrapper {
	ns := copyFields(w.fields)
	ns[constant.LogErrorKey] = err

	return &Wrapper{
		log:    w.log,
		fields: ns,
	}
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
