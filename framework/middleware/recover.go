package middleware

import "selfmade-webframework/framework"

func Recover() framework.ControllerHandler{
	//使用函數回調
	return func (c *framework.Context)  error{
		// 核心在增加這個 recover 機制，捕獲 c.Next()出現的panic 
		defer func() { 
			if err := recover(); err != nil { 
				c.Json(500, err) 
			} 
		}()

		c.Next()

		return nil
	}
}