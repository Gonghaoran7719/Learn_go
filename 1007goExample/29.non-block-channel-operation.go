package main

import (
	"fmt"
)

//channel 发送和接收是阻塞的，可以使用select，实现阻塞或非阻塞发送接收
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//这是一个非阻塞的接收，如果channel上有可用的值，select将接收，如果没有将立刻执行默认的case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	//非阻塞的消息无法发送到channel，因为该channel没有缓冲区，也没有接收器，所以执行默认case
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//可以使用多个case实现非阻塞的select，尝试用信号进行非阻塞接收
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

//输出：
//no message received
//no message sent
//no activity