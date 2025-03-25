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

// @secretKey JWT 加解密密钥
// @iat	时间戳
// @exp	过期时间[秒]
// @payload数据载体
func GetJwtToken(secretKey string, iat, exp int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + exp
	claims["iat"] = iat
	claims[TOKEN_PAYLOAD] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func HttpResult[T any](w http.ResponseWriter, r http.Request, data T, err error) {

	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, Success(data))
	}
}

func HandlerAction(w http.ResponseWriter, r *http.Request, err error, resp any) {

	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

func LogicAction(ok *Result, fail *Result, err error) (*Result, error) {
	if err != nil {
		return fail, err
	}
	return ok, nil
}
