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

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s.Greeting, s.Punct, s.Who)
}

func main() {
	http.Handle("/string", String("whatever string"))
	http.Handle("/struct", &Struct{"Hello", ":", "Golang"})
	http.ListenAndServe("localhost:4000", nil)
}
