/// 為接口說明文件，定義這個服務的key、設計這個服務的接口
/// 這裏了 demo.Service 接口，它有一個 GetFoo 方法，返回了 Foo 的數據結構。

package demo

// Demo 服務的 key
const Key = "hi:demo"

// Demo 服務的接口
type Service interface {
  GetFoo() Foo
}

// Demo 服務接口定義的一個數據結構
type Foo struct {
  Name string
}