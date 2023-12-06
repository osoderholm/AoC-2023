package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(realData))                              // 3317888
	fmt.Printf("Part 2: %d\n", part1(strings.ReplaceAll(realData, " ", ""))) // 24655068
}

func part1(data string) int {
	var prod int

	races := getRaces(data)

	for _, race := range races {
		var wins int
		for hold := 1; hold < race[0]; hold++ {
			if (race[0]-hold)*hold > race[1] {
				wins++
			}
		}
		if wins == 0 {
			continue
		}
		if prod == 0 {
			prod = 1
		}
		prod *= wins
	}

	return prod
}

func getRaces(data string) [][2]int {
	races := make([][2]int, 0)

	re := regexp.MustCompile(`(\d+)`)

	for i, line := range strings.Split(data, "\n") {
		if i > 1 {
			break
		}
		nums := re.FindAllString(line, -1)
		for c, num := range nums {
			val, _ := strconv.Atoi(num)
			if i == 0 {
				races = append(races, [2]int{val, 0})
				continue
			}
			races[c][i] = val
		}
	}

	return races
}

const testData = `Time:      7  15   30
Distance:  9  40  200`

const realData = `Time:        42     89     91     89
Distance:   308   1170   1291   1467`
