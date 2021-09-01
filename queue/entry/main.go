package main

import (
	"fmt"
	"learn_go/queue"
)

/**
 *
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-09-01 10:01<br/>
 * @version 2021.09
 * @since 2021.09
 */
func main() {
	q := queue.Queue{1} // 原始值

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
