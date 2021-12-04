package binarydiagnostic

import "fmt"

func GammaRate(input []int, width int) int {
	counts := make([]int, width)

	for _, number := range input {
		for i := 0; i < width; i++ {
			mask := 1 << i

			if number&mask == mask {
				counts[i] += 1
			}
		}
	}

	var output int
	for i := 0; i < width; i++ {
		c := counts[i]
		if c > len(input)/2 {
			output = output | (1 << i)
		}
	}

	return output
}

func EpsilonRateFromGamma(gamma, width int) int {
	mask := (1 << width) - 1
	return (^gamma) & mask
}

func PowerConsumption(input []int, width int) int {
	gamma := GammaRate(input, width)
	return gamma * EpsilonRateFromGamma(gamma, width)
}

type Select func([]int, []int) []int

func rating(input []int, width int, s Select) int {
	numbers := input

	for i := 0; i < width; i++ {
		if len(numbers) == 1 {
			return numbers[0]
		}

		mask := 1 << (width - 1 - i)

		var left []int
		var right []int

		for _, number := range numbers {
			if number&mask == mask {
				left = append(left, number)
			} else {
				right = append(right, number)
			}
		}

		numbers = s(left, right)
	}

	for _, n := range numbers {
		fmt.Printf("%b\n", n)
	}

	return numbers[0]
}

func OxygenGeneratorRating(input []int, width int) int {
	return rating(input, width, func(left []int, right []int) []int {
		if len(left) >= len(right) {
			return left
		}
		return right
	})
}

func C02ScrubberRating(input []int, width int) int {
	return rating(input, width, func(left []int, right []int) []int {
		if len(left) < len(right) {
			return left
		}
		return right
	})
}

func LifeSupportRating(input []int, width int) int {
	return OxygenGeneratorRating(input, width) * C02ScrubberRating(input, width)
}
