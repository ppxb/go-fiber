package middleware

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	v4 "github.com/golang-jwt/jwt/v4"
	"github.com/ppxb/go-fiber/pkg/constant"
	"github.com/ppxb/go-fiber/pkg/log"
	"github.com/ppxb/go-fiber/pkg/req"
	"github.com/ppxb/go-fiber/pkg/resp"
	"github.com/ppxb/go-fiber/pkg/utils"
	"net/http"
	"strings"
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

func JwtLogin(options ...func(*JwtOptions)) gin.HandlerFunc {
	ops := getJwtOptions(nil)
	for _, f := range options {
		f(ops)
	}
	mw := initJwt(*ops)
	return func(c *gin.Context) {
		user, data, err := login(c, *ops)

		if err != nil {
			unauthorized(c, http.StatusUnauthorized, err, *ops)
			return
		}

		token := v4.New(v4.GetSigningMethod(mw.SigningAlgorithm))
		claims := token.Claims.(v4.MapClaims)

		for k, v := range payload(data) {
			claims[k] = v
		}

		expire := mw.TimeFunc().Add(mw.Timeout)
		claims["exp"] = expire.Unix()
		claims["orig_iat"] = mw.TimeFunc().Unix()
		tokenString, err := signedString(mw.Key, token)

		if err != nil {
			unauthorized(c, http.StatusUnauthorized, err, *ops)
			return
		}

		loginResponse(c, http.StatusOK, user, tokenString, expire, *ops)
	}
}

func initJwt(ops JwtOptions) *jwt.GinJWTMiddleware {
	j, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         ops.realm,
		Key:           []byte(ops.key),
		Timeout:       time.Hour * time.Duration(ops.timeout),
		MaxRefresh:    time.Hour * time.Duration(ops.maxRefresh),
		TokenLookup:   ops.tokenLookup,
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

func payload(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		return jwt.MapClaims{
			jwt.IdentityKey:                  v[constant.MiddlewareJwtUserCtxKey],
			constant.MiddlewareJwtUserCtxKey: v[constant.MiddlewareJwtUserCtxKey],
		}
	}
	return jwt.MapClaims{}
}

func login(c *gin.Context, ops JwtOptions) (interface{}, interface{}, error) {
	var r req.LoginCheck
	c.ShouldBind(&r)
	r.Username = strings.TrimSpace(r.Username)
	r.Password = strings.TrimSpace(r.Password)

	user, err := ops.loginPwdCheck(c, r)
	if err != nil {
		return nil, nil, err
	}
	return user, map[string]interface{}{
		constant.MiddlewareJwtUserCtxKey: fmt.Sprintf("%d", user.Id),
	}, nil
}

func signedString(key []byte, token *v4.Token) (tokenString string, err error) {
	tokenString, err = token.SignedString(key)
	return
}

func loginResponse(c *gin.Context, ok int, user interface{}, tokenString string, expire time.Time, ops JwtOptions) {
	resp.SuccessWithData(c, map[string]interface{}{
		"user":    user,
		"token":   tokenString,
		"expires": expire,
	})
}
