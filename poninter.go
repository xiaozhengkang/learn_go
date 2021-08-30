package main

import "fmt"

/**
 * 指针
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-30 9:56<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	a, b := 3, 4
	fmt.Println(swap(a, b))
	fmt.Println(a, b) //3,4

	//调用指针
	swapt(&a, &b)
	fmt.Println(a, b) //4,3
}
func swap(a, b int) (c, d int) {
	b, a = a, b
	return a, b
}

//指针
func swapt(a, b *int) {
	*b, *a = *a, *b
}
