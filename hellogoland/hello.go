package main

import (
	"math/rand"
	"fmt"
	"time"
	"math"
)

func add(x , y int) int{
	return x+y
}

func swap(x , y string) (a,b string) {
	a = y
	b =x
	return
}

func swapV2(x,y string)(string,string)  {
	return y,x
}

func main() {
	// 随机
	rand.Seed(time.Now().Unix())
	fmt.Println("hello, 世界",rand.Float64())
	// 宏
	fmt.Println(math.Pi)
	// 函数
	fmt.Println(add(1,2))
	a, b :=swapV2("1h","h2")
	fmt.Println(a,b)
	fmt.Println(swap("1h","h2"))

}

