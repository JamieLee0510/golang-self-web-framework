package middleware

import (
	"fmt"
	"selfmade-webframework/framework"
)

/// 測試用的中間件

func Test1() framework.ControllerHandler {
	// 使用函數回調
	return func(c *framework.Context) error {
	  fmt.Println("middleware pre test1")
	  c.Next()  // 調用Next()往下調用，讓contxt.index自增
	  fmt.Println("middleware post test1")
	  return nil
	}
  }
  
  func Test2() framework.ControllerHandler {
	// 使用函數回調
	return func(c *framework.Context) error {
	  fmt.Println("middleware pre test2")
	  c.Next() // 調用Next()往下調用，讓contxt.index自增
	  fmt.Println("middleware post test2")
	  return nil
	}
  }
  