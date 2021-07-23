package main

import "fmt"

/**
 * 枚举
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-07-23 20:29<br/>
 * @version 2021.07
 * @since 2021.07
 */

func emus() {
	const (
		c      = 1
		python = 2
		golang = 3
		e      = iota //自增
		d
		f
		_ //跳过
		g
		t
		r
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(c, python, golang, e, d, f, g, t, r) //1 2 3 3 4 5 7 8 9
	fmt.Println(b, kb, mb, gb, tb)                   //1 1024 1048576 1073741824 1099511627776
}
func main() {
	emus()
}
