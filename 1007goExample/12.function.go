package main

import "fmt"

//定义一个函数实现加法
func plus(a, b int) int {
	return a+b
}

//可忽略相同类型的名称
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	//函数调用
	res := plus(1,2)
	fmt.Println("1+2=",res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3=",res)
}
//输出：
//1+2= 3
//1+2+3= 6
