package main

import (
	"fmt"
)

func main() {
	//不需要条件括号，但是需要大括号
	if 7%2 == 0 {
		fmt.Println("7 is even")
	}else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0	{
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { //在条件中声明的内容，可在其他else中继续使用
		fmt.Println("is negative")
	} else if num < 10 {
		fmt.Println(num,"has 1 digit")
	} else {
		fmt.Println(num,"has multiple digits")
	}
}

//out :
//7 is odd
//8 is divisible by 4
//9 has 1 digit