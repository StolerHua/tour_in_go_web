package main

import "fmt"

// 函数外必须用var，func开头
var python, c ,java bool

func main()  {
	// 没有赋值时，
	var i int
	// 类型明确时，等价于 var python , c = true,"wq"
	python , c := true,"wq"

	fmt.Println(i,python,c,java)
}