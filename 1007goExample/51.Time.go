package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

//时间
//2021-10-10 15:11:25.5561798 +0800 CST m=+0.003784601
//2009-11-17 20:34:58.651387237 +0000 UTC
//2009
//November
//17
//20
//34
//58
//651387237
//UTC
//Tuesday
//true
//false
//false
//104266h36m26.904792563s
//104266.6074735535
//6.255996448413209e+06
//3.7535978690479255e+08
//375359786904792563
//2021-10-10 07:11:25.5561798 +0000 UTC
//1997-12-26 09:58:31.746594674 +0000 UTC