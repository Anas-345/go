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

func (b *bankAccount) withdraw() error {
	var am int
	fmt.Print("Enter amount you want to withdraw: ")
	_, err := fmt.Scan(&am)
	if err != nil {
		return fmt.Errorf("Invalid input")
	} else if am > int(b.balance) {
		return fmt.Errorf("Insufficient balance")
	} else {
		b.balance -= float64(am)
		return nil
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
	accounts := map[int]bankAccount{}
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
			num := len(accounts)
			accounts[num] = bankAccount{
				number:  num,
				balance: initialAmount,
			}
			fmt.Printf(`Successfully made with account no. %d and with balance %.2f`, accounts[num].number, accounts[num].balance)
			fmt.Println()
		case 2:
			fmt.Println(accounts[0].number, accounts[0].balance)
		case 3:
			fmt.Print("Enter your account number: ")
			var accNum int
			_, err := fmt.Scan(&accNum)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			temp := accounts[accNum]
			temp.deposit()
			accounts[accNum] = temp
		case 4:
			fmt.Print("Enter your account number: ")
			var accNum int
			_, err := fmt.Scan(&accNum)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			temp := accounts[accNum]
			err = temp.withdraw()
			if err != nil {
				fmt.Println(err)
				continue
			}
			accounts[accNum] = temp
		case 5:
			isContinue = false
		default:
			fmt.Println("Invalid input")
		}
	}
	for _, v := range accounts {
		fmt.Println(v.number, v.balance)
	}
}
