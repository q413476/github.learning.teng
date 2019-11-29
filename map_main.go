package main

import (
	"fmt"
	"sort"
)

//map 是可以动态增长的

//key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float
//所以数组、切片和结构体不能作为 key，但是指针和接口类型可以。如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，
//这样可以通过结构体的域计算出唯一的数字或者字符串的 key。
//map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节
//map 是 引用类型 的： 内存用 make 方法来分配。

func main() {

	var map1 map[int]string

	//map1[1] = "sdfsd" 报错

	fmt.Println(map1)

	//mapCreated := make(map[string]float) 相当于：mapCreated := map[string]float{}

	map2 := map[int]string{1: "ab", 2: "cd"}

	fmt.Println(map2)

	map3 := make(map[int]string)

	fmt.Println(map3)
	map3[1] = "kk"
	map3[2] = "kk"
	map3[3] = "kk"
	map3[4] = "kk"
	fmt.Println(map3)

	//map4 := map[int]func() int{}
	map4 := make(map[int]func() int)

	map4[1] = func() int {

		return 1
	}

	fmt.Println(map4)

	//和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。
	// 但是你也可以选择标明 map 的初始容量 capacity
	map5 := make(map[string]float64, 100)

	fmt.Println(map5)
	//当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。
	//所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明

	//用切片作为 map 的值
	map6 := make(map[int][]int)
	map6[1] = make([]int, 3, 6)
	map6[2] = make([]int, 6, 9)
	fmt.Println(map6)

	//测试键值对是否存在及删除元素
	map7 := map[int]string{1: "ab", 2: "cd", 3: "ef", 4: "gh", 5: "ij", 6: "kl"}
	val, ok := map7[3]
	fmt.Println(val, ok)
	delete(map7, 3)
	val1, ok1 := map7[3]
	fmt.Println(val1, ok1)

	//无序的，每次输出结果顺序不一样
	for val2, key2 := range map7 {
		fmt.Println(val2, key2)
	}
	//排序
	//map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
	//如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序
	//然后可以使用切片的 for-range 方法打印出所有的 key 和 value
	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	keys := make([]string, len(barVal))

	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Println(k, barVal[k])
	}

}
