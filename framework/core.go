package framework

import (
	"log"
	"net/http"
	"strings"
)

//框架核心結構
type Core struct{
	router map[string]map[string]ControllerHandler //二級map
}

//初始化框架核心結構
//回傳Core引用類型
func NewCore() *Core{
	//定義二級map
	getRouter := map[string]ControllerHandler{} 
	postRouter := map[string]ControllerHandler{} 
	putRouter := map[string]ControllerHandler{}
	 deleteRouter := map[string]ControllerHandler{}

	//將二級map寫入一級map
	 router := map[string]map[string]ControllerHandler{} 
	 router["GET"] = getRouter 
	 router["POST"] = postRouter 
	 router["PUT"] = putRouter 
	 router["DELETE"] = deleteRouter
	return &Core{router: router}
}

// #region 路由註冊

// 對應 Method = Get
func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
  }
  
  // 對應 Method = POST
  func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
  }
  
  // 對應 Method = PUT
  func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
  }
  
  // 對應 Method = DELETE
  func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
  }

// #endregion

//Core的鏈式函數宣告
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	// 一个简单的路由选择器，这里直接写死为测试路由foo
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	router(ctx)
}