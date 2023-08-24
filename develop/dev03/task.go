package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Line struct {
	Text     string
	SortKey  interface{}
	Original string
}

type BySortKey []Line

func (a BySortKey) Len() int           { return len(a) }
func (a BySortKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySortKey) Less(i, j int) bool { return less(a[i].SortKey, a[j].SortKey) }

func less(a, b interface{}) bool {
	switch aVal := a.(type) {
	case string:
		bVal := b.(string)
		return aVal < bVal
	case int:
		bVal := b.(int)
		return aVal < bVal
	default:
		return false // Unsupported type
	}
}

var (
	flagColumn  int
	flagNumeric bool
	flagReverse bool
	flagUnique  bool
)

func main() {
	flag.IntVar(&flagColumn, "k", 1, "Column number for sorting (1-indexed)")
	flag.BoolVar(&flagNumeric, "n", false, "Sort by numeric value")
	flag.BoolVar(&flagReverse, "r", false, "Sort in reverse order")
	flag.BoolVar(&flagUnique, "u", false, "Remove duplicate lines")
	flag.Parse()

	lines := readLines()

	var linesToSort []Line
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= flagColumn {
			sortKey := fields[flagColumn-1]
			if flagNumeric {
				if num, err := strconv.Atoi(sortKey); err == nil {
					linesToSort = append(linesToSort, Line{Text: line, SortKey: num, Original: line})
				} else {
					fmt.Fprintf(os.Stderr, "Line '%s' cannot be converted to a number\n", sortKey)
				}
			} else {
				linesToSort = append(linesToSort, Line{Text: line, SortKey: sortKey, Original: line})
			}
		} else {
			fmt.Fprintf(os.Stderr, "Line '%s' doesn't have enough columns for sorting\n", line)
		}
	}

	sort.Sort(BySortKey(linesToSort))

	if flagReverse {
		reverseSlice(linesToSort)
	}

	if flagUnique {
		linesToSort = removeDuplicates(linesToSort)
	}

	for _, sortedLine := range linesToSort {
		fmt.Println(sortedLine.Original)
	}
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

func reverseSlice(slice []Line) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func removeDuplicates(lines []Line) []Line {
	seen := make(map[interface{}]struct{})
	var uniqueLines []Line
	for _, line := range lines {
		if _, exists := seen[line.SortKey]; !exists {
			uniqueLines = append(uniqueLines, line)
			seen[line.SortKey] = struct{}{}
		}
	}
	return uniqueLines
}
