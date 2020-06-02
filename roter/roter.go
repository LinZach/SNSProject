package roter

import (
	"SNSProject/contrl"
	"github.com/buaazp/fasthttprouter"
)

func InitRouter() *fasthttprouter.Router {

	user := contrl.UserContrl{}
	post := contrl.Postcontrl{}
	contact := contrl.ContactContrl{}

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

	router.POST("/contact/add", contact.AddToContactList)
	router.POST("/contact/list", contact.GetContactList)

	return router
}
