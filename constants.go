package main

import (
	"fmt"
	"math"
)

/**
 * 常量
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-28 21:03<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	consts()
}

const filename = "abc.txt"

func consts() {
	const (
		a, b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
}
