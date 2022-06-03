package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

// Method CalcPerimeter
func (c *Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Method CalcArea
func (c *Circle) CalcArea() float64 {
	return math.Pi * c.Radius
}
