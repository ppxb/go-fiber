package middleware

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/ppxb/go-fiber/pkg/constant"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/resp"
	"github.com/ppxb/go-fiber/pkg/utils"
	"net/http"
	"time"
)

func Jwt(options ...func(*JwtOptions)) gin.HandlerFunc {
	ops := getJwtOptions(nil)
	for _, f := range options {
		f(ops)
	}
	mw := initJwt(*ops)

	return func(c *gin.Context) {
		claims, err := mw.GetClaimsFromJWT(c)
		if err != nil {
			unauthorized(c, http.StatusUnauthorized, err, *ops)
			return
		}

		if claims["exp"] == nil {
			unauthorized(c, http.StatusBadRequest, jwt.ErrMissingExpField, *ops)
			return
		}

		if _, ok := claims["exp"].(float64); !ok {
			unauthorized(c, http.StatusBadRequest, jwt.ErrWrongFormatOfExp, *ops)
			return
		}

		if int64(claims["exp"].(float64)) < mw.TimeFunc().Unix() {
			unauthorized(c, http.StatusUnauthorized, jwt.ErrExpiredToken, *ops)
			return
		}

		c.Set("JWT_PAYLOAD", claims)
		i := identity(c)

		if i != nil {
			c.Set(mw.IdentityKey, i)
		}

		if !authorizator(i, c) {
			unauthorized(c, http.StatusForbidden, jwt.ErrForbidden, *ops)
			return
		}

		c.Next()
	}
}

func initJwt(ops JwtOptions) *jwt.GinJWTMiddleware {
	j, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         ops.realm,
		Key:           []byte(ops.key),
		Timeout:       time.Hour * time.Duration(ops.timeout),
		MaxRefresh:    time.Hour * time.Duration(ops.maxRefresh),
		TokenLookup:   ops.tokenLookUp,
		TokenHeadName: ops.tokenHeaderName,
		SendCookie:    ops.sendCookie,
		CookieName:    ops.cookieName,
		TimeFunc:      time.Now,
	})

	if err != nil {
		panic(nil)
	}
	return j
}

func authorizator(i interface{}, c *gin.Context) bool {
	if v, ok := i.(string); ok {
		userId := utils.Str2Int64(v)
		c.Set(constant.MiddlewareJwtUserCtxKey, userId)
		return true
	}
	return false
}

func unauthorized(c *gin.Context, code int, err error, ops JwtOptions) {
	log.WithContext(c).WithError(err).Warn("[Auth] jwt auth check failed, code: %d", code)
	msg := fmt.Sprintf("%v", err)
	if msg == resp.LoginCheckErrorMsg ||
		msg == resp.ForbiddenMsg ||
		msg == resp.UserLockedMsg ||
		msg == resp.UserDisabledMsg ||
		msg == resp.InvalidCaptchaMsg {
		ops.failWithMsg(msg)
		return
	}
	ops.failWithCodeAndMsg(resp.Unauthorized, msg)
}

func identity(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return claims[jwt.IdentityKey]
}
