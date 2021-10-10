package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

//输出：
//request 1 2021-10-10 14:57:02.4960533 +0800 CST m=+0.209950701
//request 2 2021-10-10 14:57:02.6987228 +0800 CST m=+0.412620201
//request 3 2021-10-10 14:57:02.90154 +0800 CST m=+0.615437401
//request 4 2021-10-10 14:57:03.1046995 +0800 CST m=+0.818596901
//request 5 2021-10-10 14:57:03.2923056 +0800 CST m=+1.006203001
//request 1 2021-10-10 14:57:03.2923056 +0800 CST m=+1.006203001
//request 2 2021-10-10 14:57:03.2923056 +0800 CST m=+1.006203001
//request 3 2021-10-10 14:57:03.2923056 +0800 CST m=+1.006203001
//request 4 2021-10-10 14:57:03.4949237 +0800 CST m=+1.208821101
//request 5 2021-10-10 14:57:03.6965284 +0800 CST m=+1.410425801