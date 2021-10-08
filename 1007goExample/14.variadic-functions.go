package main

import "fmt"

//Variadic functions 可以接受任意长度的参数
//fmt.Println 是一个常见的变量函数
//将任意个整数设置为函数的参数
func sum(nums ...int)  {
	fmt.Print(nums," ")
	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main() {
	//可接受任意个参数
	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4} //如果切片中已经有多个args，使用可变参数函数可以用 func(slice...)
	sum(nums...)
}

//输出；
//[1 2] 3
//[1 2 3] 6
//[1 2 3 4] 10