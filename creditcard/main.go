package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Application crashed: %v", r)
		}
	}()

	// Check if a feature is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./creditcard <feature> [options]")
		os.Exit(1)
	}

	// Find the feature
	feature := os.Args[1]

	// Shift the args for flag parsing (removes the feature argument)
	os.Args = os.Args[1:]

	// Execute the chosen feature
	switch feature {
	case "validate":
		handleValidate()
	case "generate":
		handleGenerate()
	case "information":
		handleInformation()
	case "issue":
		handleIssue()
	default:
		fmt.Printf("Unknown feature: %s\n", feature)
		os.Exit(1)
	}
}

func runApplication() error {
	fmt.Println("Running common pre-execution tasks...")

	return nil
}

func handleValidate() {
	// define flags
	stdinFlag := flag.Bool("stdin", false, "Read input from stdin")
	flag.Parse()

	var cardNumbers []string
	if *stdinFlag {
		cardNumbers = readFromStdin()
	} else {
		cardNumbers = flag.Args()
	}

	if len(cardNumbers) == 0 {
		fmt.Fprintln(os.Stderr, "No card numbers provided.")
		os.Exit(1)
	}

	// call card validation function
	cardValidation(cardNumbers)
}

func handleGenerate() {
	// define flags
	pickFlag := flag.Bool("pick", false, "Pick a random valid card number")
	flag.Parse()
	pattern := flag.Arg(0)

	if len(flag.Args()) != 1 && pattern != "*" {
		fmt.Fprintln(os.Stderr, "Usage of flag pick: ./creditcard generate [--pick] <card-pattern>")
		os.Exit(1)
	}
	// pattern := flag.Arg(0)

	handleGeneration(pattern, *pickFlag)
}

func handleInformation() {
	// define flags for brands and issuers files
	brandsFile := flag.String("brands", "", "Path to brands.txt file")
	issuersFile := flag.String("issuers", "", "Path to issuers.txt file")
	stdinFlag := flag.Bool("stdin", false, "Read card numbers from stdin")
	flag.Parse()

	// check if both brands and issuers files are provided
	if *brandsFile == "" || *issuersFile == "" {
		fmt.Fprintln(os.Stderr, "Error: --brands and --issuers flags are required.")
		os.Exit(1)
	}

	// load brands and issuers data
	brands, err := LoadData(*brandsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load brands from %s: %v\n", *brandsFile, err)
		os.Exit(1)
	}

	issuers, err := LoadData(*issuersFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load issuers from %s: %v\n", *issuersFile, err)
		os.Exit(1)
	}

	// Read card numbers (from stdin or command-line arguments)
	var cardNumbers []string
	if *stdinFlag {
		cardNumbers = readFromStdin()
	} else {
		cardNumbers = flag.Args()
	}

	// Check if card numbers are provided
	if len(cardNumbers) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No card numbers provided.")
		os.Exit(1)
	}

	// Display card information
	for _, number := range cardNumbers {
		brand, issuer, valid := DisplayCardInfo(number, brands, issuers)
		fmt.Println(number)
		fmt.Println("Correct:", valid)
		fmt.Println("Card Brand:", brand)
		fmt.Println("Card Issuer:", issuer)
	}
}

func handleIssue() {
	// define flags
	brandsFile := flag.String("brands", "", "Path to brands.txt")
	issuersFile := flag.String("issuers", "", "Path to issuers.txt")
	brand := flag.String("brand", "", "Card brand")
	issuer := flag.String("issuer", "", "Card issuer")
	flag.Parse()

	// checks if required flags are present
	if *brandsFile == "" || *issuersFile == "" || *brand == "" || *issuer == "" {
		fmt.Fprintln(os.Stderr, "Error: --brands and --issuers flags are required.")
		os.Exit(1)
	}

	brands, err := LoadData(*brandsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load brands: %v\n", err)
		os.Exit(1)
	}

	issuers, err := LoadData(*issuersFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load issuers: %v\n", err)
		os.Exit(1)
	}

	// call the issuance function
	reportIssuance(*brand, *issuer, brands, issuers)
}

func readFromStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []string
	for scanner.Scan() {
		line := scanner.Text()
		// split the line into individual card numbers by whitespace
		cardNumbers := strings.Fields(line)
		inputs = append(inputs, cardNumbers...)
	}
	return inputs
}
