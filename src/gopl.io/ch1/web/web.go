package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"gopl.io/demo"
)

var mu sync.Mutex
var count map[string]int

func main() {
	//	count = make(map[string])
	http.HandleFunc("/", handler)
	http.HandleFunc("/sin", sinHandler)
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	mu.Lock()
	defer mu.Unlock()
	if count == nil {
		count = make(map[string]int)
	}
	count[url]++
	fmt.Fprintf(w, "URL.Path = %q : Count : %d\n", url, count[url])
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintln(w, "RemoteAddr = %q", r.RemoteAddr)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
func sinHandler(w http.ResponseWriter, r *http.Request) {
	demo.Singif(w)
}
func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	demo.Lissajous(w)
}
