package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	testFor()
	fmt.Println("sqrt(2)=", Sqrt(2))
	fmt.Println("sqrt(2)=", math.Sqrt(2))
	testSwitch()
	testSwitch2()
	testSwitch3()

	testDefer()

	testDefer2()
}

func testFor() {
	num := 0
	for i := 0; i <= 100; i++ {
		num += i
	}
	fmt.Println("sum:", num)

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum:", sum)
}

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func testSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func testSwitch2() {
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

func testSwitch3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Printf("Good evening!")
	}
}

func testDefer() {
	defer fmt.Println("World!")
	fmt.Println("Hello")
}

func testDefer2() {
	fmt.Println("Couting:")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
