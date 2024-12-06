package types

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Printstruct() {
	fmt.Println(Vertex{1, 2})

}
func PrintWithFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
func PrintPointerStruct() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func PrintStructLiterals() {
	fmt.Println(v1, p, v2, v3)
}
