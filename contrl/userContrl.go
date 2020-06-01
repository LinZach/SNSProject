package contrl

import (
	"SNSProject/Crypt"
	"SNSProject/helper"
	"SNSProject/middleware"
	"SNSProject/model"
	"SNSProject/validatar"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"reflect"
)

type UserContrl struct {

}

type RegistParam struct {
	Account string `name:"account" rule:"digit" empty:"false" min:"8" max:"999999999999999" msg:"用户昵称[2-30]位英文与数字"`
	Password string `name:"password"`
	Username string `name:"username" rule:"plainText" empty:"false" min:"2" max:"10" msg:"昵称长短在,[2-10]文字内"`
}

type LoginRsp struct {
	Token string `json:"token"`
}

type userProfile struct {
	UserName string `name:"username" rule:"plainText" empty:"false" min:"2" max:"10" msg:"昵称长度在[2-10]文字内"`
	Avatar string `name:"avatar" rule:"url" empty:"false" msg:"请传入正确url"`
	Slogan string `name:"slogan" rule:"plainText" empty:"false" min:"1" max:"33" msg:"个人签名长度在[1-33]字内"`
	Gender int `name:"gender" rule:"digit" empty:"false" min:"1" max:"1" msg:"请检查参数长度"`
	Uid int `uid:"uid" rule:"digit" empty:"false" min:"1" msg:"用户id无法识别"`
}

//获取盐值
func (that *UserContrl)GetSlatHandle(ctx *fasthttp.RequestCtx)  {
	helper.Print(ctx, "0", Crypt.Salt)
}

//注册
func (that *UserContrl)RegisterHandle(ctx *fasthttp.RequestCtx) {
	var regParam RegistParam
	err := validatar.Bind(ctx, &regParam)
	if err != nil {
		helper.Print(ctx, "500", err.Error())
		return
	}

	user := model.User{
		Username:regParam.Username,
		Account:regParam.Account,
		Uid:0,
		Password:regParam.Password,
	}

	err = model.Insert(user)
	if err != nil {
		helper.Print(ctx, "500", err.Error())
	}
}

//登录
func (that *UserContrl)Login(ctx *fasthttp.RequestCtx) {
	var regParam RegistParam
	err := validatar.Bind(ctx, &regParam)
	if err != nil {
		helper.Print(ctx, "500", err.Error())
		return
	}

	user := model.User{
		Username:regParam.Username,
		Account:regParam.Account,
		Uid:0,
		Password:regParam.Password,
	}

	if !validataUser(user) {
		helper.Print(ctx, "0", "账号不存在")
		return
	}

	if !model.ValidataUser(&user) {
		helper.Print(ctx, "500", "密码错误")
		return
	}

	var uid = user.Uid
	var dd = uint(uid)
	tokenStr, err := middleware.CreatToken([]byte(middleware.SecrtKey),dd)
	if err != nil {
		fmt.Print(err)
		return
	}
	err = model.SetToken(user.Uid, tokenStr)
	if err != nil {
		fmt.Print(err)
		return
	}

	data := LoginRsp{Token:tokenStr}

	helper.Print(ctx, "0", data)
}

//修改用户资料
func (that *UserContrl)UpdateUserProfile(ctx *fasthttp.RequestCtx)  {
	uid := getUidWithCurrentTokenString(ctx)

	var usrPro userProfile
	err := validatar.Bind(ctx, &usrPro)
	if err != nil {
		helper.Print(ctx, "500", err.Error())
		return
	}

	var user = model.User{
		Username:usrPro.UserName,
		Gender:usrPro.Gender,
		Avatar:usrPro.Avatar,
		Slogan:usrPro.Slogan,
		Uid:int32(uid),
	}

	excit := validataUserWithUid(user)
	if !excit {
		helper.Print(ctx, "400", "用户不存在")
	}

	err = model.UpdateUserProfile(user)
	if err == nil {
		helper.Print(ctx, "0", "修改成功")
	}else {
		helper.Print(ctx, "400", "信息修改失败")
	}
}

//验证用户是否存在
func validataUser(user model.User) bool {
	var isExcat = false
	if model.Query(user) > 0 {
		isExcat = true
	}
	return isExcat
}

//验证用户是否存在 使用uid
func validataUserWithUid(user model.User) bool {
	var isExcat = false
	if model.QueryWithUid(user) > 0 {
		isExcat = true
	}
	return isExcat
}

//解析token 获得 uid
func getUidWithCurrentTokenString(ctx *fasthttp.RequestCtx) int {
	tokenStr := ctx.Request.Header.Peek("AuthToken")
	if tokenStr == nil {
		return 0
	}

	claims, err := middleware.ParesToken(string(tokenStr), []byte(middleware.SecrtKey))
	if err != nil {
		return 0
	}
	uid := claims.(jwt.MapClaims)["uid"]
	value := reflect.ValueOf(uid)
	return int(value.Int())
}