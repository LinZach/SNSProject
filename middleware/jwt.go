package middleware

import (
	"SNSProject/session"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"reflect"
	"time"
)

func CheckTokenMiddleware(ctx *fasthttp.RequestCtx) error {
	path := string(ctx.Path())
	allows := map[string]bool{
		"/user/register": true,
		"/login":	true,
		"/post/getAll":	true,
		"/post/class":	true,
		"/post/commend": true,
		"/post/comment": true,
	}

	if _, ok := allows[path]; ok {
		return nil
	}

	tokenStr := ctx.Request.Header.Peek("AuthToken")
	claims, err := ParesToken(string(tokenStr), []byte(SecrtKey))
	if err != nil {
		return errors.New("token 解析失败")
	}

	uid := claims.(jwt.MapClaims)["uid"]
	value := reflect.ValueOf(uid)
	if value.Kind() == reflect.String {
		locToken, err := session.Get(value.String())
		if err != nil {
			return err
		}

		if locToken == string(tokenStr) {
			return nil
		}
	}
	
	return nil
}


type jwtToken struct {
	jwt.StandardClaims

	Uid   uint `json:"uid"`
	Admin bool `json:"admin"`
}

const SecrtKey  = "cudishcuihef728h7fhhiue"
const issuer  = "sky"

func CreatToken(secretKey []byte,uid uint) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		uid,
		true,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	return
}

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid   uint `json:"uid"`
	Admin bool `json:"admin"`
}

func ParesToken(tokenString string, secrtKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return secrtKey, err
	})

	claims = token.Claims

	if !token.Valid {
		return
	}

	return
}