package main

import "os"

func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

//输出：
//panic: a problem
//
//goroutine 1 [running]:
//main.main()
