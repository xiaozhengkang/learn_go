package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
 *   定义变量
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-07-21 15:31<br/>
 * @version 2021.07
 * @since 2021.07
 */
var (
	aa, bb, cc = 1, 2, 3
)

//定义一个初始化的变量
func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}

func variable() {
	var a int = 1
	var b string = "刘军"
	//规定类型 后面必须是int类型的
	var c, d int
	//不规定类型，可以自定义
	var e, f, g, h = 1, 'a', "", true
	fmt.Println(e, f, g, h)                 //1 97  true
	fmt.Printf("%d %q %q %q\n", a, b, c, d) //1 "刘军" '\x00' '\x00'
}

//更简便的自定义
func varialbeShorter() {
	a, b, c, d, e, f := 1, 2, 3, 4, 5, 6

	//覆盖上面的值
	b = 11
	b = 12

	fmt.Println(a, b, c, d, e, f) //1 12 3 4 5 6
}

//验证欧拉公式
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

//强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	triangle()
	euler()
	fmt.Println("Hello Word")
	variable()
	variableZeroValue()
	varialbeShorter()

	println(aa, bb, cc)
}
