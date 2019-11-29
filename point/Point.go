package point

import "math"

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {

	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

type NamedPoint struct {
	Point
	Name string
}

//覆盖上一级的方法
func (n *NamedPoint) Abs() float64 {

	return n.Point.Abs() * 100.
}
