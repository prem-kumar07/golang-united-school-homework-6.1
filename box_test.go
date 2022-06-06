package golang_united_school_homework

import (
	"testing"
)

func Test_box(t *testing.T) {
	box := NewBox(3)
	box.AddShape(&Triangle{Side: 3})
	box.AddShape(&Circle{Radius: 3})
    t.Log(box.shapes)
}