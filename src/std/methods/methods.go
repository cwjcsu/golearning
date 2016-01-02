package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	testMethods()
	testInterface()
	testImplictedInterface()
	testString()
	testIPAddr()
	testError()
	exerciseErrors()

	testRot13Reader()

	//testWebServers()

	exerciseHttpHandlers()
}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(s))
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s%s", s.Greeting, s.Punct, s.Who)
}

func exerciseHttpHandlers() {
	http.Handle("/string", String("I'm a luck boy!"))
	http.Handle("/struct", &Struct{"Hello", ":", "Atlas"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

type Hello struct {
}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func testWebServers() {
	var h Hello
	var bindAddr = "localhost:4000"
	err := http.ListenAndServe(bindAddr, h) //blocked here
	if err != nil {
		log.Fatal(err)
	} else { //never here
		log.Println("listening ", bindAddr)
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func testMethods() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	//-------------
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	//-------------

	v = &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

type Abser interface {
	Abs() float64
}

func testInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func testImplictedInterface() {
	var w Writer

	// os.Stdout 实现了 Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func testString() {
	a := Person{"Atlas", 28}
	b := Person{"秦始皇", 2218}
	fmt.Println(a, b)
}

func testIPAddr() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it don't work",
	}
}

func testError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintln("cannot Sqrt negative number:", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func exerciseErrors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	if err != nil {
		return n, err
	}
	for i, v := range p {
		if between(97, 122, v) {
			new := v + 13
			if new > 122 {
				new -= 26
			}
			p[i] = new
		} else if between(65, 90, v) {
			new := v + 13
			if new > 90 {
				new -= 26
			}
			p[i] = new
		}
	}
	return n, err
}

func between(start, end, value byte) bool {
	if value > end {
		return false
	} else if value < start {
		return false
	}
	return true
}

func testRot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
