package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pipes = make([][]string, 0)
var startingLoc []int

/*
*
123
4X5
678

	ENTRY      LOCATION  EXIT

| swaps between 2 and 7 or (1,0) from (1,1) and (1,2) AKA adds/removes 0,1
- swaps between 4 and 5 or (0,1) from (1,1) and (2,1) AKA adds/removes 1,0
F swaps between 5 and 7 or (2,1) from (1,1) and (1,2) AKA adds/removes 0,1
L swaps between 2 and 5 or (1,0) from (1,1) and (2,1) AKA adds/removes 1,0
J swaps between 2 and 4 or (1,0) from (1,1) and (0,1) AKA adds/removes -1,0
7 swaps between 4 and 7 or (0,1) from (1,1) and (1,2) AKA adds/removes 0,1
*/
// type Pipe struct {
// 	Location []int
// }

func VerticalPipe(location, EntryLocation []int) ([]int, []int) {
	if EntryLocation[0] < location[0] {
		return []int{location[0] + 1, location[1]}, location
	} else {
		return []int{location[0] - 1, location[1]}, location
	}
}
func HorizontalPipe(location, EntryLocation []int) ([]int, []int) {
	if EntryLocation[1] < location[1] {
		return []int{location[0], location[1] + 1}, location
	} else {
		return []int{location[0], location[1] - 1}, location
	}
}
func FPipe(location, EntryLocation []int) ([]int, []int) {
	if EntryLocation[0] > location[0] {
		return []int{location[0], location[1] + 1}, location
	} else {
		return []int{location[0] + 1, location[1]}, location
	}
}
func LPipe(location, EntryLocation []int) ([]int, []int) {
	if EntryLocation[0] < location[0] {
		return []int{location[0], location[1] + 1}, location
	} else {
		return []int{location[0] - 1, location[1]}, location
	}
}
func JPipe(location, EntryLocation []int) ([]int, []int) {
	if EntryLocation[0] < location[0] {
		return []int{location[0], location[1] - 1}, location
	} else {
		return []int{location[0] - 1, location[1]}, location
	}
}
func SevenPipe(location, EntryLocation []int) ([]int, []int) {
	if EntryLocation[0] > location[0] {
		return []int{location[0], location[1] - 1}, location
	} else {
		return []int{location[0] + 1, location[1]}, location
	}
}

func main() {
	parseInput()

	loc1, loc2 := findStartingDirections(startingLoc)
	oloc1 := startingLoc
	oloc2 := startingLoc
	counter := 1
	for {
		fmt.Printf("1: One turn %v, location: %v, pipe:%s\n", counter, loc1, pipes[loc1[0]][loc1[1]])
		fmt.Printf("2: One turn %v, location: %v, pipe:%s\n", counter, loc2, pipes[loc2[0]][loc2[1]])
		// loc1
		pipeT1 := pipes[loc1[0]][loc1[1]]
		pipeT2 := pipes[loc2[0]][loc2[1]]
		switch pipeT1 {
		case "|":
			loc1, oloc1 = VerticalPipe(loc1, oloc1)
		case "-":
			loc1, oloc1 = HorizontalPipe(loc1, oloc1)
		case "F":
			loc1, oloc1 = FPipe(loc1, oloc1)
		case "J":
			loc1, oloc1 = JPipe(loc1, oloc1)
		case "7":
			loc1, oloc1 = SevenPipe(loc1, oloc1)
		case "L":
			loc1, oloc1 = LPipe(loc1, oloc1)
		default:
			panic("unknonw direction")
		}

		switch pipeT2 {
		case "|":
			loc2, oloc2 = VerticalPipe(loc2, oloc2)
		case "-":
			loc2, oloc2 = HorizontalPipe(loc2, oloc2)
		case "F":
			loc2, oloc2 = FPipe(loc2, oloc2)
		case "J":
			loc2, oloc2 = JPipe(loc2, oloc2)
		case "7":
			loc2, oloc2 = SevenPipe(loc2, oloc2)
		case "L":
			loc2, oloc2 = LPipe(loc2, oloc2)
		default:
			panic("unknonw direction")
		}

		counter++
		if testEq(loc1, loc2) {
			fmt.Println(counter)
			break
		}
	}

}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// if Im checking up, has to be | or F or 7

func isUpOption(startingLoc []int) bool {
	if startingLoc[0] == 0 {
		return false
	}
	switch pipes[startingLoc[0]-1][startingLoc[1]] {
	case "7":
		return true
	case "F":
		return true
	case "|":
		return true
	default:
		return false
	}
}

func isDownOption(startingLoc []int) bool {
	if startingLoc[0] == len(pipes)-1 {
		return false
	}
	switch pipes[startingLoc[0]+1][startingLoc[1]] {
	case "L":
		return true
	case "J":
		return true
	case "|":
		return true
	default:
		return false
	}
}

func isLeftOption(startingLoc []int) bool {
	if startingLoc[1] == 0 {
		return false
	}
	switch pipes[startingLoc[0]][startingLoc[1]-1] {
	case "L":
		return true
	case "F":
		return true
	case "-":
		return true
	default:
		return false
	}
}

func isRightOption(startingLoc []int) bool {
	if startingLoc[1] == len(pipes[0]) {
		return false
	}
	switch pipes[startingLoc[0]][startingLoc[1]+1] {
	case "J":
		return true
	case "7":
		return true
	case "-":
		return true
	default:
		return false
	}
}

func findStartingDirections(startingLoc []int) (posve, negve []int) {
	options := make([][]int, 0)
	if isUpOption(startingLoc) {
		options = append(options, []int{startingLoc[0] - 1, startingLoc[1]})
	}
	if isDownOption(startingLoc) {
		options = append(options, []int{startingLoc[0] + 1, startingLoc[1]})
	}
	if isLeftOption(startingLoc) {
		options = append(options, []int{startingLoc[0], startingLoc[1] - 1})
	}
	if isRightOption(startingLoc) {
		options = append(options, []int{startingLoc[0], startingLoc[1] + 1})
	}
	return options[0], options[1]
}

func parseInput() {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		ps := make([]string, 0)
		for i, l := range line {
			ps = append(ps, string(l))
			if string(l) == "S" {
				startingLoc = []int{lineCount, i}
			}
		}
		pipes = append(pipes, ps)
		lineCount++
	}
	for i, p := range pipes {
		fmt.Printf("Y%v:%v\n", i, p)
	}
	fmt.Println(pipes[startingLoc[0]][startingLoc[1]])
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
