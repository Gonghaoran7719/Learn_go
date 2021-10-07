package main

import (
	"fmt"
	"math"
)

const s string = "constants"  //定义一个常量
func main() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d)) //go为类型语言，要通过显示类型转换

	fmt.Println(math.Sin(n)) //类型可以通过上下文（赋值、或函数调用）来给定类型，所以sin的类型为int64
}

//out :
//constants
//6e+11
//600000000000
//-0.28470407323754404
