package main

import (
	"net/http"
	"fmt"
	"log"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Creeting string
	Punct    string
	Who      string
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s %s", s.Creeting, s.Punct, s.Who)
}

func main() {
	http.Handle("/string", String("This is test"))
	http.Handle("/struct", &Struct{"Hello", ":", "Demo"})

	err := http.ListenAndServe("127.0.0.1:4000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
