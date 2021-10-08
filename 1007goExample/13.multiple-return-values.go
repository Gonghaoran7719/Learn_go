package main

import "fmt"

//go的函数可返回多个值，在工程中同程用于返回结果值和error
func vals() (int,int) {
	return 3, 7
}

func main() {
	//使用两个值 来接函数的返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//如果不需要第一个返回值，可以使用 _
	_, c := vals()
	fmt.Println(c)
}

//输出：
//3
//7
//7
