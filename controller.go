package main

import (
	"selfmade-webframework/framework"
)

func FooControllerHandler(ctx *framework.Context) error {
	return ctx.Json(200, map[string]interface{}{
	  "code": 0,
	})
  }