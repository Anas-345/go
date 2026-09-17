package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	slice := []float64{}
	for i := 0; i < len(costs); i++ {
		if day == costs[i].day {
			slice = append(slice, costs[i].value)
		}
	}
	return slice
}

func main() {
	fmt.Println(getDayCosts([]cost{{day: 2, value: 12.3}, {day: 3, value: 121.3}, {day: 2, value: 12.182}}, 2))
}
