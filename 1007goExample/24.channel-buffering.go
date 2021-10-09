package main

import "fmt"

//默认情况下，channel是无缓冲的，当接受端准备好时，才会接受
//buffer channel会设置一定容量，但是不会对应同等数量的接收器
func main()  {
	//make一个容量为2的channels
	messages := make(chan string,2)

	//此通道是缓冲的，可以将这些值发送到通道里，但无需并发接受
	messages <- "buffered"
	messages <- "channel"
	//接受value
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

//输出：
//buffered
//channel