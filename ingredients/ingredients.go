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

	ranges := [][2]int{}
	freshIngredientCount := 0

	scanningRanges := true
	for scanner.Scan() {
		if scanningRanges {
			// line := string(append([]byte(nil), scanner.Text()...))
			line := scanner.Text()
			if len(line) == 0 {
				scanningRanges = false
				continue
			}
			parts := strings.Split(line, "-")
			range_start, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			range_end, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, [2]int{range_start, range_end})
			// fmt.Println(ranges[len(ranges)-1])

		} else {
			line := scanner.Text()
			ingredientId, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			for _, singleRange := range ranges {
				if ingredientId >= singleRange[0] && ingredientId <= singleRange[1] {
					freshIngredientCount++
					break
				}
			}
		}
	}
	fmt.Println(freshIngredientCount)
}
