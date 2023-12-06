package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var reg, _ = regexp.Compile("([0-9])|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)")

func main() {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if reg.MatchString(line) {
			firstNum := findValue(findL2R(line))
			lastNum := findValue(findR2L(line))

			value, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstNum, lastNum))
			total = total + value
		}
	}
	fmt.Println(total)
}

func findL2R(s string) string {
	if reg.MatchString(s) {
		return string(reg.FindString(s))
	} else {
		panic("")
	}
}

func findR2L(s string) string {
	counter := 1
	test := s[len(s)-counter:]

	for !reg.MatchString(test) {
		counter++
		test = s[len(s)-counter:]
	}
	return string(reg.FindString(test))
}

func findValue(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return s
	}
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
