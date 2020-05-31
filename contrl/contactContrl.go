package contrl

import (
	"DemoMall/helper"
	"DemoMall/model"
	"DemoMall/validatar"
	"github.com/valyala/fasthttp"
)

type contact struct {
	UserName string `name:"username" rule:"plainText" empty:"false" min:"2" max:"10" msg:"昵称长度在[2-10]文字内"`
	Avatar string `name:"avatar" rule:"url" empty:"false" msg:"请传入正确url"`
	Uid int `uid:"uid" rule:"digit" empty:"false" min:"1" msg:"用户id无法识别"`
}

type ContactContrl struct {}

//添加好友
func (that *ContactContrl)AddToContactList(ctx *fasthttp.RequestCtx)  {
	var contact contact
	err := validatar.Bind(ctx, &contact)
	if err != nil {
		helper.Print(ctx, "500", err)
	}

	uid := getUidWithCurrentTokenString(ctx)

	body := model.ContactBody{
		Uid:int32(contact.Uid),
		Avatar:contact.Avatar,
		UserName:contact.UserName,
	}

	err = model.AddContact(int32(uid), body)
	if err != nil {
		helper.Print(ctx, "400", err.Error())
	}else {
		helper.Print(ctx, "0", "好友添加成功")
	}
}

//获取好友列表
func (that *ContactContrl)GetContactList(ctx *fasthttp.RequestCtx)  {
	uid := getUidWithCurrentTokenString(ctx)
	contacts, err := model.QueryContactList(int32(uid))

	if err != nil {
		helper.Print(ctx, "400", err.Error())
	}else {
		helper.Print(ctx, "0", contacts)
	}
}