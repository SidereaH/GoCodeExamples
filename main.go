package gofix

import (
	"concurrency"
	"flowcontrlstate"
	"interfaces"
	"methods"
	"pacvarfun"
	"types"
)

func main() {
	//package, variable, functions

	pacvarfun.Main()
	//if else, for while
	flowcontrlstate.Main()
	//types of values and structures
	types.Main()
	//methods
	methods.Main()
	//interfaces
	interfaces.Main()
	concurrency.Main()
}
