package queue

/**
 * 使用别名的方式扩展
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-09-01 9:50<br/>
 * @version 2021.09
 * @since 2021.09
 */

type Queue []int

// Push  the element into the queue.
// 		e.g. q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop  element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
