package main

import "fmt"
//上个例子，对范围的数据进行了迭代，也可以用此语法来完成channel的接收
func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//由于关闭了channel，迭代在接收两次后停止
	for elem := range queue {
		fmt.Println(elem)
	}
}

//关闭非空channel，里面的值还会在里面

//输出：
//one
//two
