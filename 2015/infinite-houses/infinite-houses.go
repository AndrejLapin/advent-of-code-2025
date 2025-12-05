package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Position struct {
	x int
	y int
}

func visitingUniqueHouse(position *Position, visitedPlaces *map[int]map[int]bool) bool {
	row, okRow := (*visitedPlaces)[position.x]
	if okRow {
		_, okLocation := row[position.y]
		if okLocation {
			return false
		} else {
			(*visitedPlaces)[position.x][position.y] = true
			return true
		}
	} else {
		(*visitedPlaces)[position.x] = make(map[int]bool)
		(*visitedPlaces)[position.x][position.y] = true
		return true
	}
}

func main() {
	fileName := flag.String("input", "input.txt", "input file")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	visitedPlaces := make(map[int]map[int]bool)
	visitedPlaces[0] = make(map[int]bool)
	visitedPlaces[0][0] = true
	uniqueHousesVisited := 1
	position := Position{x: 0, y: 0}
	for {
		readByte, err := reader.ReadByte()
		if err != nil {
			break
		}
		switch readByte {
		case '^':
			position.y++
		case 'v':
			position.y--
		case '<':
			position.x--
		case '>':
			position.x++
		}
		if visitingUniqueHouse(&position, &visitedPlaces) {
			uniqueHousesVisited++
		}
	}
	fmt.Println(uniqueHousesVisited)
}
