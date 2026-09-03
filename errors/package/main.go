package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("no dividing by 0")
		return 0.0, err
	}
	return x / y, nil
}

func main() {
	val, err := divide(30, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
