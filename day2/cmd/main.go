package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	Red      int
	Green    int
	Blue     int
	Possible bool
}

func NewRound(roundData string) Round {
	round := Round{Possible: true}
	for _, d := range strings.Split(roundData, ",") {
		data := strings.Split(strings.TrimSpace(d), " ")
		value, _ := strconv.Atoi(data[0])
		switch strings.TrimSpace(data[1]) {
		case "red":
			if value > MaxRedCubes {
				round.Possible = false
			}
			round.Red = value
		case "blue":
			if value > MaxBlueCubes {
				round.Possible = false
			}
			round.Blue = value
		case "green":
			if value > MaxGreenCubes {
				round.Possible = false
			}
			round.Green = value
		}
	}
	return round
}

var MaxRedCubes = 12
var MaxGreenCubes = 13
var MaxBlueCubes = 14

func main() {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	total := 0
	totalPower := 0
	for scanner.Scan() {
		line := scanner.Text()

		lineData := strings.Split(line, ":")
		id := strings.Split(lineData[0], " ")[1]
		gameId, _ := strconv.Atoi(id)

		rounds := extractRounds(lineData[1])
		gamePossible := true

		for _, round := range rounds {
			if !round.Possible {
				gamePossible = false
			}
		}

		if gamePossible {
			total = total + gameId
		}

		totalPower = totalPower + calculatePower(rounds)
	}
	fmt.Println(total)
	fmt.Println(totalPower)
}

func calculatePower(rounds []Round) int {
	redH := rounds[0].Red
	blueH := rounds[0].Blue
	greenH := rounds[0].Green

	for _, r := range rounds {
		if redH < r.Red {
			redH = r.Red
		}
		if blueH < r.Blue {
			blueH = r.Blue
		}
		if greenH < r.Green {
			greenH = r.Green
		}
	}

	return redH * blueH * greenH
}

func extractRounds(gameData string) []Round {
	var rounds = make([]Round, 0)
	roundsData := strings.TrimSpace(gameData)
	for _, r := range strings.Split(roundsData, ";") {
		rounds = append(rounds, NewRound(r))
	}
	return rounds
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
