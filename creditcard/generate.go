package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
)

func GenerateCard(hidden string) ([]string, error) {
	result := []string{}

	// Check length
	if len(hidden) < 13 || len(hidden) > 19 {
		return nil, fmt.Errorf("invalid card number length, must be between 13 and 19 characters")
	}

	// Count number of asterisks at the end
	stars := 0
	for i := len(hidden) - 1; i >= 0 && hidden[i] == '*'; i-- {
		stars++
	}

	// Check if the number of asterisks is valid
	if stars == 0 || stars > 4 {
		return nil, fmt.Errorf("number of hidden symbols (*) must be between 1 and 4")
	}

	// Validate that the portion before the asterisks contains only digits
	original := hidden[:len(hidden)-stars]
	for _, r := range original {
		if !unicode.IsDigit(r) {
			return nil, fmt.Errorf("card number contains invalid characters")
		}
	}

	// Determine how many combinations to generate based on the number of asterisks
	required := 1
	for i := 0; i < stars; i++ {
		required *= 10
	}

	// Generate possible card numbers
	for i := 0; i < required; i++ {
		number := original + fmt.Sprintf("%0*d", stars, i) // Replace asterisks with numbers
		if LuhnCheck(number) {
			result = append(result, number) // Add to result if it passes Luhn's algorithm
		}
	}

	return result, nil
}

func handleGeneration(options string, pick bool) {
	numbers, err := GenerateCard(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if pick && len(numbers) > 0 { // if pick is true and numbers isn't empty
		rand.Seed(time.Now().UnixNano()) // seeds random num generator with current time
		// every time program runs, a different sequence is produced
		fmt.Println(numbers[rand.Intn(len(numbers))]) // print value wuth random index between 0 and size of "numbers"
	} else {
		for _, number := range numbers {
			fmt.Println(number)
		}
	}
}
