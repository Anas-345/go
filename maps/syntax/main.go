package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return userMap, errors.New("Invalid Size")
	}
	for i, val := range names {
		userMap[val] = user{name: val, phoneNumber: phoneNumbers[i]}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func main() {
	fmt.Println(getUserMap([]string{"Anas", "Ali", "Ahmad"}, []int{123, 456, 789}))
}
