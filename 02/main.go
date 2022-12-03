package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("02.txt")
	fmt.Println("Part one total:", part1(file))
	file.Close()

	file, _ = os.Open("02.txt")
	fmt.Println("Part two total:", part2(file))
	file.Close()
}

func part1(file *os.File) int {
	total := 0

	scanner := bufio.NewScanner(file)
	// Only nine possibilities, so why not just use a switch statement?
	for scanner.Scan() {
		switch scanner.Text() {
		// opponent plays rock
		case "A X":
			// we play rock (1) = draw (3)
			total += 4
		case "A Y":
			// we play paper (2) = win (6)
			total += 8
		case "A Z":
			// we play scissors (3) = loss (0)
			total += 3

		// opponent plays paper
		case "B X":
			// we play rock (1) = loss (0)
			total += 1
		case "B Y":
			// we play paper (2) = draw (3)
			total += 5
		case "B Z":
			// we play scissors (3) = win (6)
			total += 9

		// opponent plays scissors
		case "C X":
			// we play rock (1) = win (6)
			total += 7
		case "C Y":
			// we play paper (2) = loss (0)
			total += 2
		case "C Z":
			// we play scissors (3) = draw (3)
			total += 6

		default:
			panic("unknown input")
		}
	}

	return total
}

func part2(file *os.File) int {
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Only nine possibilities, so why not just use a switch statement?
		switch scanner.Text() {
		// opponent plays rock
		case "A X":
			// we want to lose so we play scissors (3) = loss (0)
			total += 3
		case "A Y":
			// we want to draw so we play rock (1) = draw (3)
			total += 4
		case "A Z":
			// we want to win so we play paper (2) = win (6)
			total += 8

		// opponent plays paper
		case "B X":
			// we want to lose so we play rock (1) = loss (0)
			total += 1
		case "B Y":
			// we want to draw so we play paper (2) = draw (3)
			total += 5
		case "B Z":
			// we want to win so we play scissors (3) = win (6)
			total += 9

		// opponent plays scissors
		case "C X":
			// we want to lose so we play paper (2) = loss (0)
			total += 2
		case "C Y":
			// we want to draw so we play scissors (3) = draw (3)
			total += 6
		case "C Z":
			// we want to win so we play rock (1) = win (6)
			total += 7

		default:
			panic("unknown input")
		}
	}

	return total
}
