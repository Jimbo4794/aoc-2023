package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	Time           int
	RecordDistance int
}

var reg, _ = regexp.Compile("[0-9]+")

func main() {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	tsl := scanner.Text()
	scanner.Scan()
	rsl := scanner.Text()

	// Part A
	times := make([]int, 0)

	if reg.MatchString(tsl) {
		timeString := reg.FindAllString(tsl, -1)
		for _, t := range timeString {
			v, _ := strconv.Atoi(t)
			times = append(times, v)
		}
	}

	records := make([]int, 0)

	if reg.MatchString(rsl) {
		recordString := reg.FindAllString(rsl, -1)
		for _, t := range recordString {
			v, _ := strconv.Atoi(t)
			records = append(records, v)
		}
	}

	total := 1
	for i := 0; i < len(times); i++ {
		total = total * findMargin(times[i], records[i])
	}
	fmt.Println(total)

	// Part B
	inTime := 0
	if reg.MatchString(tsl) {
		timeString := reg.FindAllString(tsl, -1)
		ts := ""
		for _, t := range timeString {
			ts = ts + t
		}
		inTime, _ = strconv.Atoi(ts)
	}

	inRecord := 0
	if reg.MatchString(rsl) {
		recordString := reg.FindAllString(rsl, -1)
		rs := ""
		for _, t := range recordString {
			rs = rs + t
		}
		inRecord, _ = strconv.Atoi(rs)
	}

	fmt.Println(findMargin(inTime, inRecord))
}

func findMargin(time, record int) int {
	count := 0
	for i := 0; i <= time; i++ {
		if (1*i)*(time-i) > record {
			count++
		}
	}
	return count
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
