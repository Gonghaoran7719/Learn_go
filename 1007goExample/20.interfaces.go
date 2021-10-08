package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect1 struct { //在矩形和圆形实现此接口
	width,height float64
}

type circle struct {
	radius float64
}

func (r rect1) area() float64 { //在rect实现此接口
	return r.width * r.height
}

func (r rect1) perim() float64 {
	return 2*r.width * 2*r.height
}

func (c circle) area() float64 { //在circle实现此接口
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//如果变量具有接口类型，可以用调用命名调用里面的方法
func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	//rect 和 circle都实现了接口，传参来调用
	r := rect1{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

//输出：
//{3 4}
//12
//48
//{5}
//78.53981633974483
//31.41592653589793
