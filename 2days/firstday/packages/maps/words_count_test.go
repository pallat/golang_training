package main

import "testing"
import "reflect"

func TestLetterCountEachWord(t *testing.T) {
	sentenses := []string{
		"I am learning Go!",
		"The quick brown fox jumped over the lazy dog.",
	}

	expected := []map[string]int{
		map[string]int{
			"I":        1,
			"am":       2,
			"learning": 8,
			"Go!":      3,
		},
		map[string]int{
			"The":    3,
			"quick":  5,
			"brown":  5,
			"fox":    3,
			"jumped": 6,
			"over":   4,
			"the":    3,
			"lazy":   4,
			"dog.":   4,
		},
	}

	if !reflect.DeepEqual(LettersCount(sentenses[0]), expected[0]) {
		t.Error(LettersCount(sentenses[0]), expected[0])
	}
	if !reflect.DeepEqual(LettersCount(sentenses[1]), expected[1]) {
		t.Error(LettersCount(sentenses[1]), expected[1])
	}
}
