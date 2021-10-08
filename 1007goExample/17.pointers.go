package main

import "fmt"

//go支持指针，在传递的过程中记录此值

//值传递，会拷贝一个副本
func zeroval(ival int) {
	ival = 0
}

//引用传递，会直接操作指针对应的值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:",i)

	zeroptr(&i) //传入i的地址
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
//out:
//initial: 1
//zeroval: 1   值传递不会修改i，但指针传递会修改，因为它引用了该变量的内存地址
//zeroptr: 0
//pointer: 0xc000014098