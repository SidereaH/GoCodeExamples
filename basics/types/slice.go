package types

import (
	"fmt"
	"strings"
)

func PrintSimpleSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
func PrintReferncedSlices() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
func PrintSliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}
func PrintSliceDefaults() {

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}
func PrintSliceLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func PrintNilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
func printSliceWMake() {
	a := make([]int, 5)
	printSlicew("a", a)

	b := make([]int, 0, 5)
	printSlicew("b", b)

	c := b[:2]
	printSlicew("c", c)

	d := c[2:5]
	printSlicew("d", d)
}

func printSlicew(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSliceOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
func printAppendingSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}
func pic2(dx, dy int) [][]uint8 {
	//works only in tour app
	/*
	   Exercise: Slices
	   Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
	   The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
	   (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
	   (Use uint8(intValue) to convert between types.)
	*/
	// Создаем двумерный срез длиной dy
	slice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		// Инициализируем каждую строку длиной dx
		slice[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			// Заполняем значение пикселя (например, (x+y)/2)
			slice[i][j] = uint8((i + j) / 2)
		}
	}
	return slice
}
