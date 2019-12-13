package main

import (
	"fmt"
	"time"
)

type Empty interface{}

type Semaphore chan Empty

//写入
func (s Semaphore) P(n int) {

	e := new(Empty)

	for i := 0; i < n; i++ {

		s <- e
	}
}

//取出
func (s Semaphore) V(n int) {

	for i := 0; i < n; i++ {

		<-s
	}
}

/* mutexes */
func (s Semaphore) Lock() {
	s.P(1)
}

func (s Semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s Semaphore) Wait(n int) {
	s.P(n)
}

func (s Semaphore) Signal() {
	s.V(1)
}

var II = 1

var semaphore Semaphore

func SetII() {

	semaphore.Lock()

	fmt.Println(222)

	II = II + 1

	semaphore.Unlock()

}

func main() {

	for i := 0; i < 10; i++ {

		go SetII()
	}

	time.Sleep(2e9)

	fmt.Println(II)
}
