package main

import "fmt"

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func analyzeMessage(a *Analytics, m Message) {
	a.MessagesTotal++
	if m.Success {
		a.MessagesSucceeded++
	} else {
		a.MessagesFailed++
	}
}

func main() {
	a := Analytics{
		MessagesTotal:     5,
		MessagesFailed:    2,
		MessagesSucceeded: 3,
	}
	analyzeMessage(&a, Message{Success: true})
	fmt.Println(a)
}
