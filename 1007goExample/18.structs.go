package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p //传出一个指针，让局部变量存活
}

func main()  {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name : "Alice", age : 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(person{name : "Ann", age : 40})

	fmt.Println(newPerson("Jon"))

	s := person{name : "Sean",age : 50}
	fmt.Println(s.name)

	sp := &s //用 . 调用可使指针自动取消
	fmt.Println(sp.age)

	sp.age = 51 //结构体是可变的
	fmt.Println(sp.age)
}

//输出：
//{Alice 30}
//{Fred 0}
//{Ann 40}
//&{Jon 42}
//Sean
//50
//51