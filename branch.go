package main

import (
	"fmt"
	"io/ioutil"
)

/**
 * 条件
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-07-23 20:41<br/>
 * @version 2021.07
 * @since 2021.07
 */
//if
func ift() {
	const filename = "abc.txt"
	//第一种
	/*constants, err := ioutil.ReadFile(filename) //读取一个文件内容
	  if err != nil {
	  	fmt.Println(err) //open abc.txt: The system cannot find the file specified.
	  } else {
	  	fmt.Printf("%s\n", constants) //123
	  }*/

	//第二种
	if constants, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", constants)
	}
}

//switch
func grade(score int) string { //传递一个百分制的分数 ,返回一个string
	g := ""
	switch {
	case score < 0 || score > 200:
		panic(fmt.Sprintf("wrong score %d", score))
	case score < 50:
		g = "F"
	case score < 90:
		g = "A"
	default:
		panic(fmt.Sprintf("wrong score %d", score))
	}
	return g
}
func main() {
	grade(199)
	ift()
}
