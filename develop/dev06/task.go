package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	flagF     string
	flagD     string
	flagS     bool
	fieldNums []int
)

func main() {
	flag.BoolVar(&flagS, "s", false, "only-delimited")
	flag.StringVar(&flagD, "d", "\t", "delimiter")
	flag.StringVar(&flagF, "f", "", "field")
	flag.Parse()

	lines := readLines()

	parseFieldNumbers()

	if len(flagF) > 0 && flagD != "\t" {
		cutByDelimiter(lines, flagD)
	} else {
		fmt.Println(strings.Join(lines, "\n"))
	}

}

func cutByDelimiter(arr []string, delimiter string) {
	var fields []string
	for _, line := range arr {

		parts := strings.Split(line, delimiter)

		if flagS && !strings.Contains(line, delimiter) {
			continue
		}
		var extractedFields []string

		for _, fieldNum := range fieldNums {

			if fieldNum > 0 && fieldNum <= len(parts) {

				extractedFields = append(extractedFields, strings.TrimSpace(parts[fieldNum-1]))
			}

		}
		if len(extractedFields) > 0 {
			fields = append(fields, strings.Join(extractedFields, ", "))
		}
	}

	if len(fields) > 0 {
		fmt.Println(strings.Join(fields, "\n"))
	}
}

func parseFieldNumbers() {
	fields := strings.Split(flagF, ",")
	for _, field := range fields {
		num := strings.TrimSpace(field)
		if num != "" {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Invalid field number:", err)
				os.Exit(1)
			}
			fieldNums = append(fieldNums, n)
		}
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
