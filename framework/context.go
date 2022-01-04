package framework

import (
	"context"
	"net/http"
	"sync"
)

type Context struct{
	request *http.Request 
	responseWriter http.ResponseWriter
	ctx context.Context
	handler ControllerHandler

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
	}
}

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

func (ctx *Context) BaseContext() context.Context{
	return ctx.request.Context()
}

//#endregion

func (ctx *Context) Done() <-chan struct{}{
	return ctx.BaseContext().Done()
}
