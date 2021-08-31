package main

import "fmt"

/**
 *
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-28 14:52<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	fmt.Println("Hello Word")
	variable()
	fmt.Println(aa, bb, true)
	variableInitiaal()
}

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitiaal() {
	var d = 1
	print(d)
}

var (
	aa = 3
	bb = "kkk"
	//cc = true
)
