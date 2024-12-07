package methods

func Main() {
	/*
		Methods
		Go does not have classes. However, you can define methods on types.
		A method is a function with a special receiver argument.
		The receiver appears in its own argument list between the func keyword and the method name.
		In this example, the Abs method has a receiver of type Vertex named v.
	*/
	printSimpleMethod()
	/*
		Interfaces
		An interface type is defined as a set of method signatures.
		A value of interface type can hold any value that implements those methods.
		Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
	*/


}
