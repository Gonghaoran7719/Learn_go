package main

import (
	"errors"
	"fmt"
)

//go 可以通过显式、单独的返回值来判断错误
//与ruby、java语言使用异常以及C中使用重载单个结果/错误形成鲜明对比
//go 的方法可以轻松的查看哪些返回值的错误，并进行处理

//错误为最后一个返回值，内置一个错误接口
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1,errors.New("can't work with 42") //构造一个基本的错误值
	}
	return arg + 3, nil //nil 表示没有错误
}

type argError struct { //自定义一个错误，实现Error方法，显示的表示参数错误
	arg int
	prob string
}

func (e *argError) Error() string { //使用两个字段构建一个argError
	return fmt.Sprintf("%d - %s",e.arg,e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} //例子
	}
	return arg + 3, nil
}


func main() {
	//下面用了两个循环来测试函数的返回值
	//在if上使用内联错误检查，在go中比较常见
	for _, i := range []int{7, 42}{
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42}{
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	//如果你想在自定义错误中使用数据，需要类型断言，获取自定义错误的类型
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

//out:
//f1 worked: 10
//f1 failed: can't work with 42
//f2 worked: 10
//f2 failed: 42 - can't work with it
//42
//can't work with it