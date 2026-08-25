package main

import "fmt"

func reformat(message string, formatter func(string) string) string {
	message = formatter(message)
	return "TEXTIO: " + message
}

func main() {
	addPeriod := func(message string) string { return message + "." }
	addExclamation := func(message string) string { return message + "!" }
	fmt.Println(reformat("Add Period", addPeriod))
	fmt.Println(reformat("Add Exclamation", addExclamation))
}
