package main

import (
	"fmt"
)

/**
 *
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-31 9:32<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	//结构的创建
	crNode()

	//遍历
	bl()
}

//遍历
func bl() {
	var root treeNode
	root = treeNode{value: 5}
	root.print() //5

	//root.right.left.setValue(4) //panic: runtime error: invalid memory address or nil pointer dereference没有赋予地址,要使用必须实例化
	root.right = new(treeNode)
	//使用局部变量
	root.right.left = createNode(3)
	root.right.left.setValue(4)
	root.right.left.print()

	var pRoot *treeNode
	pRoot.setValue(200) //setting value to nil
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print() //300

	root.right.left.setValue(4)
	fmt.Println()
	root.traverse() //300 4 0

}

//通过指针设置值
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	//1.先遍历它的字数
	node.left.traverse()
	//2、再遍历它自己
	node.print()
	//3、遍历它的右子树
	node.right.traverse()
}

/**
 *node:接收者
 */
func (node treeNode) print() {
	fmt.Println(node.value)
}

//创建
func crNode() {
	var root treeNode
	fmt.Println(root) //{0 <nil> <nil>}

	root = treeNode{value: 5}
	root.left = &treeNode{}
	root.right = &treeNode{3, nil, nil}
	root.right.left = new(treeNode)

	//使用局部变量
	root.right.left = createNode(3)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, nil},
	}

	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> <nil>}]
}

//创建类
type treeNode struct {
	value       int
	left, right *treeNode
}

//没有构造函数，如果特别想用，可以利用指针
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
