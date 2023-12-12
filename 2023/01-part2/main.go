package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type numOccurence struct {
	index  int
	number int
}

func GetCalibrationValue(input string, numMappings map[string]int) (value int) {
	allDigits := []numOccurence{}

	// Gather actual numbers
	for index, char := range input {
		if unicode.IsNumber(char) {
			occurence := numOccurence{
				index:  index,
				number: int(char - '0'),
			}
			allDigits = append(allDigits, occurence)
		}
	}

	// Gather numbers as text
	for k, v := range numMappings {
		numberAt := strings.Index(input, k)

		if numberAt != -1 {
			occurence := numOccurence{
				index:  numberAt,
				number: v,
			}
			allDigits = append(allDigits, occurence)
		}
	}

	var firstOccurence numOccurence
	var lastOccurence numOccurence

	firstOccurenceZero := firstOccurence
	lastOccurenceZero := lastOccurence

	for _, occurence := range allDigits {
		if firstOccurence == firstOccurenceZero {
			firstOccurence = occurence
		}

		if occurence.index < firstOccurence.index {
			firstOccurence = occurence
		}

		if lastOccurence == lastOccurenceZero || occurence.index > lastOccurence.index {
			lastOccurence = occurence
		}

		if occurence.index > lastOccurence.index {
			lastOccurence = occurence
		}
	}

	if lastOccurence == lastOccurenceZero {
		return firstOccurence.number*10 + firstOccurence.number
	}

	return firstOccurence.number*10 + lastOccurence.number
}

func GetCalibrationSumFromMultiLineString(input string, numMappings map[string]int) (result int) {
	sum := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		sum += GetCalibrationValue(line, numMappings)
	}

	return sum
}

func main() {
	var numMappings = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		sum += GetCalibrationValue(scanner.Text(), numMappings)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", sum)
}
