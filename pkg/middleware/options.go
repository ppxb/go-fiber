package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/app/models"
	"github.com/ppxb/go-fiber/pkg/constant"
	"github.com/ppxb/go-fiber/pkg/req"
	"github.com/ppxb/go-fiber/pkg/resp"
)

type CorsOptions struct {
	origin     string
	header     string
	expose     string
	method     string
	credential string
}

type JwtOptions struct {
	realm              string
	key                string
	timeout            int
	maxRefresh         int
	tokenLookup        string
	tokenHeaderName    string
	sendCookie         bool
	cookieName         string
	privateBytes       []byte
	success            func()
	successWithData    func(c *gin.Context, data interface{}) resp.Resp
	failWithMsg        func(format interface{}, a ...interface{})
	failWithCodeAndMsg func(code int, format interface{}, a ...interface{})
	loginPwdCheck      func(c *gin.Context, r req.LoginCheck) (user models.SysUser, err error)
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

func WithJwtRealm(realm string) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).realm = realm
	}
}

func WithJwtKey(key string) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).key = key
	}
}

func WithJwtTimeout(timeout int) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).timeout = timeout
	}
}

func WithJwtMaxRefresh(maxRefresh int) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).maxRefresh = maxRefresh
	}
}

func WithJwtTokenLookup(tokenLookup string) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).tokenLookup = tokenLookup
	}
}

func WithJwtTokenHeaderName(tokenHeaderName string) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).tokenHeaderName = tokenHeaderName
	}
}

func WithJwtSendCookie(flag bool) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).sendCookie = flag
	}
}

func WithJwtCookieName(cookieName string) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).cookieName = cookieName
	}
}

func WithJwtPrivateBytes(bs []byte) func(*JwtOptions) {
	return func(options *JwtOptions) {
		getJwtOptions(options).privateBytes = bs
	}
}

func WithJwtSuccess(fun func()) func(*JwtOptions) {
	return func(options *JwtOptions) {
		if fun != nil {
			getJwtOptions(options).success = fun
		}
	}
}

//func WithJwtSuccessWithData(fun func(...interface{})) func(*JwtOptions) {
//	return func(options *JwtOptions) {
//		if fun != nil {
//			getJwtOptions(options).successWithData = fun
//		}
//	}
//}

func WithJwtFailWithMsg(fun func(format interface{}, a ...interface{})) func(*JwtOptions) {
	return func(options *JwtOptions) {
		if fun != nil {
			getJwtOptions(options).failWithMsg = fun
		}
	}
}

func WithJwtFailWithCodeAndMsg(fun func(code int, format interface{}, a ...interface{})) func(*JwtOptions) {
	return func(options *JwtOptions) {
		if fun != nil {
			getJwtOptions(options).failWithCodeAndMsg = fun
		}
	}
}

func WithJwtLoginPwdCheck(fun func(c *gin.Context, r req.LoginCheck) (user models.SysUser, err error)) func(*JwtOptions) {
	return func(options *JwtOptions) {
		if fun != nil {
			getJwtOptions(options).loginPwdCheck = fun
		}
	}
}

func getJwtOptions(options *JwtOptions) *JwtOptions {
	if options == nil {
		return &JwtOptions{
			realm:           "test jwt",
			key:             "test secret",
			timeout:         24,
			maxRefresh:      168,
			tokenLookup:     "header: Authorization, repository: token, cookie: jwt",
			tokenHeaderName: "Bearer",
		}
	}
	return options
}
