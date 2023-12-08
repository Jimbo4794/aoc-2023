package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type DesertMap struct {
	Instructions []rune
	Nodes        map[string]DesertNode
}

type DesertNode struct {
	Left  string
	Right string
}

type DesertNodeP struct {
	Left  *DesertNodeP
	Right *DesertNodeP
}

var answerMap = make(map[string]int)
var locations = make([]string, 0)

func main() {
	desertMap := parseInput()

	partAMoves := partAPathFinder(desertMap)
	fmt.Printf("Took %v steps to reach ZZZ\n", partAMoves)
	partBMoves := partBPathFinder(desertMap)
	fmt.Printf("Took %v steps to reach nodes ending in Z\n", partBMoves)
}

func partBPathFinder(desertMap DesertMap) int {
	locations = findStartingLocations(desertMap.Nodes)
	fmt.Println(locations)

	i := make([]int, 0)
	for _, loc := range locations {
		i = append(i, findPathLength(loc, desertMap))
	}

	return LCM(i)
}

func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func LCM(i []int) int {
	a := i[0]
	b := i[1]
	result := a * b / GCD(a, b)

	if len(i) == 2 {
		return result
	}

	remainderI := i[2:]
	remainderI = append(remainderI, result)

	return LCM(remainderI)
}

func findPathLength(startingLoc string, desertMap DesertMap) int {
	loc := startingLoc
	zLoc := ""
	index := 0
	counter := 0
	for {
		// fmt.Println(loc)
		if index == len(desertMap.Instructions) {
			index = 0
		}
		k := desertMap.Instructions[index]

		if k == 76 {
			loc = desertMap.Nodes[loc].Left
		} else {
			loc = desertMap.Nodes[loc].Right
		}

		if loc == zLoc {
			return counter
		}

		if loc[2] == 90 {
			zLoc = loc
			counter = 0
		}

		index++
		counter++
	}
}

func findStartingLocations(nodes map[string]DesertNode) []string {
	locations := make([]string, 0)
	for k, _ := range nodes {
		if k[2] == 65 {
			locations = append(locations, k)
		}
	}
	return locations
}

func partAPathFinder(desertMap DesertMap) int {
	location := "AAA"
	counter := 0
	index := 0
	for {
		if location == "ZZZ" {
			break
		}
		if index == len(desertMap.Instructions) {
			index = 0
		}
		k := desertMap.Instructions[index]
		// Left
		if k == 76 {
			location = desertMap.Nodes[location].Left
		} else {
			location = desertMap.Nodes[location].Right
		}
		counter++
		index++
	}
	return counter
}

func parseInput() DesertMap {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	instructions := scanner.Text()

	// Quick hack to remove first empty line
	scanner.Scan()

	reg, _ := regexp.Compile("[0-9A-Z][0-9A-Z][0-9A-Z]")
	nodes := make(map[string]DesertNode)
	for scanner.Scan() {
		line := scanner.Text()
		if reg.MatchString(line) {
			v := reg.FindAllString(line, -1)
			nodes[v[0]] = DesertNode{
				Left:  v[1],
				Right: v[2],
			}
		}

	}

	locs := make([]string, 0)
	for k, _ := range nodes {
		locs = append(locs, k)
	}

	return DesertMap{
		Instructions: []rune(instructions),
		Nodes:        nodes,
	}
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
