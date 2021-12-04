package binarydiagnostic

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

var sample = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func convert(input []string) ([]int, int) {
	output := make([]int, len(input))

	for i, s := range input {
		n, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		output[i] = int(n)
	}

	return output, len(input[0])
}

func readInput(path string) ([]int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []int
	var width int

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		s := scanner.Text()
		width = len(s)
		i, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		input = append(input, int(i))
	}

	return input, width
}

func TestGammaRateSample(t *testing.T) {
	input, width := convert(sample)
	actual := GammaRate(input, width)

	expected := 22
	if actual != expected {
		t.Errorf("GammaRate was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestEpsilonRateSample(t *testing.T) {
	input, width := convert(sample)
	gamma := GammaRate(input, width)
	actual := EpsilonRateFromGamma(gamma, width)

	expected := 9
	if actual != expected {
		t.Errorf("EpsilonRate was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPowerConsumptionSample(t *testing.T) {
	input, width := convert(sample)
	actual := PowerConsumption(input, width)

	expected := 198
	if actual != expected {
		t.Errorf("EpsilonRate was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPowerConsumptionInput(t *testing.T) {
	input, width := readInput("./input.txt")
	actual := PowerConsumption(input, width)

	expected := 3969000
	if actual != expected {
		t.Errorf("PowerConsumption was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestOxygenGeneratorRatingSample(t *testing.T) {
	input, width := convert(sample)
	actual := OxygenGeneratorRating(input, width)

	expected := 23
	if actual != expected {
		t.Errorf("OxygenGeneratorRating was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestC02ScrubberRatingSample(t *testing.T) {
	input, width := convert(sample)
	actual := C02ScrubberRating(input, width)

	expected := 10
	if actual != expected {
		t.Errorf("C02ScrubberRating was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestLifeSupportRatingSample(t *testing.T) {
	input, width := convert(sample)
	actual := LifeSupportRating(input, width)

	expected := 230
	if actual != expected {
		t.Errorf("LifeSupportRating was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestLifeSupportRatingInput(t *testing.T) {
	input, width := readInput("./input.txt")
	actual := LifeSupportRating(input, width)

	expected := 4267809
	if actual != expected {
		t.Errorf("LifeSupportRating was incorrect, got: %d, want: %d.", actual, expected)
	}
}
