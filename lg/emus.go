package main

import "fmt"

/**
 * 枚举类型
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-28 21:07<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	enums()
}

func enums() {
	const (
		aa = 1
		b  = 2
		c  = 3
		d  = iota //自增
		_         //跳过
		f
	)
	fmt.Println(aa, b, c, d, f) //1 2 3 3 5
}
