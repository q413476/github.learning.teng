package main

import "fmt"

func main() {

	f11()
	fmt.Println("Returned normally from f.")


}

func f11() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g11(0)
	fmt.Println("Returned normally from g.")
}

func g11(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g11(i + 1)
}