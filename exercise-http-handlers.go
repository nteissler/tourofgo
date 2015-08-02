package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v, %v, take this: %v", s.Greeting, s.Punct, s.Who)
	return
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.Handle("/string", String("I'm a frayed knot."))
	http.ListenAndServe("localhost:4000", nil)
}
