package middleware

import (
	"DemoMall/helper"
	"DemoMall/model"
	"DemoMall/validatar"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"github.com/wyanlord/oauth2/errors"
	"github.com/wyanlord/oauth2/generates"
	"github.com/wyanlord/oauth2/manage"
	"github.com/wyanlord/oauth2/models"
	"github.com/wyanlord/oauth2/server"
	"github.com/wyanlord/oauth2/store"
)

//验证用户登录信息
type Register struct {
	Account string `name:"account" rule:"alnum" empty:"false" min:"8" max:"10" msg:"用户昵称[2-30]位英文与数字"`
	Password string `name:"password" rule:"alnum" empty:"false" min:"8" max:"16" msg:"用户密码[2-30]位英文与数字"`
	Username string `name:"username" rule:"plainText" empty:"false" min:"2" max:"10" msg:"昵称长短在,[2-10]文字内"`
}

func Auth() fasthttp.RequestHandler {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	//token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("000000"), jwt.SigningMethodHS256))

	clientStore := store.NewClientStore()
	err := clientStore.Set("222222", &models.Client{
		ID:"222222",
		Secret:"222222",
		Domain:"http://www.baidu.com",
	})

	if err != nil {
		fmt.Print(err)
	}

	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetUserAuthorizationHandler(func(ctx *fasthttp.RequestCtx) (userID string, err error) {
		var register Register
		err = validatar.Bind(ctx, &register)
		if err != nil {
			helper.Print(ctx, "500", err.Error())
			return
		}

		user := model.User{
			Username:register.Username,
			Account:register.Account,
			Uid:0,
			Password:register.Password,
		}

		if !model.ValidataUser(&user) {
			helper.Print(ctx, "500", "密码错误")
			return
		}

		return string(user.Uid), err
	})

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		fmt.Print(err)
		return
	})
	
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		fmt.Print(re.Description)
	})

	h := fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/authorize":
			err := srv.HandleAuthorizeRequest(ctx)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				_, _ = ctx.WriteString(err.Error())
			}

		case "/token":
			_ = srv.HandleTokenRequest(ctx)

		default:
			ctx.SetStatusCode(fasthttp.StatusNotFound)

		}
	})

	return h
}
