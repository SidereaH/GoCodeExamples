package pacvarfun

import (
	"fmt"
	"math/rand"
)

// variables list
var c, python, java bool

func Main() {

	fmt.Println("My favorite number is", rand.Intn(10))
	//func example
	fmt.Println(Add(42, 13))
	//multiple function example func
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	//return named value
	fmt.Println(split(7))
	/*
		Variables
		The var statement declares a list of variables; as in function argument lists, the type is last.
		A var statement can be at package or function level. We see both in this example.
	*/
	var i int
	python = true
	fmt.Println(i, c, python, java)
	/*
		Variables with initializers
		A var declaration can include initializers, one per variable.
		If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	*/
	testVariablesWithInitializers()
	/*
		Short variable declarations
		Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
		Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	*/
	testShortVariableDeckarations()

	/*

		Basic types
		Go's basic types are
		bool
		string
		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr
		byte // alias for uint8
		rune // alias for int32
		     // represents a Unicode code point
		float32 float64
		complex64 complex128
		The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
		The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
	*/
	testBasicTypes()
	/*

		Type conversions

		The expression T(v) converts the value v to the type T.

		Some numeric conversions:

		var i int = 42
		var f float64 = float64(i)
		var u uint = uint(f)

		Or, put more simply:

		i := 42
		f := float64(i)
		u := uint(f)

		Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

	*/
	testTypeConversion()

	/*

		Type inference

		When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

		When the right hand side of the declaration is typed, the new variable is of that same type:

		var i int
		j := i // j is an int

		But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

		i := 42           // int
		f := 3.142        // float64
		g := 0.867 + 0.5i // complex128

		Try changing the initial value of v in the example code and observe how its type is affected.
	*/
	testTypeInference()
	/*

		Constants

		Constants are declared like variables, but with the const keyword.

		Constants can be character, string, boolean, or numeric values.

		Constants cannot be declared using the := syntax.
	*/
	testConsts()
	/*
		Numeric Constants
		Numeric constants are high-precision values.
		An untyped constant takes the type needed by its context.
		Try printing needInt(Big) too.
		(An int can store at maximum a 64-bit integer, and sometimes less.)
	*/
	testNumericValues()

}
