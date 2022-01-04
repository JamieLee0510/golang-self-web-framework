package main

import (
	"net/http"
	"selfmade-webframework/framework"
)

func main(){
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr: ":8080",
	}
	server.ListenAndServe()
}