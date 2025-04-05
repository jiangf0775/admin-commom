package web

import (
	"common/tool"
	"common/types"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
	"strings"
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

// @secretKey JWT 加解密密钥
// @iat	时间戳
// @exp	过期时间[秒]
// @payload数据载体
func GetJwtTokenByUser(secretKey string, iat, exp int64, user types.User) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + exp
	claims["iat"] = iat
	//拼接payload值,*为防止Name中含有分隔符，提前替换
	var strs = make([]string, 3)
	strs[USER_NAME] = strings.ReplaceAll(user.Name, TOKEN_PAYLOAD_SPLIT, " ")
	strs[USER_ID] = strconv.FormatUint(user.Id, 10)
	if len(user.RoleId) > 0 {
		strs[USER_ROLE_IDS] = strings.Join(tool.Uint64ToStrings(user.RoleId), TOKEN_PAYLOAD_SPLIT2)
	}
	claims[TOKEN_PAYLOAD] = strings.Join(strs, TOKEN_PAYLOAD_SPLIT)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetUserByJwt(r *http.Request) (*types.User, error) {
	payload := r.Context().Value(TOKEN_PAYLOAD)
	jsonPayload, ok := payload.(json.Number)
	if !ok {
		return nil, JwtGetUserInfoError
	}
	str := jsonPayload.String()
	if str == "" {
		return nil, JwtGetUserInfoError
	}
	split := strings.Split(str, TOKEN_PAYLOAD_SPLIT)
	var user types.User
	user.Name = split[USER_NAME]
	parseUint, err := strconv.ParseUint(split[USER_ID], 10, 64)
	if err != nil {
		return nil, JwtGetUserInfoError
	}
	user.Id = parseUint
	user.RoleId = []uint64{}
	roleStr := split[USER_ROLE_IDS]
	if roleStr != "" {
		roleIds := strings.Split(roleStr, TOKEN_PAYLOAD_SPLIT2)
		user.RoleId = tool.StringToUint64s(roleIds)
	}

	return &user, nil
}

func HandlerAction(w http.ResponseWriter, r *http.Request, err error, resp any) {

	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}

func HandlerActionJsonOK(w http.ResponseWriter, r *http.Request, resp *Result) {

	httpx.OkJsonCtx(r.Context(), w, resp)
}
