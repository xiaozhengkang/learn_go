package main

import (
	"fmt"
	"io/ioutil"
)

/**
 * 循环
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-29 14:13<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	ifUat()
	fmt.Println(
		grade(1011),
		grade(0),
		grade(59),
		grade(80),
		grade(90),
		grade(100),
	)
}

//switch
func grade(sorce int) string {
	g := ""
	switch {
	case sorce < 0 || sorce > 100:
		panic(fmt.Sprintf("Wrong score: %d", sorce)) //报错
	case sorce < 60:
		g = "F"
	case sorce < 80:
		g = "C"
	case sorce < 90:
		g = "B"
	case sorce <= 100:
		g = "A"
		/*	default:
			panic(fmt.Sprintf("Wrong score: %d", sorce)) //报错*/
	}
	return g
}

// if 判断
func ifUat() {
	const filename = "abc.txt"
	//方式一
	/*contents, err := ioutil.ReadFile(filename)
	  if err != nil {
	  	fmt.Println(err)
	  } else {
	  	fmt.Printf("%s\n", contents) //open abc.txt: The system cannot find the file specified.
	  }*/

	// 改造二 推荐
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
