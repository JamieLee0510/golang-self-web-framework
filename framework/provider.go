/// 主體框架的服務容器(service container)
/// 服務容器綁定多個服務接口協議，每個接口協議都由一個service provider
/// 提供服務

package framework

// NewInstance 定義如何創建新實例，這是所有服務容器的創建服務
type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider 定義一個service provider需要實現的接口
type ServiceProvider interface {
	// Register 在service container中註冊一個實例化服務的方法
	// 是否在註冊時就實例化這個服務，需要参考IsDefer接口
	Register(Container) NewInstance

	// Boot 在調用實例化服務的時候會調用，可以把一些準備工作：基礎配置、初始化參數放這裡。
	// 如果Boot返回error，整個服務實例化就會失敗，返回錯誤
	Boot(Container) error

	// IsDefer 決定是否在註冊時就實例化這個服務；
	// **如果不是註冊時就實例化，那就是在第一次make的時候進行
	// false 表示不需要延遲實例化、在註冊的時候就實例化 / true表延遲實例化
	IsDefer() bool

	// Params 定義傳給NewInstance的參數，可以自定義多個
	// 建議將container作為第一個參數
	Params(Container) []interface{}

	// Name 代表這個service provider的憑證
	Name() string
}

