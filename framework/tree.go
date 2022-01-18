package framework

import (
	"errors"
	"strings"
)


type Tree struct{
	root *node //根節點
}

type node struct{
	isLast bool  //是否為葉子節點，代表是否為最終路由規則
	segment string // 該路由的某段字符串
	handlers []ControllerHandler // 這個節點中包含的handler，用於最終加載
								//改為hander array--> 中間件+控制器
	childs []*node // 這個節點下的子節點
}


func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root}
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
		// 如果此 segment 是最後一個節點，則判斷這些 cnode 是否有 isLast 的 flag
		for _, tn := range cnodes{
			if tn.isLast{
				return tn
			}
		}

		//假如都不是最後一個節點
		return nil
	}

	// 如果有2個segment，遞迴每個子節點繼續向下查找
	for _, tn := range cnodes{
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil{
			return tnMatch
		}
	}

	return nil
}


// 增加路由節點
func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error{
	n := tree.root

	//確認路由是否衝突
	if n.matchNode(uri) != nil{
		return errors.New("router is exist: "+ uri)
	}

	segments := strings.Split(uri, "/")

	// 對每一個segment
	for index, segemt := range segments{

		//最終進入segment的字段
		if !isWildSegmemt(segemt){
			segemt = strings.ToUpper(segemt)
		}

		isLast := index==len(segments) -1
		
		var objNode *node //標記是否有合適的子節點

		childNodes := n.filterChildNodes(segemt)

		//假如有匹配的子節點
		if len(childNodes)>0{
			// 假如有和 segment 相同的子節點，則選擇該節點
			for _, cnode := range childNodes{
				if cnode.segment == segemt{
					objNode = cnode
					break
				}
			}
		}

		//假如都沒有匹配到
		if objNode == nil{
			//創建一個當前的節點
			cnode := newNode()
			cnode.segment = segemt

			if isLast{
				cnode.isLast = true
				cnode.handlers = handlers
			}

			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

		// ?? 為啥這邊還要把 n 改為 objNode
		n = objNode
	}
	return nil
}


// 匹配uri
func (tree *Tree) FindHandler(uri string) []ControllerHandler{

	//直接複用 matchNode函數
	matchNode := tree.root.matchNode(uri)

	if matchNode == nil{
		return nil
	}

	return matchNode.handlers

}


