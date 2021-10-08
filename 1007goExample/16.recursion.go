package main

import (
	"fmt"
)

//递归自己调用自己，直到fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	var fib  func(n int) int //闭包也可以是递归的，但在闭包之前要键入变量显示声明闭包

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

//输出：
//5040
//13