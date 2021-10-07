package main

import (
	"fmt"
)

func main() {
	var a = "initial"  //var 声明变量
	fmt.Println(a)

	var b,c int = 1,2 //可初始化多个值
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int  //为被初始化的值会被赋值为0
	fmt.Println(e)

	f := "apple"  //可用:= 直接初始化，编译器来判断类型
	fmt.Println(f)
}

//out :
//initial
//1 2
//true
//0
//apple