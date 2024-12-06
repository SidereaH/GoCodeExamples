module main

replace (
	flowcontrlstate => ./basics/flowcontrlstate/
	interfaces => ./interfaces
	methods => ./methods
	pacvarfun => ./basics/pacvarfun/
	types => ./basics/types
)

go 1.23

require (
	flowcontrlstate v0.0.0-00010101000000-000000000000
	interfaces v0.0.0-00010101000000-000000000000
	methods v0.0.0-00010101000000-000000000000
	pacvarfun v0.0.0-00010101000000-000000000000
	types v0.0.0-00010101000000-000000000000
)

require golang.org/x/tour v0.1.0 // indirect
