package concurrency

import (
	"fmt"
	"time"

)
func Main()  {
	/*
	Goroutines
	A goroutine is a lightweight thread managed by the Go runtime.
	go f(x, y, z)
	starts a new goroutine running
	f(x, y, z)
	The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
	Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)
	*/
	go say("world")
	say("hello")
	/*
	Channels
	Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
	ch <- v    // Send v to channel ch.
	v := <-ch  // Receive from ch, and
	           // assign value to v.
	(The data flows in the direction of the arrow.)
	Like maps and slices, channels must be created before use:
	ch := make(chan int)
	By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
	The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.
	*/
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
