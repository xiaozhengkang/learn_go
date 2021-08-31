package main

import "fmt"

/**
 * 切片
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-30 10:59<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	t1()
	t2()
	addSlice()
	cjSlices()
	copySlice()
	delSlices()
}

//删除
func delSlices() {
	arr := [...]int{2, 4, 6, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//删除8
	s2 := append(arr[:3], arr[4:]...)
	printSlice(s2) //[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0],len=15,cap=16

	//删除头
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front) //2

	//删除尾
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail) //0
}

// 复制
func copySlice() {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16) //len=16
	copy(s2, s1)
	printSlice(s2) //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0],len=16,cap=16
}

//创建slice
func cjSlices() {
	//方式一
	var s []int //Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	//方式二
	s1 := []int{2, 4, 6, 8}
	printSlice(s1) //len=4,cap=4

	//方式三
	s2 := make([]int, 16) //len=16
	printSlice(s2)        //len=16,cap=16

	s3 := make([]int, 16, 32) //len=16 cap=32
	printSlice(s3)            //len=16,cap=32
}

// 打印 len  和 cap
func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}

//添加元素
func addSlice() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3,s4,s5=", s3, s4, s5) //s3,s4,s5= [6 7 10] [6 7 10 11] [6 7 10 11 12]
	fmt.Println("arr=", arr)             //arr= [1 2 3 4 5 6 7]
}

//获取len 和 caps
func t2() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[3 4 5 6],len(s1)=4,cap(s1)=5
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2)) //s2=[6 7],len(s2)=2,cap(s2)=2
}

func t1() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[2:6]) //[3 4 5 6]
	fmt.Println(arr[2:])  //[3 4 5 6 7]
	fmt.Println(arr[:6])  //[1 2 3 4 5 6]
	fmt.Println(arr[:])   //[1 2 3 4 5 6 7]

	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println(s1)  //[100 4 5 6 7]
	fmt.Println(arr) //[1 2 100 4 5 6 7]
}

func updateSlice(s []int) {
	s[0] = 100
}
