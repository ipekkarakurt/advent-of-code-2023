package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	//Read input file
	input, _ := os.Open("day-1/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	result := 0

	for sc.Scan() {
		line := sc.Text()

		calibrationVal := 0
		firstDigit := 0
		lastDigit := 0
		foundDigitCount := 0

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
