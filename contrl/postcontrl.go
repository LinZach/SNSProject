package contrl

import (
	"DemoMall/helper"
	"DemoMall/middleware"
	"DemoMall/model"
	"DemoMall/validatar"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"reflect"
	"strconv"
)

type Postcontrl struct {
}

//发布帖子请求实体化模型
type postArgs struct {
	Title string `name:"title" rule:"plainText" empty:"false" min:"1" max:"26" msg:"帖子标题[1-26]文字"`
	Uperid int32 `name:"uperid" rule:"digit" empty:"false"`
	Content string `name:"content" empty:"false"`
	Files string `name:"files"`
}

//点赞实体化请求
type postCommend struct {
	Add int `name:"add" empty:"false" min:"1" max:"1"`
	Pid int `name:"pid" rule:"digit" empty:"false"`
}

//评论请求体实体化
type postComment struct {
	Content string `name:"content" empty:"false" min:"1" max:"225" msg:"帖子标题[1-225]文字"`
	Pid int32 `name:"pid" rule:"digit" empty:"false"`
}

type commentsReq struct {
	Pid int32 `name:"pid" rule:"digit" empty:"false"`
	Index int `name:"index" rule:"digit" empty:"false"`
	size int `name:"size" rule:"digit" empty:"false"`
}

//发布帖子
func (that *Postcontrl)PushPost(ctx *fasthttp.RequestCtx)  {

	var postAg postArgs
	err := validatar.Bind(ctx, &postAg)
	if err != nil {
		fmt.Print(err)
	}

	post := model.Post{
		Title:postAg.Title,
		Uperid:postAg.Uperid,
		Content:postAg.Content,
		Files:postAg.Files,
	}

	err = model.PostUp(post)

	if err != nil {
		fmt.Print(err)
	}

	helper.Print(ctx, "0", "发表成功")
}

//查询帖子列表
func (that *Postcontrl)GetPostList(ctx *fasthttp.RequestCtx)  {

	list := model.QueryPostList(1, 10, getUidWithCurrentTokenString(ctx))

	helper.Print(ctx, "0", list)
}

//查询帖子分类
func (that *Postcontrl)GetPostClass(ctx *fasthttp.RequestCtx)  {
	class := model.QueryPostClass()

	helper.Print(ctx, "0", class)
}

//点赞
func (that *Postcontrl)PostCommend(ctx *fasthttp.RequestCtx)  {
	uid := getUidWithCurrentTokenString(ctx)
	if uid == "" {
		helper.Print(ctx, "400", "token解析失败.未登录,或token失效")
		return
	}

	var postCd postCommend
	err := validatar.Bind(ctx, &postCd)

	if err != nil {
		fmt.Print(err)
		helper.Print(ctx, "500", "请检查参数是否合法")
		return
	}

	intUId, _ := strconv.Atoi(uid)

	err = model.SetPostCommend(postCd.Add, int32(intUId), int32(postCd.Pid))
	if err == nil {
		helper.Print(ctx, "0", "success")
	} else {
		helper.Print(ctx, "500", "fail")
	}
}

//评论
func (that *Postcontrl)PostComment(ctx *fasthttp.RequestCtx)  {
	uid := getUidWithCurrentTokenString(ctx)
	intUid, _ := strconv.Atoi(uid)

	var cObject postComment
	err := validatar.Bind(ctx, &cObject)
	if err != nil {
		fmt.Print(err)
		helper.Print(ctx, "500","请检查参数是否合法")
		return
	}

	_, err = model.QueryPost(cObject.Pid)
	if err != nil {
		helper.Print(ctx, "450", "帖子不存在")
		return
	}

	commend := model.Comment{
		Content:cObject.Content,
		Uid:int32(intUid),
		Pid:cObject.Pid,
	}

	err = model.AddComment(commend)

	if err != nil {
		helper.Print(ctx, "500", "评论添加失败")
	}else {
		helper.Print(ctx, "0", "评论添加成功")
	}
}

//获得帖子下评论
func (that *Postcontrl)GetPostComents(ctx *fasthttp.RequestCtx)  {
	var cReq commentsReq
	err := validatar.Bind(ctx, &cReq)
	if err != nil {
		helper.Print(ctx, "500", "请检查参数是否错误")
		return
	}

	comments, err := model.QueryCommentWithPid(cReq.Pid, cReq.Index, cReq.size)
	if err != nil {
		helper.Print(ctx, "450", "为找到该帖子的评论")
		return
	}

	helper.Print(ctx, "0", comments)
}