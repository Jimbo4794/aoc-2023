package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Rule struct {
	SourceStart      int
	DestinationStart int
	Range            int
}

var rules = [][]Rule{
	make([]Rule, 0),
	make([]Rule, 0),
	make([]Rule, 0),
	make([]Rule, 0),
	make([]Rule, 0),
	make([]Rule, 0),
	make([]Rule, 0),
}

var locations = make([]int, 0)

func (r *Rule) Effects(i int) bool {
	if i < r.SourceStart {
		return false
	}
	return (r.SourceStart + r.Range - 1) >= i
}

func (r *Rule) EvaluateRule(i int) int {
	return i - r.SourceStart + r.DestinationStart
}

func SeedLocation(i int) int {
	for _, r := range rules {
		for _, sr := range r {
			if sr.Effects(i) {
				i = sr.EvaluateRule(i)
				break
			}
		}
	}
	return i
}

func main() {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	//Take Seeds from first line
	scanner.Scan()
	seedsString := strings.Split(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]), " ")
	seedsRanges := make([]int, 0)
	for _, v := range seedsString {
		vInt, _ := strconv.Atoi(v)
		seedsRanges = append(seedsRanges, vInt)
	}

	rule := 0
	for scanner.Scan() {
		line := scanner.Text()
		r := make([]Rule, 0)
		if strings.Contains(line, "map:") {
			for scanner.Scan() {
				line = strings.TrimSpace(scanner.Text())
				if line == "" {
					break
				}
				numString := strings.Split(line, " ")
				numInt := make([]int, 0)
				for _, v := range numString {
					vInt, _ := strconv.Atoi(v)
					numInt = append(numInt, vInt)
				}
				r = append(r, Rule{
					DestinationStart: numInt[0],
					SourceStart:      numInt[1],
					Range:            numInt[2],
				})
			}

		} else {
			continue
		}
		rules[rule] = r
		rule++
	}
	var wg sync.WaitGroup
	for i := 0; i < len(seedsRanges); i = i + 2 {
		wg.Add(1)
		fmt.Printf("Starting thread looking at starting number of %v over %v range\n", seedsRanges[i], seedsRanges[i+1])
		go lowestInRange(seedsRanges[i], seedsRanges[i+1], &wg)
	}
	wg.Wait()
	lowestAll := slices.Min(locations)
	fmt.Printf("Lowest found location over all seeds: %v", lowestAll)
}

func lowestInRange(start, seedRange int, wg *sync.WaitGroup) {
	lowest := SeedLocation(start)
	for i := start; i < start+seedRange; i++ {
		ii := SeedLocation(i)
		if ii < lowest {
			lowest = ii
		}
	}
	locations = append(locations, lowest)
	fmt.Printf("Finishing thread looking at starting number of %v over %v range. Lowest found: %v\n", start, seedRange, lowest)
	defer wg.Done()
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
