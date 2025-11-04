package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleDeferError() {
	if r := recover(); r != nil {
		fmt.Printf("Recovered from error: %v\n", r)
	}
}

func LuhnCheck(number string) bool {
	// algo to check if card number is valid
	sum := 0
	second := false

	for i := len(number) - 1; i >= 0; i-- {
		n := int(number[i] - '0')
		if second {
			// double every other digit, starting from the first
			n *= 2
			// if result ends up with 2 digits, add 2 digits together
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		second = !second
	}
	// if final sum is divisible by 10, then the card number is valid
	return sum%10 == 0
}

func LoadData(filename string) (map[string]string, error) {
	data := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2) // Split line into exactly 2 parts
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid file format")
		}
		data[parts[1]] = parts[0] // Key = Prefix, Value = Name
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file")
	}

	return data, nil
}
