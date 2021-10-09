package main

import "fmt"

//使用channel传递参数时，可以指定channel是仅接收还是发送，增加了安全性

//ping为接收发送值channel，尝试在channel上接收，会编译错误
func ping(pings chan <- string, msg string)  {
	pings <- msg
}

//pong为一个接收channel和一个发送的channel
func pong(pings <- chan string, pongs chan <- string)  {
	msg := <- pings
	pongs <- msg
}

func main()  {
	pings := make(chan string,1)
	pongs := make(chan string,1)
	ping(pings, "passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}

//输出：
//passed message