package main

import "fmt"

/**
 * map
 * @author 康小征<kangxiaozheng@gamutsoft.com>,2021-08-30 15:02<br/>
 * @version 2021.08
 * @since 2021.08
 */
func main() {
	minimaps()
}

func minimaps() {
	//方式一
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m) //map[course:golang name:ccmouse quality:notbad site:imooc]

	//方式二
	m2 := make(map[string]int) //m2== empty map
	fmt.Println(m2)            //map[]

	//方式三
	var m3 map[string]int //m3==nil
	fmt.Println(m3)       //map[]

	//遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	//根据k得v
	course, ok := m["course"] //查看k是否存在并得到k得值
	fmt.Println(course, ok)   //golang true

	if v, flag := m["course"]; flag {
		fmt.Println(v)
	} else {
		panic("不存在健：" + "course")
	}

	//删除
	if _, ok := m["name"]; ok {
		delete(m, "name")
		fmt.Println(m) //map[course:golang quality:notbad site:imooc]
	} else {
		panic("不存在健：" + "name")
	}
}
