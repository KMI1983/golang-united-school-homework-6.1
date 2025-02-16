package golang_united_school_homework

import (
	"errors"
)

var (
	errorIndexOutOfRange      = errors.New("Index went out of the range")
	errorShapeByIndexNotExist = errors.New("Shape by index doesn't exist")
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

	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("Out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errorIndexOutOfRange
	}
	s := b.shapes[i]
	if s == nil {
		return nil, errorShapeByIndexNotExist
	}
	return s, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errorIndexOutOfRange
	}
	s := b.shapes[i]
	if s == nil {
		return nil, errorShapeByIndexNotExist
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {

	if len(b.shapes) <= i {
		return nil, errorIndexOutOfRange
	}
	s := b.shapes[i]
	if s == nil {
		return nil, errorShapeByIndexNotExist
	}
	b.shapes[i] = shape
	return s, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := 0.0
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum := 0.0
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	circlesCount := 0
	var newShapes []Shape
	for _, shape := range b.shapes {
		_, ok := shape.(*Circle)
		if !ok {
			newShapes = append(newShapes, shape)
		} else {
			circlesCount += 1
		}
	}
	if circlesCount == 0 {
		return errors.New("There are no circles")
	}
	b.shapes = newShapes
	return nil
}
