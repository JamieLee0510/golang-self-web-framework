package framework

// IGroup 代表前缀分组 接口
type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	//實現嵌套group
	Group(string) IGroup

	// Use()
  }

// Group struct 实现了IGroup
type Group struct { 
	core *Core 
	prefix string
	parent *Group //指向上一個Group，如果有的話
	middlewares []ControllerHandler
}

// 初始化 Group
func NewGroup(core *Core, prefix string) *Group{
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
		middlewares: []ControllerHandler{},
	}
}


// 註冊中間件
func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
 }

// 獲取某個 group 的 middlewares
// 這裏就是獲取除了 Get/Post/Put/Delete 之外設置的 middleware
func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}

	return append(g.parent.getMiddlewares(), g.middlewares...)
}

// 實現Get方法
func (g *Group) Get(uri string, handlers ...ControllerHandler) { 
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(uri, allHandlers...)
}

// 實現Post方法
func (g *Group) Post(uri string, handlers ...ControllerHandler) { 
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Post(uri, allHandlers...)
}

// 實現Put方法
func (g *Group) Put(uri string, handlers ...ControllerHandler) { 
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Put(uri, allHandlers...)
}

// 實現Delete方法
func (g *Group) Delete(uri string, handlers ...ControllerHandler) { 
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Delete(uri, allHandlers...)
}

// 獲取當前group的絕對路徑（遞迴）
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// 从core中初始化這個Group
/// 使用 IGroup 接口後，core.Group 這個方法返回的是一個約定，
/// 而不依賴具體的 Group 實現
func (g *Group) Group(uri string) IGroup { 
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}
