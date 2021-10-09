package main

import "fmt"

//channels是goroutine的一种通信的方式
//当goroutine并发执行时，一个goroutine可以将值发送到通道，另一个goroutine可以接受该值
func main() {
	//创建一个channel， make(chan val-type)
	message := make(chan string)

	//使用 <- 语法将值发送到通道，通过make一个新的goroutine发送
	go func() {message <- "ping"}()

	//通过这个语法接受值，并打印
	msg := <-message
	fmt.Println(msg)
}

//默认情况下，channel的接受和发送都是阻塞的直到双方都准备好，无需做同步

//输出：
//ping
