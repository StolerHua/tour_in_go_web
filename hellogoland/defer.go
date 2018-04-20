package main

import "fmt"

func main() {
	i :=1
	// 异步 --压栈，推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	defer fmt.Println(i)
	i+=1
	fmt.Println(i)

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
