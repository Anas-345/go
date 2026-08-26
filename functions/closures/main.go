package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(val int) int { sum += val; return sum }
}

func main() {
	customer1 := adder()
	customer2 := adder()
	customer1(5)
	customer1(5)
	fmt.Println(customer1(3))
	customer2(10)
	customer2(50)
	fmt.Println(customer2(1))
}
