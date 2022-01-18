package framework

import (
	"log"
	"net/http"
	"strings"
)

//框架核心結構
type Core struct{
	router map[string]*Tree // all routers
}

//初始化框架核心結構
//回傳Core引用類型
func NewCore() *Core{
	//初始化路由
	router := map[string]*Tree{} 
	router["GET"] = NewTree() 
	router["POST"] = NewTree() 
	router["PUT"] = NewTree() 
	router["DELETE"] = NewTree() 

	return &Core{router: router}
}

//所有請求都進入這個函數,這個函數負責路由分發
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	log.Println("core.serveHTTP")
	// 封裝自定義的context
	ctx := NewContext(request, response)

	// 尋找路由
	handlers := c.FindRouteByRequest(request)
	if handlers == nil { 
		// 如果没有找到，这里打印日志 
		ctx.Json(404, "not found") 
		return 
	}
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
	if err := c.router["GET"].AddRouter(url, handler); err != nil { 
		log.Fatal("add router error: ", err) 
	}
  }
  
  // 對應 Method = POST
  func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil { 
		log.Fatal("add router error: ", err) 
	}
  }
  
  // 對應 Method = PUT
  func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil { 
		log.Fatal("add router error: ", err) 
	}
  }
  
  // 對應 Method = DELETE
  func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil { 
		log.Fatal("add router error: ", err) 
	}
  }

// #endregion




// 匹配路由，如果沒有匹配到，則return nil
func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {

	uri := request.URL.Path
	method := request.Method

	// method 全部轉換為大寫
	upperMethod := strings.ToUpper(method)

  
	// 查找第一層map
	if methodHandlers, ok := c.router[upperMethod]; ok{
		return methodHandlers.FindHandler(uri)
	}
	
	return nil
  }