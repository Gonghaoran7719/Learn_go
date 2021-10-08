package main

import "fmt"

func main() {
	m := make(map[string]int) //创建一个空的map

	m["k1"] = 7 //使用[键]来赋值
	m["k2"] = 13

	fmt.Println("map:",m) //打印所有键值对

	v1 := m["k1"] //get一个值
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))  //返回键值对的个数

	delete(m,"k2")
	fmt.Println("map:", m) //删除一个键值对

	_,prs := m["k2"] //取k2对应的值，但是我们不需要k2值本身，可以使用标识符 “_”
	fmt.Println("prs:", prs)


	n := map[string]int{ //直接初始化map
		"foo" : 1,
		"bar" : 2,
		}
	fmt.Println("map:", n)
}

//输出：
//map: map[k1:7 k2:13]
//v1: 7
//len: 2
//map: map[k1:7]
//prs: false
//map: map[bar:2 foo:1]