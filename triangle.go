package main

import (
	"fmt"
	"math"
)

/**
 *
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-28 20:57<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	var a, b = 3, 4
	var c int
	// 开平方根
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
