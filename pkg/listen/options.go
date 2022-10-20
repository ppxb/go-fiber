package listen

import (
	"context"
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
	exit      func()
}

func WithHttpCtx(ctx context.Context) func(*HttpOptions) {
	return func(options *HttpOptions) {
		if !utils.InterfaceIsNil(ctx) {
			getHttpOptions(options).ctx = ctx
		}
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
			port:      8080,
			urlPrefix: "api",
		}
	}
	return options
}
