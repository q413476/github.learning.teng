package main

import "fmt"

type Ints []int

func (i Ints) Iterator() *Iterator {
	return &Iterator{
		data:  i,
		index: 0,
	}
}

//用闭包实现
func (i Ints) Do(fn func(int) int) int {
	ri := 0
	for _, v := range i {
		ri = ri + fn(v)
	}
	return ri
}

//用闭包实现迭代器
func (i Ints) IteratorNew() func() (int, bool) {
	index := 0
	return func() (val int, ok bool) {
		if index >= len(i) {
			return
		}

		val, ok = i[index], true
		index++
		return
	}
}

//下面是自定义迭代器
type Iterator struct {
	data  Ints
	index int
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)
}

func (i *Iterator) Next() (v int) {
	v = i.data[i.index]
	i.index++
	return v
}

func main() {

	ints := Ints{1, 2, 3}
	for it := ints.Iterator(); it.HasNext(); {
		fmt.Println(it.Next())
	}
	fmt.Println("----------")
	it := ints.IteratorNew()

	for {
		val, ok := it()
		if !ok {
			break
		}
		fmt.Println(val)
	}
	fmt.Println("----------")
	val := ints.Do(func(i int) int {
		return i * 8
	})
	fmt.Println(val)

}
