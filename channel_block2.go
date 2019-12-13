package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumber(from, to int, c chan int) {

	for x := from; x <= to; x++ {
		fmt.Printf("%d\n", x)
		time.Sleep(1 * time.Millisecond)
	}

	c <- 0
}

func pump() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

var wait1 sync.WaitGroup


func main() {

	c := make(chan int, 3)

	go printNumber(1, 3, c)
	go printNumber(4, 6, c)

	_ = <-c
	_ = <-c

	suck(pump())

	time.Sleep(1e9)
}
