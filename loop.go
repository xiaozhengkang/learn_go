package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * for
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-07-24 14:06<br/>
 * @version 2021.07
 * @since 2021.07
 */
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fliename string) {
	file, err := os.Open(fliename)
	if err != nil {
		panic(err)
	}
	scane := bufio.NewScanner(file) //一行行读
	for scane.Scan() {
		fmt.Println(scane.Text())
	}
}

func forerver() {
	for { //死循环
		fmt.Println("abc")
	}
}
func main() {
	forerver()
	printFile("abc.txt")
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1011 --> 1101
	)
}
