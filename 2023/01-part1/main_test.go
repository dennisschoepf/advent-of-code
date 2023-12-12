package main

import (
	"testing"
)

func TestGetCalibrationValueTwoNumbers(t *testing.T) {
	input := "1abc2"
	want := 12
	result := GetCalibrationValue(input)

	if want != result {
		t.Fatalf(`Expected %q, got %#q`, input, want)
	}
}

func TestGetCalibrationValueOneNumber(t *testing.T) {
	input := "treb7uchet"
	want := 77
	result := GetCalibrationValue(input)

	if want != result {
		t.Fatalf(`Expected %q, got %#q`, input, want)
	}
}

func TestGetCalibrationValueThreeNumbers(t *testing.T) {
	input := "treb7uchet32"
	want := 72
	result := GetCalibrationValue(input)

	if want != result {
		t.Fatalf(`Expected %q, got %#q`, input, want)
	}
}
