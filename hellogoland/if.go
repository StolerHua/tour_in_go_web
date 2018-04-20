package main

import (
	"math"
	"fmt"
)

func pow(x ,n, lim float64) float64  {
	// v的作用域在if内
	if v :=math.Pow(x,n);v<lim {
		return v
	}
	return lim
}
func main() {
	fmt.Println(
			pow(3,2,10),
			pow(3,3,20),
		)
	var small uint16 = 1 <<12
	fmt.Printf("%T,%v\n",small,small)
	fmt.Println(1/2)
}
