package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
 * 欧拉公式
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-28 20:47<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	euler()
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi), cmplx.Pow(math.E, 1i*math.Pi)+1)
}
