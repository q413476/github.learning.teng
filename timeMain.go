package main

import (
	"fmt"
	"time"
)

type ch chan time.Time

func (c ch) P(i int) {

	go func(j int, c ch) {
		kk := 0
		for {
			if kk >= j {
				//退出管道并关闭通道
				close(c)
				break
			} else {
				//每秒执行一次
				time.Sleep(1e9)
				c <- time.Now()
			}
			kk++

		}
	}(i, c)

}

func main() {

	time1 := time.Now()

	fmt.Printf("%T\n", time1)

	fmt.Println(time1)
	fmt.Println(time1.Date())
	fmt.Println(time1.UTC())
	fmt.Println(time1.Year(), time1.Month(), time1.Day(), time1.Weekday())
	fmt.Println(time1.Format("2006-01-02 15:04:03"))

	fmt.Println(time1.Unix())

	time2 := time.Unix(1542868458, 0)

	fmt.Println(time2)

	fmt.Println(time2.Add(time.Hour))
	fmt.Println(time2.Add(-time.Hour))
	fmt.Println(time2.Add(time.Microsecond))
	fmt.Println(time.April)

	//定时器
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
		break
	}

	fmt.Println("-----------")

	//自己写个定时器
	v := make(ch, 1)
	v.P(1)

	for num := range v {
		fmt.Println(num)
	}

	fmt.Println(int64(1e9))

	fmt.Printf("%T\n",1E9)
	fmt.Println(int64(1E9))
	fmt.Println(int64(1e9))


}
