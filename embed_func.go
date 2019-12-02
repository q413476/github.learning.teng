package main

import (
	"fmt"
)

type Log struct {
	msg string
}

func (l *Log) SetMsg(str string) {
	l.msg = str
}


//进行多重继承
type Msg struct {
	code int
	Log
}

func (m *Msg) setMsgType(code int, msg string) {
	m.Log.SetMsg(msg)
	m.code = code
}

func (m Msg) getMsg() (int, string) {

	return m.code, m.msg
}



type LogArray [2]Log

type Customer struct {
	Name string
	log  [] *Log
	Log
	LogArray
	Msg
}

//聚合（或组合）：包含一个所需功能类型的具名字段
//内嵌的类型不需要指针
func main() {

	customer := new(Customer)
	customer.Name = "时间"

	customer.log = append(customer.log, &Log{"第一次添加"})
	customer.log = append(customer.log, &Log{"第二次添加"})
	customer.log = append(customer.log, &Log{"第三次添加"})

	customer.log[0].SetMsg("通过方法添加")
	customer.SetMsg("通过customer的方法添加")
	customer.Log.SetMsg("通过customer的Log属性方法添加")

	customer.LogArray[0].msg = "添加到属性数组里面1"
	customer.LogArray[1].msg = "添加到属性数组里面2"

	customer.setMsgType(1000, "好消息")
	customer.Msg.SetMsg("好消息2")

	fmt.Println(customer)
	fmt.Println(customer.log, len(customer.log), cap(customer.log))
	fmt.Println(customer.log[0])
	fmt.Println(customer.log[1])

}
