package main

import "fmt"

type payment interface {
	pay(amount float64) string
	getBalance() float64
}

type creditCard struct {
	cardHolder  string
	balance     float64
	creditLimit float64
}

func (c creditCard) pay(amount float64) string {
	if amount > c.creditLimit {
		return "Credit limit exceeded"
	}
	c.balance -= amount
	return fmt.Sprintf("Paid $%.2f via Credit Card", amount)
}

func (c creditCard) getBalance() float64 {
	return c.balance
}

type bankTransfer struct {
	accountHolder string
	balance       float64
}

func (b bankTransfer) pay(amount float64) string {
	if amount > b.balance {
		return "Insufficient funds"
	}
	b.balance -= amount
	return fmt.Sprintf("Paid $%.2f via bank transfer", amount)
}

func (b bankTransfer) getBalance() float64 {
	return b.balance
}

type crypto struct {
	walletAddress string
	balance       float64
	coinType      string
}

func (c crypto) pay(amount float64) string {
	if amount > c.balance {
		return "Insufficient Crypto"
	}
	c.balance -= amount
	return fmt.Sprintf("Paid $%.2f via %s", amount, c.coinType)
}

func (c crypto) getBalance() float64 {
	return c.balance
}

func processPayment(p payment, amount float64) string {
	switch t := p.(type) {
	case creditCard:
		return fmt.Sprintf("%s | Remaining credit: $%.2f", t.pay(amount), t.getBalance())
	case bankTransfer:
		return fmt.Sprintf("%s | Remaining balance: $%.2f", t.pay(amount), t.getBalance())
	case crypto:
		return fmt.Sprintf("%s | Remaining coins: $%.2f", t.pay(amount), t.getBalance())
	default:
		return "Please enter valid data and valid transfer type"
	}
}

func main() {
	cc := creditCard{
		cardHolder:  "Anas",
		balance:     10000,
		creditLimit: 8000,
	}

	bt := bankTransfer{
		accountHolder: "Anas Munir",
		balance:       5000,
	}

	cr := crypto{
		walletAddress: "0x123abc",
		balance:       2.5,
		coinType:      "Bitcoin",
	}

	cc2 := creditCard{
		cardHolder:  "Ali",
		balance:     10000,
		creditLimit: 3000,
	}

	bt2 := bankTransfer{
		accountHolder: "Ahmed",
		balance:       500,
	}

	cr2 := crypto{
		walletAddress: "0x456def",
		balance:       0.5,
		coinType:      "Ethereum",
	}

	fmt.Println(processPayment(cc, 5000))
	fmt.Println(processPayment(bt, 2000))
	fmt.Println(processPayment(cr, 1.0))
	fmt.Println("---")
	fmt.Println(processPayment(cc2, 5000))
	fmt.Println(processPayment(bt2, 1000))
	fmt.Println(processPayment(cr2, 1.0))
}
