package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("01.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Create map of elves and their calorie totals
	elves := make(map[int]int)
	count := 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			calories, _ := strconv.Atoi(scanner.Text())
			_, ok := elves[count]
			if !ok {
				elves[count] = 0
			}
			elves[count] += calories

		} else {
			count++
		}
	}

	// Sort the elves by their calorie totals
	type kv struct {
		Key   int
		Value int
	}

	var ss []kv
	for k, v := range elves {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	// Print results
	fmt.Println("Part one total:", ss[0].Value)

	var partTwoTotal int
	for i := 0; i < 3; i++ {
		partTwoTotal += ss[i].Value
	}

	fmt.Println("Part two total:", partTwoTotal)
}
