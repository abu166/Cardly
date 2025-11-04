package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func issueProcess() error {
	defer handleDeferError()

	fmt.Println("Processing issues...")
	// Simulate issue processing
	return errors.New("issue processing encountered an error")
}

func reportIssuance(brand, issuer string, brands, issuers map[string]string) {
	// collect all prefixes for the specified brand
	var brandPrefixes []string
	for prefix, name := range brands {
		if name == brand {
			brandPrefixes = append(brandPrefixes, prefix)
		}
	}

	// check if any brand prefixes were found
	if len(brandPrefixes) == 0 {
		fmt.Fprintf(os.Stderr, "Error: Brand %s not found.\n", brand)
		os.Exit(1)
		return
	}

	// find the prefix for the specified issuer
	var issuerPrefix string
	for prefix, name := range issuers {
		if name == issuer {
			issuerPrefix = prefix
			break
		}
	}

	// check if issuer prefix was found
	if issuerPrefix == "" {
		fmt.Fprintf(os.Stderr, "Error: Issuer %s not found.\n", issuer)
		os.Exit(1)
		return
	}

	// check if any brand prefix matches the issuer prefix
	compatible := false
	for _, brandPrefix := range brandPrefixes {
		if strings.HasPrefix(issuerPrefix, brandPrefix) {
			compatible = true
			break
		}
	}

	if !compatible {
		fmt.Fprintf(os.Stderr, "Error: Issuer %s is not compatible with brand %s.\n", issuer, brand)
		os.Exit(1)
		return
	}

	// generate a valid card number
	number, err := IssueCard(brandPrefixes, issuerPrefix)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to issue card for brand %s and issuer %s: %v\n", brand, issuer, err)
		os.Exit(1)
		return
	}

	fmt.Println(number)
}

func IssueCard(brandPrefixes []string, issuerPrefix string) (string, error) {
	for {
		number := issuerPrefix + fmt.Sprintf("%08d", rand.Intn(100000000))
		for _, brandPrefix := range brandPrefixes {
			if strings.HasPrefix(number, brandPrefix) && LuhnCheck(number) {
				return number, nil
			}
		}
	}
	return "", fmt.Errorf("could not generate a valid card number")
}
