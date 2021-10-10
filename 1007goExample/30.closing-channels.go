package main

import "fmt"
//关闭channel表示不在接收发送值，在双方通信的时候可以控制
func main() {
	//make工作channel，用goroutine传达给work goroutine，当没有工人工作时，关闭channel
	jobs := make(chan int, 5)
	done := make(chan bool)

	//这是一个工作的goroutine，反复收到job，用两个值来接收，如果channel关闭，则more为false，则通知done
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	//发送3个job给channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	//结束工作后，关闭channel
	close(jobs)
	fmt.Println("sent all jobs")

	//使用同步的方法等待工人
	<-done
}

//输出：
//sent job 1
//sent job 2
//sent job 3
//sent all jobs
//received job 1
//received job 2
//received job 3
//received all jobs