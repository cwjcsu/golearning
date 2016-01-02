package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func main() {
	fmt.Println("A Randon Number is ", rand.Intn(10))
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
	fmt.Println("Add(42,13) is ", add(42, 13))
	var x, y = "Hello,", "World"
	x, y = swap(x, y)
	fmt.Println("swap(Hello,World)=", x, y)
	fmt.Print("split(17)=")
	fmt.Println(split(17))
	varTest()
	varTest2()
	varTest3()
	testBasicTypes()
	testZeros()
	typeConversions()
	typeInference()
	testConstants()
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func testConstants() {
	const greet = "Hello,World"
	//	fmt.Println(Big)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func typeInference() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
}

func typeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func testZeros() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func testBasicTypes() {
	fmt.Println("Basic Types:")
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

func add(x, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func varTest() {
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)
}

func varTest2() {
	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)
}

func varTest3() {
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
}
