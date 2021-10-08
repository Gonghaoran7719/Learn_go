package main

import "fmt"

//go支持匿名函数，这些函数称之为闭包
//当想内联定义函数，无需命名

//定义一个函数，返回值为一个闭包函数（返回的函数在i上关闭以形成闭包）
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main()  {
	//调用intSeq函数，将结果值返回给nextInt，函数捕获自己的i值，每次调用都会更新i值
	nextInt := intSeq()

	//调用多次来验证闭包
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//创建一个新的函数，会重新开始intSeq函数
	newInts := intSeq()
	fmt.Println(newInts())
}

//输出：
//1
//2
//3
//1
