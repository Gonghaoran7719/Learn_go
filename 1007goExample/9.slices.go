package main

import "fmt"

func main()  {
	//比数组更强大的：切片
	// a key data key in GO
	s := make([]string,3)
	fmt.Println("emp:", s) //输出为空

	s[0] = "a"//设置
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2]) //获取，

	fmt.Println("len:", len(s)) //长度

	s = append(s,"d") //可添加新的元素，但要重新赋值给s，可能会扩容，元素迁移
	s = append(s,"e","f")
	fmt.Println("app:", s)

	c := make([]string,len(s)) //切片也可复制
	copy(c,s)
	fmt.Println("cpy:", c)

	l1 := s[2:5] //切片可取到指定区间的值 slice[low : high],就可获取到 s[2],s[3],s[4]
	fmt.Println("sl1:", l1)

	l2 := s[:5] //取到5，但不包括s[5]
	fmt.Println("sl2", l2)

	l3 := s[2:] //从2开始取值，但包括s[2]
	fmt.Println("sl3", l3)

	t := []string{"g","h","i"} //可直接定义变量和初始化
	fmt.Println("del:", t)

	twoD := make([][]int,3) //切片可组成多维数组结构，与多维数组不同，内部切片长度可能也会有所不同
	for i := 0; i < 3; i++{
		innerLen := i + 1
		twoD[i] = make([]int,innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:",twoD)
}

//out :
//emp: [  ]
//set: [a b c]
//get: c
//len: 3
//app: [a b c d e f]
//cpy: [a b c d e f]
//sl1: [c d e]
//sl2 [a b c d e]
//sl3 [c d e f]
//del: [g h i]
//2d: [[0] [1 2] [2 3 4]]