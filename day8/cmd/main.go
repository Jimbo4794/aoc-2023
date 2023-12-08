package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
)

type DesertMap struct {
	Instructions []rune
	Nodes        map[string]DesertNode
}

type DesertNode struct {
	Left  string
	Right string
}

var answerMap = make(map[string]int)
var locations = make([]string, 0)

func main() {
	desertMap := parseInput()

	// partAMoves := partAPathFinder(desertMap)
	// fmt.Printf("Took %v steps to reach ZZZ\n", partAMoves)
	partBMoves := partBPathFinder(desertMap)
	fmt.Printf("Took %v steps to reach nodes ending in Z\n", partBMoves)
}

func partBPathFinder(desertMap DesertMap) int {
	locations = findStartingLocations(desertMap.Nodes)
	fmt.Println(locations)

	index := 0
	counter := 0
	for !checkAtFinalLocations(locations, counter) {
		if index == len(desertMap.Instructions) {
			index = 0
		}
		instruction := desertMap.Instructions[index]
		var wg sync.WaitGroup
		for k, v := range locations {
			wg.Add(1)
			go setNextLocation(v, desertMap.Nodes, instruction, k, &wg)
		}
		wg.Wait()
		counter++
		index++
	}

	return counter
}

func setNextLocation(currentLoc string, nodes map[string]DesertNode, instruction rune, index int, wg *sync.WaitGroup) {
	if instruction == 76 {
		locations[index] = nodes[currentLoc].Left
	} else {
		locations[index] = nodes[currentLoc].Right
	}
	wg.Done()
}

func checkAtFinalLocations(locations []string, counter int) bool {
	for i, v := range locations {
		if v[2] != 90 {
			return false
		}
		if i > 0 {
			fmt.Printf("found %v Z - counter: %v\n", i, counter)
		}
	}
	return true
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
