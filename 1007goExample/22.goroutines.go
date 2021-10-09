package main

import (
	"fmt"
	"time"
)

//goroutine是轻量级的执行线程
func f(from string)  {
	for i := 0; i < 3; i++ {
		fmt.Println(from,":", i)
	}
}

func main() {
	//普通调用
	f("direct")
	//在goroutine中调用，使用go f(s),这个新的goroutine将与被调用的函数并行执行
	go f("goroutine")
	//可以为匿名函数启动goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//在两个异步执行的goroutine结束之后在执行下面两个函数（要获得更强大的功能，请使用WaitGroup）
	time.Sleep(time.Second)
	fmt.Println("done")
}

//当执行此程序的时候，首先看到的是阻塞调用的输出
//然后是两个goroutine的输出，goroutine的输出可能会交织在一起，因为goroutine是并行的

//输出：
//direct : 0
//direct : 1
//direct : 2
//going
//goroutine : 0
//goroutine : 1
//goroutine : 2
//done