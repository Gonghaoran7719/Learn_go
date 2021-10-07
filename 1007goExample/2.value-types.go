package main

import "fmt"

//go的一些值类型，可以使用“+” 运算符
func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 = ",1+1)
	fmt.Println("7.0/3.0 = ",7.0/3.0)

	fmt.Println(true && true)
	fmt.Println(true || true)
	fmt.Println(!true)
}

//out :
//golang
//1+1 =  2
//7.0/3.0 =  2.3333333333333335
//true
//true
//false