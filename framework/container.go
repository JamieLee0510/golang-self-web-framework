/// 服務容器

package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Container 是一個服務容器，提供綁定服務和獲取服務的功能
type Container interface{
	// Bind用來註冊 serviceProvider，return error 代表是否成功
	Bind(provider ServiceProvider) error
	// IsBind
	IsBind(key string) bool

	// Make 根據 key 來獲取在Container中已經實例化的 serviceProvider
	Make(key string) (interface{}, error) 
	// MustMake 就像Make，但不會返回error，所以發生錯誤會panic
	// 所以在使用MustMake時要確保Container已經綁定這個key的serviceProvider
	MustMake(key string) interface{}
	// MakeNew根據key獲取serviceProvider，但這個serviceProvider不是單例模式
	// 它是根據serviceProvider 註冊的啟動函數和params來實例化
	// 因此，MakeNew在需要不同參數來實例化不同serviceProvider的場景好用
	MakeNew(key string, params []interface{}) (interface{}, error) 
}

// DingContainer是服務容器的具體實現
type DingContainer struct{
	Container //強制要求DingContainer直接實現Container接口

	//providers 存註冊的ServiceProvider，key為字符串憑證
	providers map[string]ServiceProvider

	// instance 存具體的實例，key為字符串憑證
	instances map[string]interface{}

	// lock用於鎖住對容器的變更操作
	lock sync.RWMutex

}

// NewDineContainer 創建一個服務容器
func NewDingContainer() *DingContainer {
	return &DingContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 輸出服務容器所註冊的key
func (dingContainer *DingContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range dingContainer.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

//Bind 將key和服務容器作綁定
func (dingContainer *DingContainer)Bind(provider ServiceProvider)error{
	dingContainer.lock.Lock()
	defer dingContainer.lock.Unlock()

	key := provider.Name()
	dingContainer.providers[key] = provider
	// IsDefer為false時，立即註冊不要延遲
	if !provider.IsDefer(){
		if err := provider.Boot(dingContainer); err != nil { 
			return err
		 }

		// 實例化方法
		params := provider.Params(dingContainer)
		method := provider.Register(dingContainer)
		instance, err := method(params...)
		if err != nil { 
			return errors.New(err.Error()) 
		} 

		//Container 存 serviceProvider實例
		dingContainer.instances[key] = instance
	}

	return nil
}

func (dingContainer *DingContainer) IsBind(key string) bool {
	return dingContainer.findServiceProvider(key) != nil
}

// findServiceProvider為查詢是否有註冊過該key的serviceProvider
func (dingContainer *DingContainer) findServiceProvider(key string) ServiceProvider {
	dingContainer.lock.RLock()
	defer dingContainer.lock.RUnlock()
	if sp, ok := dingContainer.providers[key]; ok {
		return sp
	}
	return nil
}

func (dingContainer *DingContainer) Make(key string) (interface{}, error) {
	return dingContainer.make(key, nil, false)
}

func (dingContainer *DingContainer) MustMake(key string) interface{} {
	serv, err := dingContainer.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (dingContainer *DingContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return dingContainer.make(key, params, true)
}

// newInstance進行 serciveProvider 實例化
func (dingContainer *DingContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	
	if err := sp.Boot(dingContainer); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(dingContainer)
	}
	method := sp.Register(dingContainer)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}


// 真正的實例化一個服務
func (dingContainer *DingContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	dingContainer.lock.RLock()
	defer dingContainer.lock.RUnlock()
	// 查詢是否已經註冊了該serviceProvider，如果有則返回errors
	sp := dingContainer.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return dingContainer.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := dingContainer.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := dingContainer.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	dingContainer.instances[key] = inst
	return inst, nil
}
