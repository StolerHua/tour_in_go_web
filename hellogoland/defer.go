package main

import "fmt"
//https://blog.go-zh.org/defer-panic-and-recover
func f() (i int)  {
	defer func() {i++}()
	return 1
}

func x() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func c() int  {
	i :=1
	defer fmt.Println(i)
	defer func() {i++
	fmt.Println(i)
	}()
	return i
}
func main() {
	x()
	fmt.Println("Returned normally from f.")

	// Deferred functions may read and assign to the returning function's named return values .
	fmt.Println(f())
	fmt.Println(c())
	//i :=1
	//// 异步 --压栈，推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	//defer fmt.Println(i)
	//i+=1
	//fmt.Println(i)
	//
	//fmt.Println("counting")
	//
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//
	//fmt.Println("done")
}
