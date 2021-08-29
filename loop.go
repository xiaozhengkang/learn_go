package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * 循环
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-29 14:30<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	fmt.Println("开始........")
	fmt.Println(
		convertToBin(5),         //101
		convertToBin(13),        //1011 -> 1101
		convertToBin(777777733), //101
	)

	printFile("abc.txt")
	forerver()
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err) //报错
	}
	scane := bufio.NewScanner(file) //一行行读
	for scane.Scan() {              //可以替换while
		fmt.Println(scane.Text())
	}
}

func forerver() {
	for { //死循环
		fmt.Println("abc")
	}
}
