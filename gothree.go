package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v  years)", p.Name, p.Age)
}

func main() {
	a := Person{"Test", 23}
	z := Person{"Test2", 46}
	fmt.Println(a, z)
}
