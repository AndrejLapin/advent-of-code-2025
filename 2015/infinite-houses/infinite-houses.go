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
	positions := [2]Position{}
	for index, _ := range positions {
		positions[index] = Position{x: 0, y: 0}
	}
	turnIndex := 0
	for {
		readByte, err := reader.ReadByte()
		if err != nil {
			break
		}
		switch readByte {
		case '^':
			positions[turnIndex].y++
		case 'v':
			positions[turnIndex].y--
		case '<':
			positions[turnIndex].x--
		case '>':
			positions[turnIndex].x++
		}
		if visitingUniqueHouse(&positions[turnIndex], &visitedPlaces) {
			uniqueHousesVisited++
		}
		turnIndex = (turnIndex + 1) % len(positions)
	}
	fmt.Println(uniqueHousesVisited)
}
