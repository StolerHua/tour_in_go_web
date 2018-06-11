package main

import "fmt"

// go没有指针运算
//	i, j := 42, 2701
//p := &i
//fmt.Println(*p)

//结构体hhh
type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{x: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p2  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1,2}
	p := &v
	// 可以简单使用结构体指针
	p.x = 1234

	fmt.Println(p)
	fmt.Println(v)
	fmt.Println(v1, p2, v2, v3)
}
