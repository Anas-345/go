package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, v := range msg {
		for _, el := range badWords {
			if el == v {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(indexOfFirstBadWord([]string{"Hello", "World"}, []string{"World"}))
}
