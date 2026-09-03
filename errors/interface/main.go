package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by zero", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main() {
	division, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The result is", division)
}
