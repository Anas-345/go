package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msg := [3]string{primary, secondary, tertiary}
	cost := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}
	return msg, cost
}

func main() {
	msg, cost := getMessageWithRetries("Hello", "World", "!")
	fmt.Println(msg, cost)
}
