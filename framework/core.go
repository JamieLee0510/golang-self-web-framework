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

//Core的鏈式函數宣告
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	// 尋找路由
	router := c.FindRouteByRequest(request) 
	if router == nil { 
		// 如果沒有找到，log 
		ctx.Json(404, "not found") 
		return 
	}

	// 調用路由函數，如果返回err 代表內部有錯，返回 status 500
	if err := router(ctx); err != nil { 
		ctx.Json(500, "inner error") 
		return 
	}
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



// 匹配路由，如果沒有匹配到，則return nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {

	uri := request.URL.Path
	method := request.Method

	// uri 和 method 全部轉換為大寫
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(uri)
  
	// 查找第一層map
	if methodHandlers, ok := c.router[upperMethod]; ok {
	  // 查找第二層map
	  if handler, ok := methodHandlers[upperUri]; ok {
		return handler
	  }
	}
	return nil
  }