package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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
	ribbonTotal := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		dimensionsStrings := strings.Split(line, "x")
		dimensions := [3]int{}
		volume := 1
		for index, value := range dimensionsStrings {
			dimensions[index], _ = strconv.Atoi(value)
			volume *= dimensions[index]
		}
		smallestPerimiter := math.MaxInt
		sides := [3]int{} // i don't think we even need this
		for index, _ := range sides {
			sides[index] = 2*dimensions[index] + 2*dimensions[(index+1)%len(dimensions)]
			if sides[index] < smallestPerimiter {
				smallestPerimiter = sides[index]
			}
			// wrappingPaperTotal += sides[index] * 2
		}
		ribbonTotal += smallestPerimiter + volume
	}
	fmt.Println(ribbonTotal)
}
