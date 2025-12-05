package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("input", "input.txt", "input file")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	const substringCount int = 2

	niceStrings := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		// fmt.Printf("Checking string %s\n", line)

		var twoCharsAgo rune = 0
		var previousChar rune = 0

		smallerLength := len(line) - 3
		substringsSatisfied := false
		satisfiesTwoChar := false
		for index, char := range line {
			if !substringsSatisfied && index < smallerLength {
				twoLetter := line[index : index+2]
				count := strings.Count(line, twoLetter)
				if count >= substringCount {
					// fmt.Printf("In string %s\n", line)
					// fmt.Printf("Pair that appeared %d times - %s\n",
					// 	count, twoLetter)
					substringsSatisfied = true
				}
			}
			if char == twoCharsAgo {
				satisfiesTwoChar = true
			}
			twoCharsAgo = previousChar
			previousChar = char
		}
		// fmt.Printf("%s:\n", line)
		// fmt.Printf("%s: substring - %t, two char - %t\n",
		// 	line, substringsSatisfied, satisfiesTwoChar)
		if substringsSatisfied && satisfiesTwoChar {
			niceStrings++
			// fmt.Println(line)
		}
	}
	fmt.Println(niceStrings)
}
