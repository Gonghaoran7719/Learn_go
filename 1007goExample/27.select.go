package main

import (
	"fmt"
	"time"
)

//go允许等待多个通道，将goroutine和channel以select结合，是go一个强大的功能

func main()  {
	c1 := make(chan string)
	c2 := make(chan string)

	//给channel发送值的时候隔一段时间，以模拟goroutine在并发执行时阻止执行RPC操作
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	//用select同时等待这两个值并打印他们
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("receive", msg2)
		}
	}
}
//由于1秒和2秒的睡眠同时执行，所以总的执行时间大约只有2秒

//输出：
//received one