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

func (ctx *Context) BaseContext() context.Context{
	return ctx.request.Context()
}

func (ctx *Context) Done() <-chan struct{}{
	return ctx.BaseContext().Done()
}
