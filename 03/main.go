package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("03.txt")
	fmt.Println("Part one total:", part1(file))
	file.Close()

	file, _ = os.Open("03.txt")
	fmt.Println("Part two total:", part2(file))
	file.Close()
}

func part1(file *os.File) int {
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		item := getDuplicate(scanner.Text())
		total += getNumericRepresentationOfLetter(item)
	}

	return total
}

func part2(file *os.File) int {
	scanner := bufio.NewScanner(file)

	total := 0

	current := 1
	var firstGroup string
	var secondGroup string
	var thirdGroup string
	for scanner.Scan() {
		switch current {
		case 1:
			firstGroup = scanner.Text()
			current++
		case 2:
			secondGroup = scanner.Text()
			current++
		case 3:
			thirdGroup = scanner.Text()
			c := getCharacterFromGroups(firstGroup, secondGroup, thirdGroup)
			total += getNumericRepresentationOfLetter(c)
			current = 1
		default:
			panic("invalid current value")
		}
	}

	return total
}

func getCharacterFromGroups(s, ss, sss string) string {
	// for a group of 3 strings, find the character that appears in all 3 strings
	for i := 0; i < len(s); i++ {
		if strings.Contains(ss, string(s[i])) && strings.Contains(sss, string(s[i])) {
			return string(s[i])
		}
	}
	panic("no character found")
}

func getDuplicate(s string) string {
	for i := 0; i < len(s)/2; i++ {
		for j := len(s) - 1; j > len(s)/2-1; j-- {
			if s[i] == s[j] {
				return string(s[i])
			}
		}
	}
	panic("no duplicate found")
}

func getNumericRepresentationOfLetter(letter string) int {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
		"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52,
	}
	return m[letter]
}
