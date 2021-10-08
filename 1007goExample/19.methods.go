package main

import "fmt"

type rect struct {
	width,height int
}

func (r *rect) area() int  { //area的方法接收者为rect
	return r.width * r.height
}

func (r rect) perim() int  { //可以是指针也可以是值
	return 2*r.width + 2*r.height
}

func main()  {
	r := rect{width: 10, height: 5}

	//两种方法的调用
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	//go的自动处理的方法调用是值和指针之间的转换，
	//使用指针接收器类型，可以避免在方法调用的时候复制，或者允许方法更改结构体
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim", rp.perim())
}

//输出：
//area: 50
//perim: 30
//area: 50
//perim 30
