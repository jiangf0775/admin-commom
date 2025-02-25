package web

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetLoginJwtToken(secretKey string, iat, exp int64, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + exp
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetJwtTokenFunc(secretKey string, iat, exp int64, setting func(claims jwt.MapClaims)) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + exp
	claims["iat"] = iat
	setting(claims)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func HttpResult[T any](w http.ResponseWriter, r http.Request, data T, err error) {

	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, Success[T](data))
	}
}

func LogicAction(w http.ResponseWriter, r *http.Request, err error, resp any) {

	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
