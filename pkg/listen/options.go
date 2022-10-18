package listen

import (
	"context"
	"github.com/ppxb/go-fiber/configs"
	"github.com/ppxb/go-fiber/pkg/utils"
	"net/http"
)

type HttpOptions struct {
	ctx       context.Context
	host      string
	port      int
	pprofPort int
	urlPrefix string
	handler   http.Handler
	config    *configs.Config
	exit      func()
}

func WithHttpConfig(str string) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptions(options).config = configs.New(str)
	}
}

func WithHttpCtx(ctx context.Context) func(*HttpOptions) {
	return func(options *HttpOptions) {
		if !utils.InterfaceIsNil(ctx) {
			getHttpOptions(options).ctx = ctx
		}
	}
}

func WithHttpHost(s string) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptions(options).host = s
	}
}

func WithHttpPort(i int) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptions(options).port = i
	}
}

func WithHttpHandler(h http.Handler) func(*HttpOptions) {
	return func(options *HttpOptions) {
		getHttpOptions(options).handler = h
	}
}

func WithHttpExit(f func()) func(*HttpOptions) {
	return func(options *HttpOptions) {
		if f != nil {
			getHttpOptions(options).exit = f
		}
	}
}

func getHttpOptions(options *HttpOptions) *HttpOptions {
	if options == nil {
		return &HttpOptions{
			ctx:       context.Background(),
			host:      "0.0.0.0",
			port:      8888,
			urlPrefix: "api",
		}
	}
	return options
}
