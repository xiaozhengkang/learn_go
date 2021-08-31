package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
 * 函数
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-29 14:46<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	fmt.Println(eval(1, 2, "/"))
	fmt.Println(dev(1, 2))
	q, r := dev2(1, 2)
	fmt.Println(q, r)
	dev3(1, 2)

	if i, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(i)
	}
	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a int, b int) int {
		i := int(math.Pow(float64(a), float64(b)))
		return i
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}

//可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func dev(a, b int) (int, int) {
	return a / b, a % b
}
func dev2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

// 推荐
func dev3(a, b int) (q, r int) {
	return a / b, a % b
}

/**
op:传入一个函数
*/
func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args"+"(%d,%d)", name, a, b)
	return op(a, b)
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := dev3(a, b) //只返回一个
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}
