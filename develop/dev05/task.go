package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	flagA int
	flagB int
	flagF bool
	flagC int
)

func main() {

	flag.IntVar(&flagA, "A", 1, "after")
	flag.IntVar(&flagB, "B", 1, "before")
	flag.BoolVar(&flagF, "F", false, "fixed")
	flag.IntVar(&flagC, "C", 1, "context")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Enter the keyword with args")
		return
	}

	keyword := args[len(args)-1]

	lines := readLines()

	switch os.Args[len(os.Args)-3] {
	case "-A":
		fmt.Println(SearchWithLinesAfter(lines, keyword, flagA))
	case "-B":
		fmt.Println(SearchWithLinesBefore(lines, keyword, flagB))
	case "-C":
		before := SearchWithLinesBefore(lines, keyword, flagC)
		after := SearchWithLinesAfter(lines, keyword, flagC)
		before = append(before[:len(before)-1], after...)
		fmt.Println(before)
	}

}

func SearchWithLinesBefore(arr []string, keyword string, n int) []string {
	var foundLines []string
	var linesBefore []string
	var found bool
	//var before bool

	for i := 0; i < len(arr); i++ {
		if isFixed(arr[i], keyword) {
			found = true
			foundLines = append(foundLines, arr[i])

			if i == 0 {
				fmt.Println("No lines before")
				return foundLines
			}
			for j := i - n; j < i; j++ {
				linesBefore = append(linesBefore, arr[j])
				//before = true
			}
		}

	}
	if !found {
		fmt.Println("Keyword not found in the input.")
	}
	linesBefore = append(linesBefore, foundLines...)
	return linesBefore
}

func SearchWithLinesAfter(arr []string, keyword string, n int) []string {
	var foundLines []string
	var found bool

	for i := 0; i < len(arr); i++ {
		if isFixed(arr[i], keyword) {
			found = true
			foundLines = append(foundLines, arr[i])

			// Check if the next line contains the keyword
			nextIndex := i + 1
			for nextIndex < len(arr) && strings.Contains(arr[nextIndex], keyword) {
				nextIndex++
			}

			// Print lines after the current occurrence
			for j := 1; j <= n && i+j < len(arr); j++ {
				foundLines = append(foundLines, arr[i+j])
			}

			// Update i to the index after the last occurrence of keyword
			i = nextIndex - 1
		}
	}

	if !found {
		fmt.Println("Keyword not found in the input.")
	}

	return foundLines
}

func isFixed(s, substr string) bool {
	if flagF {
		return s == substr
	}
	return strings.Contains(s, substr)
}

func readLines() []string {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	return lines
}
