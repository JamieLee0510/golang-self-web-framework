package main

import "selfmade-webframework/framework"

func UserLoginController(c *framework.Context) error{
	// 打印控制器名字 
	c.Json(200, "ok, UserLoginController") 
	return nil
}