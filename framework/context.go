package framework

import (
	"context"
	"encoding/json"
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


// #region response

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)

	//json.Marshal-->把stuct轉成json字串
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}

//#endregion





