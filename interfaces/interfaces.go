package interfaces

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type I interface {
	M()
}

type T struct {
	S string
}

func printInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
	a = &v

	fmt.Println(a.Abs())
	/*
		Interfaces are implemented implicitly
		A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
		Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
	*/
	var i I = &T{"hello"}
	i.M()
	/*
		Interface values
		Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
		(value, type)
		An interface value holds a value of a specific underlying concrete type.
		Calling a method on an interface value executes the method of the same name on its underlying type.
	*/

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
	/*
		Interface values with nil underlying values
		If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
		In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method Ms in this example.)
		Note that an interface value that holds a nil concrete value is itself non-nil.
	*/

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	/*
		Nil interface values
		A nil interface value holds neither value nor concrete type.
		Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	*/
	var inter I
	describe(inter)
	inter = &T{"hello"} //comment to check
	inter.M()
	/*
		The empty interface
		The interface type that specifies zero methods is known as the empty interface:
		interface{}
		An empty interface may hold values of any type. (Every type implements at least zero methods.)
		Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
	*/
	var il interface{}
	describeInterface(il)

	il = 42
	describeInterface(il)

	il = "hello"
	describeInterface(il)

	/*
		Type assertions
		A type assertion provides access to an interface value's underlying concrete value.
		t := i.(T)
		This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
		If i does not hold a T, the statement will trigger a panic.
		To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
		t, ok := i.(T)
		If i holds a T, then t will be the underlying value and ok will be true.
		If not, ok will be false and t will be the zero value of type T, and no panic occurs.
		Note the similarity between this syntax and that of reading from a map.
	*/
	var in interface{} = "hello"

	s := in.(string)
	fmt.Println(s)

	s, ok := in.(string)
	fmt.Println(s, ok)

	fs, ok := in.(float64)
	fmt.Println(fs, ok)

	//fs = in.(float64) // panic
	fmt.Println(fs)

	/*
		Type switches
		A type switch is a construct that permits several type assertions in series.
		A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.
		switch v := i.(type) {
		case T:
		    // here v has type T
		case S:
		    // here v has type S
		default:
		    // no match; here v has the same type as i
		}
		The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.
		This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.
	*/
	do(21)
	do("hello")
	do(true)
}
func (t T) Ms() {
	fmt.Println(t.S)
}
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Interface values
type F float64

func (f F) M() {
	fmt.Println(f)
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func describeInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
