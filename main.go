package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := ""
	fmt.Print("Enter a sentence: ")
	_, _ = fmt.Scanln(&input)

	res, err := isValid(input)
	if res {
		fmt.Println("Valid sentence")
	} else {
		fmt.Println("Invalid sentence:", err)
	}
}

func isValid(sentence string) (bool, error) {
	if len(sentence) == 0 {
		return false, errors.New("sentence is empty")
	}

	if !isUppercase(sentence) {
		return false, errors.New("sentence doesn't start with a capital letter")
	}

	if !hasEvenQuotes(sentence) {
		return false, errors.New("sentence has an odd number of quotation marks")
	}

	if !hasValidEnding(sentence) {
		return false, errors.New("sentence doesn't end with a valid termination character")
	}

	if !hasValidPeriod(sentence) {
		return false, errors.New("sentence contains a period character other than the last character")
	}

	if !hasValidNumerics(sentence) {
		return false, errors.New("sentence contains numeric values less than 13")
	}

	// if all checks pass, return true & nil error
	return true, nil
}

// check if string starts with a capital
func isUppercase(sentence string) bool {
	// ensure no index out of range runtime error if empty string
	if len(sentence) <= 0 {
		return false
	}

	return unicode.IsUpper(rune(sentence[0]))
}

// check if string has an even number of quotes
func hasEvenQuotes(sentence string) bool {
	return strings.Count(sentence, "\"")%2 == 0
}

// check if string ends with a valid termination character
func hasValidEnding(sentence string) bool {
	punctuation := []string{".", "?", "!"}
	for _, p := range punctuation {
		if strings.HasSuffix(sentence, p) {
			return true
		}
	}

	return false
}

// check if string contains a period character other than the last character
func hasValidPeriod(sentence string) bool {
	// ensure no slice bounds out of range runtime error if empty string
	if len(sentence) <= 0 {
		return false
	}

	return strings.Count(sentence[:len(sentence)-1], ".") == 0
}

// check if string contains numeric values less than 13
func hasValidNumerics(sentence string) bool {
	// remove punctuation characters from sentence
	punctuation := []string{",", ".", "?", "!"}
	for _, p := range punctuation {
		sentence = strings.Replace(sentence, p, "", -1)
	}

	// split sentence into words
	for _, word := range strings.Fields(sentence) {
		// parse word to int & check if less than 13 provided there is no error
		if num, err := strconv.Atoi(word); err == nil && num < 13 {
			return false
		}
	}

	return true
}
