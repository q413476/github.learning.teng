package main

import (
	"fmt"
	"time"
)

func sum(bigArray [2]int, c chan int) {

	fmt.Println("a", "正在执行别的,估计5秒")

	time.Sleep(5 * 1e9)

	var i = 0

	for _, v := range bigArray {

		i = v + i
	}
	fmt.Println()
	fmt.Println("a", "执行别的完了")
	c <- i
}

func main() {

	c := make(chan int)

	go func() {

		fmt.Println("准备发送")

		time.Sleep(2 * 1e9)

		c <- 20

		fmt.Println("发送完成")

	}()

	fmt.Println("等待")

	x := <-c

	fmt.Println("结束", x)

	go sum([2] int{3, 411}, c)

	fmt.Println("趁你你5秒钟，我干点别的")

	fmt.Println(<-c)

	fmt.Println("你们终于完成了")
}
