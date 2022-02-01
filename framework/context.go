package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct{
	request *http.Request 
	responseWriter http.ResponseWriter
	ctx context.Context

	// 當前請求的 hander 鏈條
	handlers []ControllerHandler 
	index int // 當前請求調用到 調用鏈 的哪一個節點

	params map[string]string // url 路由匹配的參數

	//是否超時的標記
	hasTimeout bool
	//寫保護機制
	writerMux *sync.Mutex


}

//context newer
func NewContext(r *http.Request, w http.ResponseWriter) *Context{
	return &Context{
		request: r,
		responseWriter: w,
		ctx: r.Context(),
		writerMux: &sync.Mutex{},
		index: -1, //index 初始值為-1，因為每次調用都會自增1
	}
}

//#region  核心函數，調用context的下一個函數
func (ctx *Context) Next() error{
	ctx.index ++
	if ctx.index < len(ctx.handlers){
		if err := ctx.handlers[ctx.index](ctx); err!=nil{
			return err
		}
	}
	return nil
}

//#endregion


//#region base function

func (ctx *Context) GetRequest() *http.Request{
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter{
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout(){
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

// 為context設置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// 為context設置params
func (ctx *Context) setParams(params map[string]string){
	ctx.params = params
}

//#endregion

func (ctx *Context) BaseContext() context.Context{
	return ctx.request.Context()
}

// #region implement context.Context

func (ctx *Context) Deadline() (deadline time.Time, ok bool){
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{}{
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

//#endregion




