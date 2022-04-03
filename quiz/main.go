package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer")
	flag.Parse()
	// _ = csvFileName

	file, err := os.Open(*csvFileName)
	if err != nil {
		// fmt.Printf("Failed to open the CSV file: %s\n", *csvFileName)
		// os.Exit(1)
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	// _ = file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	// fmt.Println(lines)
	problems := parseLines(lines)
	fmt.Println(problems)

	correctCount := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("Correct!")
			correctCount++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctCount, len(problems))
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type Problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
