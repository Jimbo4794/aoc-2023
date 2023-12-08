package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
	Value int64
}

func base10HandValue(hand string) int64 {
	base13String := ""
	for _, v := range hand {
		base13String = base13String + m[string(v)]
	}
	value, _ := strconv.ParseInt(base13String, 13, 64)
	return value
}

var set5 = make([]Hand, 0)
var set4 = make([]Hand, 0)
var setFH = make([]Hand, 0)
var set3 = make([]Hand, 0)
var set2P = make([]Hand, 0)
var set1P = make([]Hand, 0)
var setH = make([]Hand, 0)

func NewHand(s string) Hand {
	v := strings.Split(s, " ")
	bid, _ := strconv.Atoi(v[1])

	return Hand{
		Cards: v[0],
		Bid:   bid,
		Value: base10HandValue(v[0]),
	}
}

func (h *Hand) DefineType(partA bool) string {
	mv := make(map[rune]int)
	for _, v := range h.Cards {
		if counter, ok := mv[v]; ok {
			mv[v] = counter + 1
		} else {
			mv[v] = 1
		}
	}

	if partA {
		return partAGameLogic(mv)
	}
	return partBGameLogic(mv)

}

func partBGameLogic(mv map[rune]int) string {
	if len(mv) == 5 {
		if _, ok := mv[74]; ok {
			return "1Pair"
		}
		return "HighC"
	}

	if len(mv) == 4 {
		if _, ok := mv[74]; ok {
			return "3ofK"
		}
		return "1Pair"
	}

	if len(mv) == 3 {
		for _, v := range mv {
			if v == 3 {
				if _, ok := mv[74]; ok {
					return "4ofK"
				}
				return "3ofK"
			}

		}
		if val, ok := mv[74]; ok {
			if val == 1 {
				return "FullH"
			} else {
				return "4ofK"
			}
		}
		return "2Pair"
	}

	if len(mv) == 2 {
		if _, ok := mv[74]; ok {
			return "5ofK"
		}
		for _, v := range mv {
			if v == 2 {
				return "FullH"
			}
		}
		return "4ofK"
	}

	if len(mv) == 1 {
		return "5ofK"
	}

	return ""
}

func partAGameLogic(mv map[rune]int) string {
	if len(mv) == 5 {
		return "HighC"
	}

	if len(mv) == 4 {
		return "1Pair"
	}

	if len(mv) == 3 {
		for _, v := range mv {
			if v == 3 {
				return "3ofK"
			}
		}
		return "2Pair"
	}

	if len(mv) == 2 {
		for _, v := range mv {
			if v == 2 {
				return "FullH"
			}
		}
		return "4ofK"
	}

	if len(mv) == 1 {
		return "5ofK"
	}

	return ""
}

var m = make(map[string]string)

func main() {

	m["A"] = "c"
	m["K"] = "b"
	m["J"] = "a"
	m["Q"] = "9"
	m["T"] = "8"
	m["9"] = "7"
	m["8"] = "6"
	m["7"] = "5"
	m["6"] = "4"
	m["5"] = "3"
	m["4"] = "2"
	m["3"] = "1"
	m["2"] = "0"
	fmt.Printf("PartA: %v\n", RunGameMaths(m, true))
	reset()

	m["A"] = "c"
	m["K"] = "b"
	m["Q"] = "a"
	m["T"] = "9"
	m["9"] = "8"
	m["8"] = "7"
	m["7"] = "6"
	m["6"] = "5"
	m["5"] = "4"
	m["4"] = "3"
	m["3"] = "2"
	m["2"] = "1"
	m["J"] = "0"
	fmt.Printf("PartB: %v\n", RunGameMaths(m, false))

}

func reset() {
	setH = make([]Hand, 0)
	set1P = make([]Hand, 0)
	set2P = make([]Hand, 0)
	set3 = make([]Hand, 0)
	setFH = make([]Hand, 0)
	set4 = make([]Hand, 0)
	set5 = make([]Hand, 0)
}

func RunGameMaths(m map[string]string, partA bool) int {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		h := NewHand(line)
		switch h.DefineType(partA) {
		case "5ofK":
			set5 = append(set5, h)
		case "4ofK":
			set4 = append(set4, h)
		case "FullH":
			setFH = append(setFH, h)
		case "3ofK":
			set3 = append(set3, h)
		case "2Pair":
			set2P = append(set2P, h)
		case "1Pair":
			set1P = append(set1P, h)
		case "HighC":
			setH = append(setH, h)
		default:
			panic("")
		}
	}

	// Sort all sets
	sort.Slice(setH, func(i, j int) bool {
		return setH[i].Value < setH[j].Value
	})
	sort.Slice(set1P, func(i, j int) bool {
		return set1P[i].Value < set1P[j].Value
	})
	sort.Slice(set2P, func(i, j int) bool {
		return set2P[i].Value < set2P[j].Value
	})
	sort.Slice(set3, func(i, j int) bool {
		return set3[i].Value < set3[j].Value
	})
	sort.Slice(setFH, func(i, j int) bool {
		return setFH[i].Value < setFH[j].Value
	})
	sort.Slice(set4, func(i, j int) bool {
		return set4[i].Value < set4[j].Value
	})
	sort.Slice(set5, func(i, j int) bool {
		return set5[i].Value < set5[j].Value
	})

	total := 0
	rank := 1
	for _, v := range setH {
		total = total + v.Bid*rank
		rank++
	}
	for _, v := range set1P {
		total = total + v.Bid*rank
		rank++
	}
	for _, v := range set2P {
		total = total + v.Bid*rank
		rank++
	}
	for _, v := range set3 {
		total = total + v.Bid*rank
		rank++
	}
	for _, v := range setFH {
		total = total + v.Bid*rank
		rank++
	}
	for _, v := range set4 {
		total = total + v.Bid*rank
		rank++
	}
	for _, v := range set5 {
		total = total + v.Bid*rank
		rank++
	}
	return total
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
