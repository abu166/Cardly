package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// checks the validity of a card number
func ValidateCardNumber(number string) (bool, string) {
	// check length
	if len(number) < 13 || len(number) > 19 {
		return false, "Card number must be between 13 and 19 digits."
	}

	// check if all characters are digits
	for _, r := range number {
		if !unicode.IsDigit(r) {
			return false, "Card number contains invalid characters."
		}
	}

	// check Luhn's algorithm check
	if !LuhnCheck(number) {
		return false, "Card number failed Luhn's algorithm validation."
	}

	// if all checks pass
	return true, ""
}

// handles the validation of multiple card numbers
func cardValidation(numbers []string) {
	for _, number := range numbers {
		// trim whitespace and newline characters
		trimmed := strings.TrimSpace(number)

		// validate card number
		valid, errMsg := ValidateCardNumber(trimmed)
		if valid {
			fmt.Println("OK")
		} else {
			// Print detailed error message to stderr
			fmt.Fprintf(os.Stderr, "INCORRECT: %s\n", errMsg)
			os.Exit(1)
		}
	}
}
