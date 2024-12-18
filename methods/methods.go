package methods

import (
	"fmt"
	"math"
)

/*
Methods
Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
In this example, the Abs method has a receiver of type Vertex named v.
*/
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func printSimpleMethod() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	/*
		Pointer receivers
		You can declare methods with pointer receivers.
		This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
		For example, the Scale method here is defined on *Vertex.
		Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
		Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.
		With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
		передача объекта по ссылке
	*/

	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())

	/*
		Pointers and functions
		Here we see the Abs and Scale methods rewritten as functions.
		Again, try removing the * from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?
		(If you're not sure, continue to the next page.)
	*/
	v = Vertex{3, 4}
	ScaleFunc(&v, 10)
	fmt.Println(Abs(v))

	v = Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	/*
		Methods and pointer indirection
		Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:
		var v Vertex
		ScaleFunc(v, 5)  // Compile error!
		ScaleFunc(&v, 5) // OK
		while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
		var v Vertex
		v.Scale(5)  // OK
		p := &v
		p.Scale(10) // OK
		For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
	*/
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
	/*
		Methods and pointer indirection (2)
		The equivalent thing happens in the reverse direction.
		Functions that take a value argument must take a value of that specific type:
		var v Vertex
		fmt.Println(AbsFunc(v))  // OK
		fmt.Println(AbsFunc(&v)) // Compile error!
		while methods with value receivers take either a value or a pointer as the receiver when they are called:
		var v Vertex
		fmt.Println(v.Abs()) // OK
		p := &v
		fmt.Println(p.Abs()) // OK
		In this case, the method call p.Abs() is interpreted as (*p).Abs().
	*/
	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p = &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
	/*
		Choosing a value or pointer receiver
		There are two reasons to use a pointer receiver.
		The first is so that the method can modify the value that its receiver points to.
		The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
		In this example, both Scale and Abs are methods with receiver type *Vertex, even though the Abs method needn't modify its receiver.
		In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)
	*/
	z := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", z, z.Abs())
	z.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", z, z.Abs())
}

func Abs(v Vertex) float64 {
	/*
		Methods are functions
		Remember: a method is just a function with a receiver argument.
		Here's Abs written as a regular function with no change in functionality.
	*/
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	/*
		Methods continued
		You can declare a method on non-struct types, too.
		In this example we see a numeric type MyFloat with an Abs method.
		You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
	*/
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
