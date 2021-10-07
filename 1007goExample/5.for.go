package main

import "fmt"

func main()  {
	i := 1
	for i <= 3 { //只有一个条件 （go没有while）
		fmt.Println(i)
		i++
	}

	for j := 7; j < 10; j++ { //初始化/条件/after the loop
		fmt.Println(j)
	}

	for  { //死循环直到打破
		fmt.Println("loop")
		break
	}

	for n := 10; n >= 5; n-- {
		if n%2 == 0 {  //可通过continue跳过此次
			continue
		}
		fmt.Println(n)
	}
}

//out :
//1
//2
//3
//7
//8
//9
//loop
//9
//7
//5
