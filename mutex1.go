package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex

	//互斥
	countGuard1 sync.RWMutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()

	return count
}
func SetCount(c int) {

	countGuard.Lock()

	fmt.Println("正在写入", c)

	cc := count

	rand.Seed(time.Now().UnixNano())

	a := rand.Float64() / 1000 * 1e9

	time.Sleep(time.Duration(a))

	count = cc + c
	fmt.Println("写入完成", c)
	countGuard.Unlock()

}

func SetCoun1(c int) {

	countGuard1.RLock()

	fmt.Println("正在写入1", c)

	cc := count

	rand.Seed(time.Now().UnixNano())

	a := rand.Float64() / 1000 * 1e9

	time.Sleep(time.Duration(a))

	count = cc + c
	fmt.Println("写入完成1", c)
	countGuard1.RUnlock()

}

func GetCount1() int {
	// 锁定
	countGuard1.RLock()
	// 在函数退出时解除锁定
	defer countGuard1.RUnlock()
	return count
}

func main() {

	// 可以进行并发安全的设置

	for i := 0; i < 10; i++ {

		go SetCount(i)
	}

	fmt.Println(GetCount())

	time.Sleep(2e9)

	fmt.Println(GetCount())

	fmt.Println("--------")


	count = 0

	for i := 0; i < 10; i++ {

		go SetCount(i)
	}

	fmt.Println(GetCount1())

	time.Sleep(2e9)

	fmt.Println(GetCount1())

}
