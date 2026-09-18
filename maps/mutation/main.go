package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (bool, error) {
	el, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !el.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	fmt.Println(deleteIfNecessary(map[string]user{"Anas": {name: "Anas", number: 123, scheduledForDeletion: false}}, "Anas"))
}
