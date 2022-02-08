package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"selfmade-webframework/framework/framework/gin"
	"syscall"
)

func main(){
	core := gin.New()
	// core 中使用 use 註冊中間件
	core.Use(gin.Recovery())
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	// 這個goroutine是啟動服務到goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 當前的Goroutine等待信號量
	quit := make(chan os.Signal)
	// 監控信號：SIGINT, SIGTERM, SIGQUIT 
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 這裡會阻塞，等待信號量。使用channel導出操作，來阻塞當前的Goroutine
	<- quit

	//調用 server.Shutdown() graceful 結束
	if err := server.Shutdown(context.Background()); err != nil { 
		log.Fatal("Server Shutdown:", err) 
	}
	
}