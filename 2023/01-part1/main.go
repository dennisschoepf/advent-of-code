package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	sum := 0

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sum += GetCalibrationValue(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", sum)
}

func GetCalibrationValue(input string) (value int) {
	allDigits := []int{}

	for _, char := range input {
		if unicode.IsNumber(char) {
			allDigits = append(allDigits, int(char-'0'))
		}
	}

	length := len(allDigits)
	first := allDigits[0]
	last := allDigits[len(allDigits)-1]

	if length < 2 {
		return first*10 + first
	}

	return first*10 + last
}
