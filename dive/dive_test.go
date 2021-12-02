package dive

import (
	"log"
	"os"
	"strings"
	"testing"
)

func readInput(path string) string {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(contents))
}

var input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestDiveSample(t *testing.T) {
	actual := Dive(input)

	expected := 150
	if actual != expected {
		t.Errorf("Dive was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestDiveInput(t *testing.T) {
	input := readInput("input.txt")
	actual := Dive(input)

	expected := 1648020
	if actual != expected {
		t.Errorf("Dive was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestDiveWithAimSample(t *testing.T) {
	actual := DiveWithAim(input)

	expected := 900
	if actual != expected {
		t.Errorf("DiveWithAim was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestDiveWithAimInput(t *testing.T) {
	input := readInput("input.txt")
	actual := DiveWithAim(input)

	expected := 1759818555
	if actual != expected {
		t.Errorf("DiveWithAim was incorrect, got: %d, want: %d.", actual, expected)
	}
}
