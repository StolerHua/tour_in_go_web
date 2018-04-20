package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z :=1.0
	times :=1
	for {
		z2 := z*z
		if math.Abs(z2 - x) <= 0.0000000000001 {
			return z
		} else {
			z -= (z*z-x)/(2*z)
		}
		fmt.Printf("this is %v times to get sqrt %v\n",times,z)
		times+=1
	}

}
func main() {
	fmt.Println(Sqrt(100))
	fmt.Println(math.Sqrt(100))
}
