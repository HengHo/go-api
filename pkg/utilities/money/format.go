package money

import (
	"fmt"
	"strconv"
	"strings"
)

func Format(amount float64) string {
	return formatCurrency(amount)
}

func formatCurrency(amount float64) string {
	// Format the float value with two decimal places
	amountStr := strconv.FormatFloat(amount, 'f', 2, 64)

	// Split the amount into integer and decimal parts
	integerPartStr, decimalPartStr := splitAmount(amountStr)

	// Add commas to the integer part
	integerPartWithCommas := addCommas(integerPartStr)

	// Combine the formatted parts with a currency symbol
	formattedAmount := integerPartWithCommas + "." + decimalPartStr

	return formattedAmount
}

func AmountPart(amount float64) string {
	amountStr := fmt.Sprintf("%v", amount)
	parts := strings.Split(amountStr, ".")
	if len(parts) == 1 {
		return "0"
	}
	return parts[1]
}

func splitAmount(amountStr string) (string, string) {
	// Split the amount into integer and decimal parts
	parts := strings.Split(amountStr, ".")

	// Handle missing decimal part
	if len(parts) == 1 {
		return parts[0], "00"
	}

	return parts[0], parts[1]
}

func addCommas(amountStr string) string {
	// Reverse the string for easier insertion of commas
	reversed := reverseString(amountStr)

	// Insert commas after every three characters
	var result string
	for i, char := range reversed {
		if i > 0 && i%3 == 0 {
			result += ","
		}
		result += string(char)
	}

	// Reverse the string back to its original order
	return reverseString(result)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
