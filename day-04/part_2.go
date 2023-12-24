package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	//Read input file
	input, _ := os.Open("day-04/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	cardCounts := make([]int, 200)
	cardIndex := 0

	for sc.Scan() {
		// draw original card
		cardCounts[cardIndex] += 1

		line := sc.Text()
		s := regexp.MustCompile(`Card\s+(\d*):\s+([\d+\s+]*)\s+\|\s+([\d+[\s+]*)`).FindStringSubmatch(line)
		winningNumbers := strings.Fields(s[2])
		myNumbers := strings.Fields(s[3])
		intersect := intersection(winningNumbers, myNumbers)
		if len(intersect) > 0 {
			index := len(intersect)
			for index > 0 {
				cardCounts[cardIndex+index] += cardCounts[cardIndex]
				index--
			}
		}
		cardIndex++
	}

	sum := 0
	for _, num := range cardCounts {
		sum += num
	}
	fmt.Println(sum)
}

func intersection(first, second []string) []string {
	out := []string{}
	bucket := map[string]bool{}

	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}

	return out
}
