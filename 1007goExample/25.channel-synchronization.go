package main

import (
	"fmt"
	"time"
)

//channel 同步,可以使用channel在goroutine中同步执行，使用阻塞接收等待goroutine完后（也可以使用WaitGroup）
//在函数已完成时，channel可以用于通知另一个goroutine该函数已完成
func worker(done chan bool)  {
	fmt.Print("working...")
	time.Sleep(time.Second)
	//fmt.Println("done")

	done <- true //发发送一个true来通知已完成
}

func main()  {
	//make一个channel用来传递函数完成的信号
	done := make(chan bool,1)
	go worker(done)

	//<- done //阻塞，直到收到channel已完成的通知
	if <-done { //也可以用if判断是否接收到true
		fmt.Println("done")
	} //如果没有阻塞接收，程序将在worker函数结束前 退出
}

//输出：
//working...done