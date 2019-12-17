package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)

	go suck1(ch1, ch2)

	time.Sleep(2e9)

}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck1(ch1, ch2 chan int) {
	k := 0
	for {

		select {
		case v := <-ch1:
			k++
			fmt.Printf("Received on channel 1: %d,number:%d\n", v, k)
		case v := <-ch2:
			k++
			fmt.Printf("Received on channel 2: %d,number:%d\n", v, k)
		}
	}
}
