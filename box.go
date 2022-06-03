package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	ErrIndexDoesntExist   = errors.New("Index of element doesn't exist")
	ErrCirclesDoesntExist = errors.New("Circles doesn't exist")
	ErrMaxOfQuantity      = errors.New("Maximum quantity of shapes reached")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	var err error
	//panic("implement me")
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		if err != nil {
			return err
		}
	} else {
		return ErrMaxOfQuantity
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	//panic("implement me")
	// index of elem higher than len of slice
	if i > len(b.shapes)-1 {
		return nil, ErrIndexDoesntExist
	}
	for idxS, valS := range b.shapes {
		if idxS == i {
			return valS, nil
		}
	}
	// index of elem is not found
	return nil, ErrIndexDoesntExist
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var newshapes []Shape
	var triggerOfNotExisttenElement bool
	var valueOfFoundElement Shape
	//panic("implement me")
	// index of elem higher than len of slice
	if i > len(b.shapes)-1 {
		return nil, ErrIndexDoesntExist
	}
	triggerOfNotExisttenElement = false
	for idxS, valS := range b.shapes {
		if idxS == i {
			triggerOfNotExisttenElement = true
			valueOfFoundElement = valS
		} else {
			// re-create shapes
			newshapes = append(newshapes, valS)
		}
	}
	if triggerOfNotExisttenElement {
		b.shapes = newshapes
		return valueOfFoundElement, nil
	} else {
		// index of elem is not found
		return nil, ErrIndexDoesntExist
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	//panic("implement me")
	// index of elem higher than len of slice
	if i > len(b.shapes)-1 {
		return nil, ErrIndexDoesntExist
	}
	for idxS, valS := range b.shapes {
		if idxS == i {
			b.shapes[idxS] = valS
			return valS, nil
		}
	}
	// index of elem is not found
	return nil, ErrIndexDoesntExist
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumP float64
	//panic("implement me")
	for _, valS := range b.shapes {
		sumP = sumP + valS.CalcPerimeter()
	}
	return sumP
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumA float64
	//panic("implement me")
	for _, valS := range b.shapes {
		fmt.Printf("Area: %f, Type of shape: %T \n", valS.CalcArea(), valS)
		sumA = sumA + valS.CalcArea()
	}
	return sumA
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var triggerCircleIsNotExist bool
	var newshapes []Shape
	//panic("implement me")
	triggerCircleIsNotExist = false
	for _, valS := range b.shapes {
		switch valS.(type) {
		case *Circle:
			triggerCircleIsNotExist = true
		default:
			newshapes = append(newshapes, valS)
		}
	}
	if !triggerCircleIsNotExist {
		// circles there are not
		return ErrCirclesDoesntExist
	} else {
		b.shapes = newshapes
		return nil
	}
}
