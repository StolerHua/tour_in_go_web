package main

import (
	"fmt"
	"math/cmplx"
)

var print = fmt.Printf

// 常量不能用 := 声明
const (
	World  = "world"
	Big = 1<<100
)
var(
	ToBe bool =false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main()  {
	const World  = 3
	fmt.Println(World)
	print("%v\n",Big*0.1)
	// 类型转换
	print("Type: %T Value: %v\n",ToBe,ToBe)
	print("Type: %T Value: %v\n",float32(MaxInt),float32(MaxInt))
	print("Type: %T Value: %v\n",z,z)
}
