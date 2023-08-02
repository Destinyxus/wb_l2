package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	anagramSets := searchAn([]string{"пятак", "пятка", "Тяпка"})
	fmt.Println(anagramSets)
}

func searchAn(arr []string) map[string][]string {
	anagramSets := make(map[string][]string)

	// Create a map where the key is the sorted characters of the word (in lowercase),
	// and the value is the word itself (in lowercase)
	anagramMap := make(map[string]string)
	for _, word := range arr {
		lowercaseWord := strings.ToLower(word)
		sortedWord := sortString(lowercaseWord)
		anagramMap[sortedWord] = lowercaseWord
	}
	// Group anagrams together in the result map
	for _, word := range arr {
		lowercaseWord := strings.ToLower(word)
		sortedWord := sortString(lowercaseWord)
		anagramKey := anagramMap[sortedWord]
		fmt.Println(anagramKey)
		// Append the word to the set only if it's the first occurrence of the anagram
		if _, ok := anagramSets[anagramKey]; !ok {
			// Create a new slice and copy the word into it
			anagramSets[anagramKey] = []string{lowercaseWord}
		} else {
			// Append the word to the existing slice
			anagramSets[anagramKey] = append(anagramSets[anagramKey], lowercaseWord)
		}
	}

	return anagramSets
}

// Helper function to sort the characters of a word and return the sorted string
func sortString(str string) string {
	sortedRunes := []rune(str)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}
