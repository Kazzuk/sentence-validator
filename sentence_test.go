package main

import "testing"

func TestEmptySentence(t *testing.T) {
	sentence := ""
	if result, _ := isValid(sentence); result {
		t.Errorf("isValid(%s) = %t, want false", sentence, result)
	}
}

func TestValidSentences(t *testing.T) {
	sentences := []string{
		"The quick brown fox said \"hello Mr lazy dog\".",
		"The quick brown fox said hello Mr lazy dog.",
		"One lazy dog is too few, 13 is too many.",
		"One lazy dog is too few, thirteen is too many.",
		"How many \"lazy dogs\" are there?",
	}

	for _, sentence := range sentences {
		if result, _ := isValid(sentence); !result {
			t.Errorf("isValid(%s) = %t, want true", sentence, result)
		}
	}
}

func TestInValidSentences(t *testing.T) {
	sentences := []string{
		"The quick brown fox said \"hello Mr. lazy dog\".",
		"the quick brown fox said \"hello Mr lazy dog\".",
		"\"The quick brown fox said \"hello Mr lazy dog.\"",
		"One lazy dog is too few, 12 is too many.",
		"Are there 11, 12, or 13 lazy dogs?",
		"There is no punctuation in this sentence",
	}

	for _, sentence := range sentences {
		if result, _ := isValid(sentence); result {
			t.Errorf("isValid(%s) = %t, want false", sentence, result)
		}
	}
}
