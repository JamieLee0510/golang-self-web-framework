// Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"
	"fmt"
	"selfmade-webframework/framework"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// engine實現container的綁定封裝
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	fmt.Print("綁定serviceProvider成功！")
	return engine.container.Bind(provider)
}

// IsBind key是否已經綁定serviceProvider
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// context 實現container的幾個封裝

// 實現Make的封裝
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

// 實現MustMake的封裝
func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

// 實現MakeNew的封裝
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}
