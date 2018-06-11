package main

import "fmt"

func main() {
	var a [2]byte
	a[0] = '1'
	a[1] = '2'
	fmt.Println(a)
	fmt.Print("%c",a[0])
	fmt.Printf("%s\n",a)

	//切片  不存储任何数据，更改它的元素会导致数组的元素改变
	t := []int{}
	fmt.Println(t)

	slice :=append([]int{1,2,3,4}, 5)
	fmt.Println(slice)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int =primes[1:4]
	fmt.Println(s)
}
