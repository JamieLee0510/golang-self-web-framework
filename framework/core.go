package framework

import (
	"log"
	"net/http"
)

//框架核心結構
type Core struct{
	router map[string]ControllerHandler
}

//初始化框架核心結構
//回傳Core引用類型
func NewCore() *Core{
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

//Core的鏈式函數宣告
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	/// 一個簡單的路由選擇器，這邊直接寫死測試路由foo
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	router(ctx)
}