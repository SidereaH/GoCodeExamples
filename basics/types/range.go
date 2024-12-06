package types

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func rangeFromLoop() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
func rangeContinued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
