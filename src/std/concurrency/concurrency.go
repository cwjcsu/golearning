package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//	testGo()
	//testChannel()
	//testBufferedChannel()

	//	testRangeAndClose()

	//	testSelect()

	//testDefaultSelection()

	//	exerciseBinaryTree()

	//	testMutex()

	testDaisyChain()
}

func f(left, right chan int) {
	left <- 1 + <-right
}
func testDaisyChain() {
	const n = 10000
	leftmost := make(chan int)
	left := leftmost
	right := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func testGo() {
	go say("World")
	say("Hello")
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func testChannel() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) //-5 17 12 ?
}

func testBufferedChannel() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

/**
把fibonacci数组前n个塞入管道c
**/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func testRangeAndClose() {
	fmt.Println("testRangeAndClose:")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func testSelect() {
	fmt.Println("Select test:")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

func testDefaultSelection() {
	fmt.Println("testDefaultSelection:")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print("tick.")
		case <-boom:
			fmt.Print("BOOM!")
		default:
			fmt.Print("  .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree, c chan int, quit chan int) { //中序遍历（左中右）
	if t != nil {
		select {
		case <-quit:
			return
		default:
			traverse(t.Left, c, quit)
			c <- t.Value
			traverse(t.Right, c, quit)
		}
	}
}
func compare(c1, c2, quit chan int) bool {
	v1, v2 := 0, 0
	ok1, ok2 := true, true
	for {
		v1, ok1 = <-c1
		v2, ok2 = <-c2
		if ok1 == ok2 && v1 == v2 {
			if ok1 {
				continue
			} else {
				return true
			}
		} else {
			quit <- 0
			return false
		}
	}
}

func isEqual(t1, t2 *Tree) bool {

	c1 := make(chan int, 2)
	c2 := make(chan int, 2)
	quit := make(chan int)
	doTraverse := func(t *Tree, c chan int, quit chan int) {
		traverse(t, c, quit)
		close(c)
	}
	go doTraverse(t1, c1, quit)
	go doTraverse(t2, c2, quit)
	return compare(c1, c2, quit)
}

func exerciseBinaryTree() {
	fmt.Println("exerciseBinaryTree:")
	var t1, t2 *Tree
	t1 = &Tree{
		&Tree{
			&Tree{Value: 1},
			1,
			&Tree{Value: 2},
		},
		3,
		&Tree{
			&Tree{Value: 5},
			8,
			&Tree{Value: 13},
		},
	}

	t2 = &Tree{
		&Tree{
			&Tree{
				&Tree{Value: 1},
				1,
				&Tree{Value: 2},
			},
			3,
			&Tree{Value: 5},
		},
		8,
		&Tree{Value: 13},
	}
	fmt.Println("isEqual:", isEqual(t1, t2))
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func testMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("someKey"))
}
