package main

import (
	"fmt"
	"unicode/utf8"
)

/**
 *
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-30 16:02<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	s := "Yes go曹甜梦"
	fmt.Println(len(s)) //15
	fmt.Printf("", []byte(s))

	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	println(utf8.RuneCountInString(s)) //9

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) //Y e s   g o 曹 甜 梦
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch) //(0 Y) (1 e) (2 s) (3  ) (4 g) (5 o) (6 曹) (7 甜) (8 梦)
	}
	fmt.Println()
}
