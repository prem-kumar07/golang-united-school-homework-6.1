package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

const customError="index out of shapesCapacity!"
// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if(b.shapes ==nil){
		b.shapes= make([]Shape,0,b.shapesCapacity)
		b.shapes=append(b.shapes,shape)
		//fmt.Println(b)
	}else{
		if(len(b.shapes)<b.shapesCapacity){
			b.shapes=append(b.shapes,shape)
			//fmt.Println(b)
		}else{
			fmt.Println("error!")
			return fmt.Errorf("out of shapesCapacity!")
		}
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	//fmt.Println(b)
	if(i>=len(b.shapes)){
		//fmt.Println("error")
		return nil,fmt.Errorf(customError)
	}else{
	//	fmt.Println(b.shapes[i])
		return b.shapes[i],nil
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {

	if(i>=len(b.shapes)){
		return nil,fmt.Errorf(customError)
	}else{
		shape:=b.shapes[i]
		//fmt.Println(shape.CalcPerimeter())
		b.shapes=append(b.shapes[:i], b.shapes[i+1:]...)
		//fmt.Println(b.shapes)
        return shape,nil
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if(i>=len(b.shapes)){
		return nil,fmt.Errorf(customError)
	}else{
		removed:=b.shapes[i]
	    b.shapes[i]=shape
        return removed,nil
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum:=0.0
	for _, shape := range b.shapes {
        sum+=shape.CalcPerimeter()
    }
   return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum:=0.0
	for _, shape := range b.shapes {
        sum+=shape.CalcArea()
    }
   return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	
	y := b.shapes[:0]
	isCircle:=false
for _, shape := range b.shapes {
	switch shape.(type){
		case *Circle:
		  // fmt.Println(index)
		   isCircle=true
	    default:
			//other than circle
          y = append(y, shape)
}
}
b.shapes=y
if(!isCircle){
  return fmt.Errorf("NO Circle!")
}
return nil
}