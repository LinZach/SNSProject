package main

import (
	"SNSProject/DB"
	"SNSProject/middleware"
	"SNSProject/roter"
	"SNSProject/session"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func main() {

	DB.InitDB()
	session.New("127.0.0.1:6379")
	//model.Insert()
	//
	//
	router := roter.InitRouter()

	server := fasthttp.Server{
		Name:								"snsApi",
		Handler:                            middleware.User(router.Handler),
		ReadTimeout:                        30 * time.Second,
		WriteTimeout:                       30 * time.Second,
		MaxRequestBodySize:                 5 * 1024 * 1024,
	}

	if err := server.ListenAndServe("127.0.0.1:8080"); err != nil {
		fmt.Print(err)
	}

	//if err := fasthttp.ListenAndServe(":9096", middleware.Auth()); err != nil {
	//	fmt.Print("conecting auth fail:", err.Error())
	//}
}
