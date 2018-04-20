package main

import (
	"fmt"
)

func main() {
	var sum int64 = 0
	j := 0
	// int 就是int 不会随着变大而自动变成int32
	fmt.Println(int64(j) + sum)

	fmt.Printf("%T\n",sum)
	var i int64
	// for v1
	for i =0 ; i<10;i++  {
		sum +=i
		//fmt.Println(sum)
	}
	// for v2
	for ;sum<1000000000000000; {
		sum+=sum
		//fmt.Println(sum)
	}

	// for == while
	for sum>1000{
		sum /= 10
	}
	//big := 100000000000000000000 不同大小的int不能通过赋值来给定类型
	//fmt.Printf("%T",big)
	fmt.Println(sum)
}