package framework

import "strings"


type Tree struct{
	root *node //根節點
}

type node struct{
	isLast bool  //是否為葉子節點，代表是否為最終路由規則
	segment string // 該路由的某段字符串
	handler ControllerHandler // 這個節點中包含的handler，用於最終加載
	childs []*node // 這個節點下的子節點
}

// 判斷一個 segment 是否為通用的 segement
func isWildSegmemt(segment string) bool{
	return strings.HasPrefix(segment,":")
}

// 過濾下一層滿足 segment 規則的子節點 
func (n *node) filterChildNodes(segment string) []*node{
	if len(n.childs)== 0{
		return nil
	}

	if isWildSegmemt(segment){
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs)) // so far len(nodes)=0, cap(nodes)=len(n.childs)

	for _, cnode := range n.childs{
		if isWildSegmemt(cnode.segment){
			// 如果下一層子節點有通配符，則滿足需求
			nodes = append(nodes, cnode)
		}else if cnode.segment == segment{
			// 如果下一層子節點沒有通配符，但文本完全匹配
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

// 判斷路由是否存在於現在的 trie tree 中
func (n *node) matchNode(uri string) *node{
	// 使用 “/” 將 uri 切割為一個長度為2的list 
	segments := strings.SplitN(uri, "/", 2) 

	// list[0]用於匹配下一個子節點
	segment := segments[0] 
	if !isWildSegmemt(segment) { 
		segment = strings.ToUpper(segment) 
	}
	// 匹配符合的下一個子節點
	cnodes := n.filterChildNodes(segment)  
	// 如果當前子節點沒有一個符合，
	// 代表這個 uri 一定不存在於 trie tree 中，
	// 直接 retur nil 
	if cnodes == nil || len(cnodes) == 0 {    
		return nil  
	}

	// 如果 segments len == 1,代表為最後一個標記
	if len(segments)==1{
		
	}
}
