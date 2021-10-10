package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

//输出:
//2021-10-10 15:12:06.0772998 +0800 CST m=+0.005030701
//1633849926
//1633849926077
//1633849926077299800
//2021-10-10 15:12:06 +0800 CST
//2021-10-10 15:12:06.0772998 +0800 CST