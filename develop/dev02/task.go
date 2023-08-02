package main

import (
	"strings"
	"unicode"
)

func unpackString(input string) string {
	if len(input) == 0 {
		return ""
	}

	stack := []rune{}
	i := 0

	for i < len(input) {
		currentChar := input[i]

		if currentChar == '\\' {
			// Handle escape sequences
			if i+1 < len(input) {
				nextChar := input[i+1]
				if nextChar == '\\' || unicode.IsDigit(rune(nextChar)) {
					i++
					stack = append(stack, rune(nextChar))
				} else {
					stack = append(stack, '\\')
				}
			} else {
				// If the backslash is at the end of the input, treat it as a literal backslash
				stack = append(stack, '\\')
			}
		} else if unicode.IsDigit(rune(currentChar)) {
			// Handle numeric counts
			count := int(currentChar - '0')
			for i+1 < len(input) && unicode.IsDigit(rune(input[i+1])) {
				i++
				count = count*10 + int(input[i]-'0')
			}
			if len(stack) > 0 {
				// Repeat the last character on the stack 'count' times
				lastChar := stack[len(stack)-1]
				stack = stack[:len(stack)-1] // Remove the last character from the stack
				stack = append(stack, []rune(strings.Repeat(string(lastChar), count))...)
			}
		} else {
			// Handle regular characters
			stack = append(stack, rune(currentChar))
		}

		i++
	}

	return string(stack)
}
