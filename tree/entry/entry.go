package main

import "learn_go/tree"

/**
 *
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-09-01 9:25<br/>
 * @version 2021.09
 * @since 2021.09
 */
func main() {
	//结构的创建
	tree.CrNode()

	//遍历
	tree.Bl()
}

//使用组合的方式扩展别人的类
/*type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil {
		return
	}
	myTreeNode{node: myNode.node.Left}
}*/
