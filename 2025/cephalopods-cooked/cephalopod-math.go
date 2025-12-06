package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

	numbers := []string{}
	var operations string

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '*' || line[0] == '+' {
			operations = line
		} else {
			numbers = append(numbers, line)
		}
	}

	total := 0
	operatorPos := 0
	actualOperatorPos := operatorPos
	reducedString := operations
	for operatorPos != -1 {
		currentOperatorPos := actualOperatorPos
		operationResult := 0
		var operation func(*int, int)
		if operations[currentOperatorPos] == '*' {
			operationResult = 1
			operation = func(a *int, b int) {
				*a *= b
			}
		} else {
			operation = func(a *int, b int) {
				*a += b
			}
		}

		reducedString = reducedString[operatorPos+1:]
		operatorPos = strings.IndexAny(reducedString, "*+")
		actualOperatorPos += operatorPos + 1
		columnEnd := actualOperatorPos - 1
		if operatorPos == -1 {
			columnEnd = len(operations)
		}

		for columnIndex := currentOperatorPos; columnIndex < columnEnd; columnIndex++ {
			digits := []byte{}
			for _, row := range numbers {
				digits = append(digits, row[columnIndex])
			}
			number, _ := strconv.Atoi(strings.TrimSpace(string(digits)))
			operation(&operationResult, number)
		}
		total += operationResult
	}
	fmt.Println(total)
}
