package main

import "fmt"

func getProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

func getProductInfo(tier string) (string, string, string) {
	switch tier {
	case "basic":
		return "1,000 texts per month", "$30 per month", "most popular"
	case "premium":
		return "50,000 texts per month", "$60 per month", "best value"
	case "enterprise":
		return "unlimited texts per month", "$100 per month", "customizable"
	default:
		return "", "", ""
	}
}

func main() {
	fmt.Println(getProductMessage("basic"))
}
