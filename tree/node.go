package tree

import (
	"fmt"
)

/**
 *
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-31 9:32<br/>
 * @version 2021.08
 * @since 2021.08
 */

//遍历
func Bl() {
	var root Node
	root = Node{value: 5}
	root.print() //5

	//root.right.left.setValue(4) //panic: runtime error: invalid memory address or nil pointer dereference没有赋予地址,要使用必须实例化
	root.right = new(Node)
	//使用局部变量
	root.right.Left = createNode(3)
	root.right.Left.setValue(4)
	root.right.Left.print()

	var pRoot *Node
	pRoot.setValue(200) //setting value to nil
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print() //300

	root.right.Left.setValue(4)
	fmt.Println()
	root.traverse() //300 4 0
}

//通过指针设置值
func (node *Node) setValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil")
		return
	}
	node.value = value
}

/**
 *node:接收者
 */
func (node Node) print() {
	fmt.Println(node.value)
}

//创建
func CrNode() {
	var root Node
	fmt.Println(root) //{0 <nil> <nil>}

	root = Node{value: 5}
	root.Left = &Node{}
	root.right = &Node{3, nil, nil}
	root.right.Left = new(Node)

	//使用局部变量
	root.right.Left = createNode(3)

	nodes := []Node{
		{value: 3},
		{},
		{6, nil, nil},
	}

	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> <nil>}]
}

//创建类
type Node struct {
	value       int
	Left, right *Node
}

//没有构造函数，如果特别想用，可以利用指针
func createNode(value int) *Node {
	return &Node{value: value}
}
