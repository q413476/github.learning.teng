package intFunc1

type Number int

func (n *Number) SetVal(i int) {

	*n = Number(i)
}

func (n Number) GetVal() int {

	return int(n)
}
