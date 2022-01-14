package framework


type Tree struct{
	root *node //根節點
}

type node struct{
	isLast bool  //是否為葉子節點，代表是否為最終路由規則
	segment string // 該路由的某段字符串
	handler ControllerHandler // 這個節點中包含的handler，用於最終加載
	childs []*node // 這個節點下的字節點
}

