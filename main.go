package main

import (
	"net/http"
	"selfmade-webframework/framework"
	"selfmade-webframework/framework/middleware"
)

func main(){
	core := framework.NewCore()
	// core 中使用 use 註冊中間件
	core.Use( 
		middleware.Test1(), 
		middleware.Test2(),
	)
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}