package demo

import (
	"fmt"
	"selfmade-webframework/framework"
)

// 具體的結構實例
type DemoService struct {
	// 實現接口
	Service
  
	  // 参数
	c framework.Container
  }
  
  // 實現接口
  func (s *DemoService) GetFoo() Foo {
	return Foo{
	  Name: "i am foo",
	}
  }


// 初始化實例的方法
func NewDemoService(params ...interface{}) (interface{}, error) {
	// 這裡需要將參數展開
	c := params[0].(framework.Container)
  
	fmt.Println("new demo service")
	// 返回實例
	return &DemoService{c: c}, nil
  }