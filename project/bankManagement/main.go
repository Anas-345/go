package main

import (
	"fmt"
)

type bankAccount struct {
	number  int
	balance float64
}

type bankFunc interface {
	deposit()
	withdraw()
}

func (b *bankAccount) deposit() {
	var dep int
	fmt.Print("Enter amount you want to deposit: ")
	_, err := fmt.Scan(&dep)
	if err != nil {
		fmt.Println("Invalid value")
	} else {
		b.balance += float64(dep)
	}
}

func (b *bankAccount) withdraw() {
	var am int
	fmt.Print("Enter amount you want to withdraw: ")
	_, err := fmt.Scan(&am)
	if err != nil {
		fmt.Println("Invalid value")
	} else if am > int(b.balance) {
		fmt.Println("Insufficient balance")
	} else {
		b.balance -= float64(am)
	}
}

func createAccount() (float64, error) {
	var initialAmount float64
	fmt.Print("Enter initial amount of your account: ")
	_, err := fmt.Scan(&initialAmount)
	if err != nil {
		return 0.0, err
	}
	return initialAmount, nil
}

func menu() int {
	var choice int
	fmt.Println("1-Create new account")
	fmt.Println("2- See details")
	fmt.Println("3- Deposit amount")
	fmt.Println("4- Withdraw amount")
	fmt.Println("5- Exit")
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid value")
	}
	return choice
}

func main() {
	accounts := []bankAccount{}
	isContinue := true
	for isContinue {
		choice := menu()
		switch choice {
		case 1:
			initialAmount, err := createAccount()
			if err != nil {
				fmt.Println(err)
				continue
			}
			s := bankAccount{
				number:  (len(accounts) + 1),
				balance: initialAmount,
			}
			accounts = append(accounts, s)
			fmt.Printf(`Successfully made with account no. %d and with balance %.2f`, s.number, s.balance)
			fmt.Println()
		case 2:
			fmt.Println(accounts[0].number, accounts[0].balance)
		case 3:
			accounts[0].deposit()
		case 4:
			accounts[0].withdraw()
		case 5:
			isContinue = false
		default:
			fmt.Println("Invalid input")
		}
	}
	fmt.Println(accounts[0].number, accounts[0].balance)

}
