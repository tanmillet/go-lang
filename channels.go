package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//c := make(chan int)
	//go sum(a[:len(a)/2], c) //开起 goroutine
	//go sum(a[:len(a)/2], c) // goroutine
	//x, y := <-c, <-c
	//fmt.Println(x, y)

	c := make(chan int, 1)
	c <- 1
	time.Sleep(2 * time.Millisecond)
	fmt.Println(<-c)
	c <- 2
	fmt.Println(<-c)
}
