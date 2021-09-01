package tree

/**
 *
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-09-01 9:31<br/>
 * @version 2021.09
 * @since 2021.09
 */
func (node *Node) traverse() {
	if node == nil {
		return
	}
	//1.先遍历它的字数
	node.Left.traverse()
	//2、再遍历它自己
	node.print()
	//3、遍历它的右子树
	node.right.traverse()
}
