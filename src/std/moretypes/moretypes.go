package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	testPointer()
	testStruct()
	testArray()
	testSlices()
	testSlices2()
	testRange()
	testMap()
	testFunctionValues()
	testFunctionClosures()
	testFibonacciClosures()
}

func fibonacci() func() int {
	n1 := 1
	n2 := 2
	return func() int {
		var tmp = n2
		n2 = n1 + n2
		n1 = tmp
		return n2
	}
}

func testFibonacciClosures() {
	fmt.Println("testFibonacciClosures:")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func testFunctionClosures() {
	fmt.Println("testFunctionClosures:")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func testFunctionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func testMap() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	m = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func testRange() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d", i, v)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func testSlices2() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func testSlices() {
	s := []int{2, 3, 4, 5, 6, 7}
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]=%d\n", i, s[i])
	}

	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"
	printBoard(game)

	s = []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])

	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func testArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func testPointer() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func testStruct() {
	fmt.Println(Vertex{1, 2})
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{3, 4}
	)
	fmt.Println(v1, v2, v3, *p)
}

type Vertex struct {
	X int
	Y int
}
