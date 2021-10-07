package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //可写多个表达式
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now()
	fmt.Println(t)
	switch{ //没有表达式的switch类似与if-else 可将时间转化为 小时
	case t.Hour() < 12:
		fmt.Println("It's before noon" )
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {  //函数为判断一个变量的类型，可将函数赋值为一个变量，参数类型为接口
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Println("Don't know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
//out:
//Write 2 as
//two
//It's weekday
//2021-10-07 21:11:07.3041197 +0800 CST m=+0.024858301
//It's after noon
//I'm a bool
//I'm a int
//Don't know type %T
//hey