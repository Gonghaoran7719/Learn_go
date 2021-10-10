package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

//输出：
//Tick at 2021-10-10 14:52:59.2518627 +0800 CST m=+0.517900501
//Tick at 2021-10-10 14:52:59.7508221 +0800 CST m=+1.016859901
//Tick at 2021-10-10 14:53:00.2497844 +0800 CST m=+1.515822201
//Ticker stopped