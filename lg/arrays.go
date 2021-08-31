package main

import "fmt"

/**
 * 数组
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-30 10:21<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	var arr1 [5]int                      //这种数量必须固定
	arr2 := [3]int{1, 3, 4}              //这种必须告诉初值
	arr3 := [...]int{2, 3, 4, 7, 5, 678} //利用切片自定义数量
	arr4 := [...]int{2, 3, 4, 7, 5}      //利用切片自定义数量
	var grid [4][5]int                   //二维数组  4行 5列
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历数组 老的方式
	/*	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}*/

	//推荐
	//for i, v := range arr3 { //下标和数值
	for _, v := range arr3 { //不要下标只要数值
		fmt.Println(v)
	}

	printArray(arr1)
	printArray(arr4)

	fmt.Println(arr1, arr4) //值不变

	printArrayPoit(&arr1)
	printArrayPoit(&arr4)

	fmt.Println(arr1, arr4) //值改变

}

func printArray(arr [5]int) {
	arr[0] = 10000000
	for _, v := range arr { //不要下标只要数值
		fmt.Println(v)
	}
	//arr[0] = 10000000
}

func printArrayPoit(arr *[5]int) {
	arr[0] = 10000000
	for _, v := range arr { //不要下标只要数值
		fmt.Println(v)
	}
	//arr[0] = 10000000
}
