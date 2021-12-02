package dive

import (
	"log"
	"strconv"
	"strings"
)

type Direction int

const (
	Forward Direction = iota
	Down
	Up
)

type Instruction struct {
	Direction Direction
	Units     int
}

func ParseDirection(direction string) Direction {
	switch direction {
	case "forward":
		return Forward
	case "down":
		return Down
	case "up":
		return Up
	}
	log.Fatalf("unknown direction: %s\n", direction)
	panic("unknown direction")
}

func ParseUnits(units string) int {
	i, err := strconv.Atoi(units)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ParseInstruction(input string) Instruction {
	instructionAndUnits := strings.Split(input, " ")
	direction := ParseDirection(instructionAndUnits[0])
	units := ParseUnits(instructionAndUnits[1])
	return Instruction{
		Direction: direction,
		Units:     units,
	}
}

func ParseInstructions(input string) []Instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, len(lines))

	for i := 0; i < len(lines); i++ {
		instructions[i] = ParseInstruction(lines[i])
	}

	return instructions
}

func Dive(input string) int {
	instructions := ParseInstructions(input)

	var distance int
	var depth int
	for _, instruction := range instructions {
		switch instruction.Direction {
		case Up:
			depth -= instruction.Units
		case Down:
			depth += instruction.Units
		case Forward:
			distance += instruction.Units
		}
	}
	return distance * depth
}

func DiveWithAim(input string) int {
	instructions := ParseInstructions(input)

	var aim int
	var depth int
	var distance int
	for _, instruction := range instructions {
		switch instruction.Direction {
		case Up:
			aim += instruction.Units
		case Down:
			aim -= instruction.Units
		case Forward:
			distance += instruction.Units
			depth -= (instruction.Units * aim)
		}
	}
	return distance * depth
}
