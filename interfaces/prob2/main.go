package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sa systemAlert) importance() int {
	return 100
}

type empty struct{}

func (e empty) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch t := n.(type) {
	case directMessage:
		return t.senderUsername, t.importance()
	case groupMessage:
		return t.groupName, t.importance()
	case systemAlert:
		return t.alertCode, t.importance()
	default:
		return "", 0
	}
}

func main() {
	d := directMessage{
		senderUsername: "Anas",
		priorityLevel:  70,
		isUrgent:       true,
		messageContent: "Hello",
	}
	gm := groupMessage{
		groupName:      "Group",
		priorityLevel:  75,
		messageContent: "Welcome",
	}
	sa := systemAlert{
		alertCode:      "101",
		messageContent: "Urgent",
	}
	fmt.Println(processNotification(d))
	fmt.Println(processNotification(gm))
	fmt.Println(processNotification(sa))
	fmt.Println(processNotification(empty{}))
}
