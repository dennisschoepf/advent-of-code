package main

import (
	"testing"
)

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

func TestMultiLineCalibration(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	want := 281
	result := GetCalibrationSumFromMultiLineString(input, numMappings)

	if want != result {
		t.Fatalf(`Expected %v, got %#v`, want, result)
	}
}

func TestGetCalibrationValueNumberString(t *testing.T) {
	input := "three9nljrrxqtxhnpsrlvbtjdznhjgmhkncrpmcpx"
	want := 39
	result := GetCalibrationValue(input, numMappings)

	if want != result {
		t.Fatalf(`Expected %v, got %#v`, want, result)
	}
}

func TestGetCalibrationValueTwoNumbers(t *testing.T) {
	input := "1abc2"
	want := 12
	result := GetCalibrationValue(input, numMappings)

	if want != result {
		t.Fatalf(`Expected %v, got %#v`, want, result)
	}
}

func TestGetCalibrationValueOneNumber(t *testing.T) {
	input := "treb7uchet"
	want := 77
	result := GetCalibrationValue(input, numMappings)

	if want != result {
		t.Fatalf(`Expected %v, got %#v`, want, result)
	}
}

func TestGetCalibrationValueThreeNumbers(t *testing.T) {
	input := "treb7uchet32"
	want := 72
	result := GetCalibrationValue(input, numMappings)

	if want != result {
		t.Fatalf(`Expected %v, got %#v`, want, result)
	}
}
