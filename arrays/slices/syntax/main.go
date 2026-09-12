package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro" {
		return messages[:], nil
	} else if plan == "free" {
		return messages[:2], nil
	} else {
		return nil, errors.New("Unsupported plan")
	}
}

func main() {
	res, err := getMessageWithRetriesForPlan(planFree, [3]string{"Hello", "Welcome", "Back"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
