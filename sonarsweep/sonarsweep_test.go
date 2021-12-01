package sonarsweep

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

var input = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func readInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		input = append(input, i)
	}

	return input
}

func TestPartOneSample(t *testing.T) {
	actual := SonarSweep(input)

	expected := 7
	if actual != expected {
		t.Errorf("SonarSweep was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPartOneInput(t *testing.T) {
	input := readInput("./input.txt")

	actual := SonarSweep(input)

	expected := 1655
	if actual != expected {
		t.Errorf("SonarSweep was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPartTwoSample(t *testing.T) {
	slidingWindow := SlidingWindow(input)

	actual := SonarSweep(slidingWindow)

	expected := 5
	if actual != expected {
		t.Errorf("SlidingWindow was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPartTwoInput(t *testing.T) {
	input := readInput("./input.txt")
	slidingWindow := SlidingWindow(input)

	actual := SonarSweep(slidingWindow)

	expected := 1683
	if actual != expected {
		t.Errorf("SlidingWindow was incorrect, got: %d, want: %d.", actual, expected)
	}
}
