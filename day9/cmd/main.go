package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	histories := parseInput()
	total := 0
	for _, h := range histories {
		levels := createLevels(h)
		total = total + findNextValue(levels)
	}
	fmt.Printf("Total for Part A: %v\n", total)

	totalB := 0
	for _, h := range histories {
		levels := createLevels(h)
		for _, l := range levels {
			fmt.Println(l)
		}
		i := findPrevValue(levels)
		fmt.Printf("first value: %v\n", i)
		fmt.Print("\n\n")
		totalB = totalB + i
	}
	fmt.Printf("Total for Part B: %v\n", totalB)
}

func findNextValue(levels [][]int) int {
	for i := len(levels) - 1; i >= 0; i-- {
		if i == len(levels)-1 {
			levels[i] = append(levels[i], 0)
		} else {
			size := len(levels[i]) - 1
			levels[i] = append(levels[i], levels[i][size]+levels[i+1][size])
		}

	}
	return levels[0][len(levels[0])-1]
}

func prepend(is []int, i int) []int {
	return append([]int{i}, is...)
}

func findPrevValue(levels [][]int) int {
	for i := len(levels) - 1; i >= 0; i-- {
		if i == len(levels)-1 {
			levels[i] = prepend(levels[i], 0)
		} else {
			// size := len(levels[i]) - 1
			levels[i] = prepend(levels[i], levels[i][0]-levels[i+1][0])
		}
		fmt.Printf("Level %v: %v\n", i, levels[i])
	}
	fmt.Print("\n\n")
	return levels[0][0]
}

func createLevels(vals []int) [][]int {
	levels := make([][]int, 0)
	levels = append(levels, vals)
	i := 0
	// fmt.Printf("Level %v: %v\n", i, vals)
	for {
		i++
		vals = findDifferences(vals)
		levels = append(levels, vals)
		// fmt.Printf("Level %v: %v\n", i, vals)

		finished := true
		for _, v := range vals {
			if v != 0 {
				finished = false
			}
		}

		if finished {
			// fmt.Println("Breaking")
			break
		}

	}

	return levels
}

func findDifferences(vals []int) []int {
	difs := make([]int, 0)
	for i := 0; i < len(vals)-1; i++ {
		difs = append(difs, vals[i+1]-vals[i])
	}
	return difs
}

func parseInput() [][]int {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)
	result := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		history := make([]int, 0)
		for _, value := range strings.Split(line, " ") {
			v, _ := strconv.Atoi(value)
			history = append(history, v)
		}
		result = append(result, history)
	}
	return result
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
