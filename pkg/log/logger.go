package log

import "github.com/ppxb/go-fiber/pkg/constant"

type Interface interface {
	Options() Options
	WithFields(fields map[string]interface{}) Interface
	Log(level Level, v ...interface{})
	LogF(level Level, format string, v ...interface{})
}

func New(options ...func(*Options)) (l Interface) {
	ops := getOptions(nil)
	for _, f := range options {
		f(ops)
	}

	switch ops.category {
	case constant.LogCategoryZap:
		l = NewZap(ops)
	default:
		l = NewZap(ops)
	}
	return l
}

func fileWithLineNum(ops Options) {

}
