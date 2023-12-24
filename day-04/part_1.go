package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	//Read input file
	input, _ := os.Open("day-04/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	sum := 0

	for sc.Scan() {
		line := sc.Text()
		s := regexp.MustCompile(`Card\s+\d*:\s+([\d+\s+]*)\s+\|\s+([\d+[\s+]*)`).FindStringSubmatch(line)
		winningNumbers := strings.Fields(s[1])
		myNumbers := strings.Fields(s[2])
		intersect := intersection(winningNumbers, myNumbers)
		if len(intersect) > 0 {
			sum += int(math.Pow(2, float64(len(intersect)-1)))
		}
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
