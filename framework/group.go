package framework

// IGroup 代表前缀分组 接口
type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
  }

// Group struct 实现了IGroup
type Group struct { 
	core *Core 
	prefix string
}

// 初始化 Group
func NewGroup(core *Core, prefix string) *Group{
	return &Group{
		core: core,
		prefix: prefix,
	}
}

// 實現Get方法
func (g *Group) Get(uri string, handler ControllerHandler) { 
	uri = g.prefix + uri 
	g.core.Get(uri, handler)
}

// 實現Post方法
func (g *Group) Post(uri string, handler ControllerHandler) { 
	uri = g.prefix + uri 
	g.core.Post(uri, handler)
}

// 實現Put方法
func (g *Group) Put(uri string, handler ControllerHandler) { 
	uri = g.prefix + uri 
	g.core.Put(uri, handler)
}

// 實現Delete方法
func (g *Group) Delete(uri string, handler ControllerHandler) { 
	uri = g.prefix + uri 
	g.core.Delete(uri, handler)
}

// 从core中初始化這個Group
/// 使用 IGroup 接口後，core.Group 這個方法返回的是一個約定，
/// 而不依賴具體的 Group 實現
func (c *Core) Group(prefix string) IGroup { 
	return NewGroup(c, prefix)
}
