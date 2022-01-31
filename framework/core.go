package framework

import (
	"log"
	"net/http"
	"strings"
)

//框架核心結構
type Core struct{
	router map[string]*Tree // all routers
	middlewares []ControllerHandler // 從core這邊設置的中間件
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


// 匹配路由，如果沒有匹配到路由，則返回 nil
func (c *Core) FindRouteNodeByRequest(request *http.Request) *node{
	// uri 和 method 全部轉換成大寫，避免大小寫敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	//查找第一層 map
	if methodHandlers, ok := c.router[upperMethod]; ok{
		return methodHandlers.root.matchNode(uri)
	}

	return nil

}

//所有請求都進入這個函數,這個函數負責路由分發
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	log.Println("core.serveHTTP")
	// 封裝自定義的context
	ctx := NewContext(request, response)

	// 尋找路由
	node := c.FindRouteNodeByRequest(request)
	if node == nil{
		// 如果没有找到，print log
		ctx.Json(404, "not found") 
		return 
	}
	ctx.SetHandlers(node.handlers)

	// 設置路由參數
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.setParams(params)


	// 調用路函數，如果返回err 代表存在內部c錯誤，返回 500 status
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}

// 註冊中間件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}



// #region 路由註冊

// 對應 Method = Get
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	// 將core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil { 
		log.Fatal("add router error: ", err) 
	}
  }
  
  // 對應 Method = POST
  func (c *Core) Post(url string,  handlers ...ControllerHandler) {
	// 將core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil { 
		log.Fatal("add router error: ", err) 
	}
  }
  
  // 對應 Method = PUT
  func (c *Core) Put(url string, handlers ...ControllerHandler) {
	// 將core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil { 
		log.Fatal("add router error: ", err) 
	}
  }
  
  // 對應 Method = DELETE
  func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	// 將core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil { 
		log.Fatal("add router error: ", err) 
	}
  }

// #endregion

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

