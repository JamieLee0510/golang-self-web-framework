package middleware

import (
	"fmt"
	"selfmade-webframework/framework/gin"
)

/// 測試用的中間件

func Test1()  gin.HandlerFunc {
	// 使用函數回調
return func(c *gin.Context) {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
	}
  }
  
  func Test2()  gin.HandlerFunc {
	// 使用函數回調
	return func(c *gin.Context)  {
	  fmt.Println("middleware pre test2")
	  c.Next() // 調用Next()往下調用，讓contxt.index自增
	  fmt.Println("middleware post test2")

	}
  }

  func Test3()  gin.HandlerFunc {
	// 使用函數回調
	return func(c *gin.Context) {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
	}
  }
  
  