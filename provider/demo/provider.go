package demo

import (
	"fmt"
	"selfmade-webframework/framework"
)

// 服務提供方
type DemoServiceProvider struct {
}

// Name func直接將服務對應的key返回，在demo的例子就是“hi:demo”
func (sp *DemoServiceProvider) Name() string {
  return Key
}

// Register 是註冊初始化服務實例的方法，這裏先暫定為 NewDemoservice
func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
  return NewDemoService
}

// IsDefer 表示是否延遲實例化，先暫定為true，將這個服務的實例化延遲到第一次make的時候
func (sp *DemoServiceProvider) IsDefer() bool {
  return true
}

// Params 為實例化的參數。我們這裡只實例化一個參數：container
// 表示我们在 NewDemoService 這個函數中只有一個參數---container
func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
  return []interface{}{c}
}

// Boot 先什麼都不執行, 只打印一行log
func (sp *DemoServiceProvider) Boot(c framework.Container) error {
  fmt.Println("demo service boot")
  return nil
}