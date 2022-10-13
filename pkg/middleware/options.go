package middleware

import "github.com/ppxb/go-fiber/pkg/constant"

type CorsOptions struct {
	origin     string
	header     string
	expose     string
	method     string
	credential string
}

func getCorsOptions(options *CorsOptions) *CorsOptions {
	if options == nil {
		return &CorsOptions{
			origin:     constant.MiddlewareCorsOrigin,
			header:     constant.MiddlewareCorsHeaders,
			expose:     constant.MiddlewareCorsExpose,
			method:     constant.MiddlewareCorsMethods,
			credential: constant.MiddlewareCorsCredentials,
		}
	}
	return options
}
