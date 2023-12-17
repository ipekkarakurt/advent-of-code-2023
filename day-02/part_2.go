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
	sum := 0

	// 12 red cubes, 13 green cubes, and 14 blue cubes
	for sc.Scan() {
		line := sc.Text()

		s := regexp.MustCompile(`Game (\d*): (.*)`).FindStringSubmatch(line)
		gamePlay := s[2]

		rounds := strings.Split(gamePlay, "; ")

		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, round := range rounds {
			picks := strings.Split(round, ", ")
			for _, pick := range picks {
				splitPick := strings.Split(pick, " ")
				count, _ := strconv.Atoi(splitPick[0])

				if splitPick[1] == "red" && count > minRed {
					minRed = count
				} else if splitPick[1] == "blue" && count > minBlue {
					minBlue = count
				} else if splitPick[1] == "green" && count > minGreen {
					minGreen = count
				}
			}
		}
		power := minRed * minGreen * minBlue
		sum += power
	}
	fmt.Println(sum)
}
