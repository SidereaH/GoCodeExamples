package types

import (
	"fmt"
	"strings"
)

type Vertexx struct {
	Lat, Long float64
}

var m map[string]Vertexx

/*
Map literals
Map literals are like struct literals, but the keys are required.
*/
var m2 = map[string]Vertexx{

	"Bell Labs": Vertexx{
		40.68433, -74.39967,
	},
	"Google": Vertexx{
		37.42202, -122.08408,
	},
}

/*
Map literals continued
If the top-level type is just a type name, you can omit it from the elements of the literal.
*/
var m3 = map[string]Vertexx{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func simpleMapPrint() {
	m = make(map[string]Vertexx)
	m["Bell Labs"] = Vertexx{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2)
}

// mutations
func printMutations() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
You might find strings.Fields helpful.
*/
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	strings := strings.Fields(s)
	for _, s := range strings {
		countString, ok := m[s]
		if ok == true {
			m[s] = countString + 1
		} else {
			m[s] = 1
		}
	}
	return m
}
