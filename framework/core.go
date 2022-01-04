package framework

import "net/http"

//框架核心結構
type Core struct{

}

//初始化框架核心結構
//回傳Core引用類型
func NewCore() *Core{
	return &Core{}
}

//Core的鏈式函數宣告
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request){
	//TODO
}