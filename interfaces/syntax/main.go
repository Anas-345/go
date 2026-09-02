package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	m := msg.getMessage()
	return m, len(m) * 3
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	birthday := birthdayMessage{
		birthdayTime:  time.Date(2008, time.May, 14, 0, 0, 0, 0, time.UTC),
		recipientName: "Anas",
	}
	fmt.Println(sendMessage(birthday))
	fmt.Println(birthday.getMessage())
}
