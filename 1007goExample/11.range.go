package main

import "fmt"

func main()  {
	//range 范围迭代元素
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//数组和切片都提供了索引和值，上面我们不需要索引，所以用 _ 忽略了它，不过我们有时候会使用
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{
		"a" : "apple",
		"b" : "banana",
	}
	//键值迭代
	for k, v := range kvs{
		fmt.Printf("%s -> %s\n",k, v)
	}
	//只迭代键
	for kvs := range kvs{
		fmt.Println("key:", kvs)
	}
	//在字符串上迭代，i 索引，c 是字符本身的asc码值
	for i, c := range "golang" {
		fmt.Println(i,c)
	}
}

//输出：
//sum: 9
//index: 1
//a -> apple
//b -> banana
//key: a
//key: b
//0 103
//1 111
//2 108
//3 97
//4 110
//5 103
