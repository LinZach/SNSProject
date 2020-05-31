package roter

import (
	"DemoMall/contrl"
)
import "github.com/buaazp/fasthttprouter"

func InitRouter() *fasthttprouter.Router {

	user := contrl.UserContrl{}
	post := contrl.Postcontrl{}

	router := fasthttprouter.New()

	router.POST("/getSlat", user.GetSlatHandle)
	router.POST("/user/register", user.RegisterHandle)
	router.POST("/login", user.Login)

	router.POST("/post/push", post.PushPost)
	router.POST("/post/getAll", post.GetPostList)
	router.POST("/post/class", post.GetPostClass)
	router.POST("/post/commend", post.PostCommend)
	router.POST("/post/comment", post.PostComment)
	router.POST("/post/getComment", post.GetPostComents)

	return router
}
