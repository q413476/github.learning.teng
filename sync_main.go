package main

import (
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

type Info struct {
	mu   sync.Mutex
	Name string
	Age  int
	num  int
}

func Update(info *Info, i int) {
	info.num = i
	if i == 1 {
		time.Sleep(1e9)
	}
	info.mu.Lock()
	info.Name = "age"
	info.Age = i
	info.mu.Unlock()
}

func main() {

	info := new(Info)

	info.Age = 1231
	go Update(info, 1)
	go Update(info, 2)

	time.Sleep(2e9)

	fmt.Println(info)

	im := big.NewInt(math.MaxInt64)

	fmt.Println(im)
	fmt.Printf("%T\n",im)
}
