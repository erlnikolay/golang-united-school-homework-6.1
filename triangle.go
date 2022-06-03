package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

// Method CalcPerimeter
func (t *Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

// Method CalcArea
func (t *Triangle) CalcArea() float64 {
	//return (2 * t.Side) / 2
	return (math.Sqrt(3) / 4) * math.Pow(t.Side, 2)
}
