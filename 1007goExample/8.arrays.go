package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("tmp:",a) //默认初始化为0

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) //通过下标索引
	fmt.Println("len:", len(a)) //数组长度

	b := [5]int{1,2,3,4,5} //直接初始化数组
	fmt.Println("del:", b)

	var twoD [2][3]int //构建多维数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ",twoD)
}

//out:
//tmp: [0 0 0 0 0]
//set: [0 0 0 0 100]
//get: 100
//len: 5
//del: [1 2 3 4 5]
//2d:  [[0 1 2] [1 2 3]]