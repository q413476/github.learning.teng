package main

import "fmt"

type container struct {
	items []int
}

func (c container) Len() int {

	return len(c.items)
}

func (c *container) Iter() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < c.Len(); i++ { // or use a for-range loop
			ch <- c.items[i]
		}
		close(ch)
	}()
	return ch
}


func main() {
	c1 := new(container)
	c1.items = []int{1, 2, 3, 4, 5}
	for val := range c1.Iter() {
		fmt.Println(val)
	}
}
