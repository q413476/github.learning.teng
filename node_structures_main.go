package main

import "fmt"

type Node struct {
	Le   *Node
	Data interface{}
	Ri   *Node
}

func NewNode(left, right *Node) *Node {

	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {

	n.Data = data
}

//超人
type Supermaner interface {
	launch() int
}

//超能力
type Abilityer interface {
	damage() int
}

//水属性超能力
type WaterAbilityer interface {
	Abilityer
	Frozen() int
}

//火属性超能力
type FireAbilityer interface {
	Abilityer
	Temperature() int
}

//全能
type God struct {
	Name string
	Age  int
}

func (g *God) damage() int {

	return 999999
}

func (g *God) Frozen() int {

	return 9
}

func (g *God) Temperature() int {

	return 9
}

func (g *God) launch() int {

	return g.damage() + g.Frozen()*2 + g.Temperature()*5
}

func main() {

	root := NewNode(nil, nil)
	root.SetData("root node")

	left := NewNode(nil, nil)
	left.SetData("left node")

	right := NewNode(nil, nil)
	right.SetData("right node")

	root.Le = left
	root.Ri = right

	fmt.Println(root)

	god := new(God)

	fmt.Println(god.launch())

	fmt.Printf("%T\n", god)

	battlefield(god)

	//接口到接口

	var god1 Supermaner
	var god2 Abilityer
	var god3 WaterAbilityer
	var god4 FireAbilityer

	god1 = god
	god2 = god
	god3 = god
	god4 = god

	fmt.Println(god1.launch())
	fmt.Println(god2.damage())

	fmt.Println(god3.Frozen())
	fmt.Println(god3.damage())

	fmt.Println(god4.damage())
	fmt.Println(god4.Temperature())

}

//进入战场
func battlefield(s interface{}) {

	switch v := s.(type) {
	case Supermaner:
		fmt.Println("总伤害", v.launch())
	case WaterAbilityer:
		fmt.Println("基础伤害", v.damage())
		fmt.Println("水伤害", v.Frozen())
	}
}
