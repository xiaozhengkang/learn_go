package main

import "fmt"

/**
 * 常量
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-07-23 20:23<br/>
 * @version 2021.07
 * @since 2021.07
 */
func conts() {
	const filename = "abc.txt"
	const a = 3
	const (
		c = 1
		d = 2
	)
	fmt.Println(filename, a, c, d)
}
func main() {
	conts()
}
