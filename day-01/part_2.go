package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//Read input file
	input, _ := os.Open("day-01/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	result := 0

	for sc.Scan() {
		line := sc.Text()

		calibrationVal := 0
		firstDigit := 0
		lastDigit := 0
		foundDigitCount := 0

		// replacement logic is a little horrendous,
		// because overlapping strings should count as separate digits
		// ex: oneight -> 18
		replacements := map[string]string{
			"one":   "o1e",
			"two":   "t2o",
			"three": "t3e",
			"four":  "f4r",
			"five":  "f5e",
			"six":   "s6x",
			"seven": "s7n",
			"eight": "e8t",
			"nine":  "n9e",
		}

		for sFrom, sTo := range replacements {
			line = strings.Replace(line, sFrom, sTo, -1)
		}

		for _, c := range line {
			if unicode.IsDigit(c) && foundDigitCount == 0 {
				// found first digit
				firstDigit = int(c - 48)
				foundDigitCount++
			}
			if unicode.IsDigit(c) && foundDigitCount >= 1 {
				// found last digit
				lastDigit = int(c - 48)
				foundDigitCount++
			}
		}

		if foundDigitCount == 1 {
			lastDigit = firstDigit
		}

		calibrationVal += firstDigit*10 + lastDigit
		result += calibrationVal
	}

	fmt.Println(result)
}
