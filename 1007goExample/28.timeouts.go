package main

import (
	"fmt"
	"time"
)

//超时是非常重要的对于连接外部资源，需要对执行时间的程序进行绑定；
//有了channel和select，是可以简单而优雅

func main() {
	c1 := make(chan string,1)
	//在调用channel时，在2秒后返回其结果，缓冲channel，发送时是非阻塞的，这是防止channel为读取时goroutine的泄露的常见模式
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	//这是超时的select的案例，select就绪准备接受，等待1s后超时
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second) :
		fmt.Println("timeout 1")
	}

	c2 := make(chan string,1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	//如果允许等待3秒，将可以接收成功，并打印
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second) :
		fmt.Println("timeout 2")
	}
}

//输出：
//timeout 1
//result 2