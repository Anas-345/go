package main

import (
	"fmt"
)

func getNameCounts(names []string) map[rune]map[string]int {
	userCount := make(map[rune]map[string]int)
	for _, val := range names {
		firstChar := rune(val[0])
		if _, ok := userCount[firstChar]; !ok {
			userCount[firstChar] = make(map[string]int)
		}
		userCount[firstChar][val]++
	}
	return userCount
}

func main() {
	names := []string{"Alice", "Bob", "Alice", "Charlie", "Anna"}
	counts := getNameCounts(names)
	fmt.Println(counts)
}
