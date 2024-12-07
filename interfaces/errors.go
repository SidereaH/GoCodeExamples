package interfaces

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %d", float64(e))
}


func Sqrt(x float64) (float64, error) {
	if x < 0{
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	prev := 0.0
	for i := 0; i < 11; i++ {

		z -= (z*z - x) / (2 * z)
		if math.Abs(prev-z) < 1e-14 {
			return prev, nil
		}
		prev = z

	}
	return z, nil
}