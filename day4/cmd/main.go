package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var scratchers = make(map[int]*ScratchCard, 0)

func main() {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := NewScratchCard(line)
		scratchers[s.Id] = &s
	}

	for index := 1; index <= len(scratchers); index++ {
		k := scratchers[index]
		total = total + k.Instances
		if k.Wins() == 0 {
			continue
		}
		for i := 0; i < k.Instances; i++ {
			for ii := 1; ii <= k.Wins(); ii++ {
				scratchers[index+ii].Instances++
			}
		}
	}
	fmt.Println(total)
}

func (s *ScratchCard) Wins() int {
	matches := 0
	for _, wn := range s.WinningNumbers {
		for _, n := range s.Numbers {
			if wn == n {
				matches++
			}
		}
	}
	return matches
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

type ScratchCard struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
	Instances      int
}

func NewScratchCard(in string) ScratchCard {
	split1 := strings.Split(in, ":")
	idString := strings.TrimSpace(strings.ReplaceAll(split1[0], "Card ", ""))
	winningNumbers := make([]int, 0)
	numbers := make([]int, 0)
	togggle := true
	for _, v := range strings.Split(strings.TrimSpace(split1[1]), " ") {
		vInt, _ := strconv.Atoi(v)
		if v == "|" {
			togggle = false
			continue
		}
		if v == "" {
			continue
		}

		if togggle {
			winningNumbers = append(winningNumbers, vInt)
		} else {
			numbers = append(numbers, vInt)
		}
	}

	id, _ := strconv.Atoi(idString)

	return ScratchCard{
		Id:             id,
		WinningNumbers: winningNumbers,
		Numbers:        numbers,
		Instances:      1,
	}
}
