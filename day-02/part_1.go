package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("day-02/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	sum := 0

	// 12 red cubes, 13 green cubes, and 14 blue cubes
	for sc.Scan() {
		line := sc.Text()
		possibleGame := true

		s := regexp.MustCompile(`Game (\d*): (.*)`).FindStringSubmatch(line)
		gameCount := s[1]
		gamePlay := s[2]

		rounds := strings.Split(gamePlay, "; ")

		for _, round := range rounds {
			if possibleGame {
				picks := strings.Split(round, ", ")
				for _, pick := range picks {
					splitPick := strings.Split(pick, " ")
					count, _ := strconv.Atoi(splitPick[0])

					if splitPick[1] == "red" && count > maxRed {
						possibleGame = false
						break
					} else if splitPick[1] == "blue" && count > maxBlue {
						possibleGame = false
						break
					} else if splitPick[1] == "green" && count > maxGreen {
						possibleGame = false
						break
					}
				}
			}
		}

		if possibleGame {
			c, _ := strconv.Atoi(gameCount)
			sum += c
		}
	}
	fmt.Println(sum)
}
