package main

import (
	"fmt"
	"runtime"
	"time"
)

func WhereIsSaturday()  {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func SwitchTrue()  {
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	fmt.Print("Go runs on ")
	// 自动添加break，除非以fallthrough结尾
	//switch 
	switch os := runtime.GOOS;os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s\n",os)
	}
	fmt.Println(runtime.GOOS)

	WhereIsSaturday()
	SwitchTrue()
}
