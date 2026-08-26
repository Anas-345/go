package main

import "fmt"

func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	if quantity > amountInStock(productID) {
		return false, accountBalance
	}
	if accountBalance < calcPrice(productID, quantity) {
		return false, accountBalance
	}
	accountBalance -= calcPrice(productID, quantity)
	return true, accountBalance
}

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	switch productID {
	case "1":
		return 1.50
	case "2":
		return 2.25
	case "3":
		return 3.00
	case "4":
		return 1.00
	case "5":
		return 2.50
	case "6":
		return 8.99
	case "7":
		return 22.50
	case "8":
		return 50.00
	case "9":
		return 999.99
	default:
		return 0.00
	}
}

func amountInStock(productID string) int {
	switch productID {
	case "1":
		return 11
	case "2":
		return 25
	case "3":
		return 4
	case "4":
		return 6
	case "5":
		return 50
	case "6":
		return 2
	case "7":
		return 0
	case "8":
		return 99
	case "9":
		return 1
	default:
		return 0
	}
}

func main() {
	fmt.Println(placeOrder("1", 10, 1000))
}
