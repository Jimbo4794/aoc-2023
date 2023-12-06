package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type gear struct {
	Ratios []int
}

type number struct {
	Value    int
	Location location
	Length   int
	GearConf gearConfig
}

type gearConfig struct {
	NextToGear   bool
	GearLocation location
}

type location struct {
	X int
	Y int
}

var symbols = "*%$=/+&-#@"
var reg, _ = regexp.Compile("[0-9]+")

var m = make([][]byte, 0)
var numberLocations = make([]number, 0)
var gears = make(map[location]gear)

func main() {
	input := readInputFile()
	defer input.Close()
	scanner := bufio.NewScanner(input)

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, []byte(line))
		checkLine(line, lineCount)
		lineCount++
	}

	total := 0
	for _, number := range numberLocations {
		if CheckComponents(number) {
			total = total + number.Value
		}
	}
	fmt.Println(total)

	gearTotal := 0
	for _, gear := range gears {
		if len(gear.Ratios) == 2 {
			gearTotal = gearTotal + (gear.Ratios[0] * gear.Ratios[1])
		}
	}
	fmt.Println(gearTotal)
}

func checkLine(line string, lineCount int) {
	if reg.MatchString(line) {
		matches := reg.FindAllString(line, -1)
		matchesLoc := reg.FindAllStringIndex(line, -1)
		for i, val := range matchesLoc {
			v, _ := strconv.Atoi(matches[i])
			numberLocations = append(numberLocations, number{
				Value: v,
				Location: location{
					X: val[0],
					Y: lineCount,
				},
				Length: val[1] - val[0],
				GearConf: gearConfig{
					NextToGear:   false,
					GearLocation: location{},
				},
			})
		}
	}
}

func CheckComponents(num number) bool {
	searchX := num.Location.X - 1
	if num.Location.X == 0 {
		searchX = 0
	}

	searchMaxX := num.Location.X + num.Length
	if searchMaxX > len(m[0])-1 {
		searchMaxX = len(m[0]) - 1
	}

	searchY := num.Location.Y - 1
	if num.Location.Y == 0 {
		searchY = 0
	}

	searchMaxY := num.Location.Y + 1
	if searchMaxY >= len(m) {
		searchMaxY = num.Location.Y
	}

	isPart := false
	for i := searchY; i <= searchMaxY; i++ {
		for ii := searchX; ii <= searchMaxX; ii++ {
			if strings.Contains(symbols, string(m[i][ii])) {
				if string(m[i][ii]) == "*" {
					num.GearConf = gearConfig{
						NextToGear: true,
						GearLocation: location{
							X: ii,
							Y: i,
						},
					}
				}

				if g, ok := gears[num.GearConf.GearLocation]; ok {
					g.Ratios = append(g.Ratios, num.Value)
					gears[num.GearConf.GearLocation] = g
				} else {
					gears[num.GearConf.GearLocation] = gear{
						Ratios: []int{num.Value},
					}
				}
				isPart = true
			}
		}
	}

	return isPart
}

func readInputFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
