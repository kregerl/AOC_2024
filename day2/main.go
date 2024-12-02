package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fi, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(fi)
	part1(str)
	part2(str)
}

func remove_index(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func all_increasing_or_decreasing(numbers []string, should_retry bool) bool {
	d := 0
	for i := 1; i < len(numbers); i++ {
		n1, _ := strconv.Atoi(numbers[i-1])
		n2, _ := strconv.Atoi(numbers[i])
		distance := n1 - n2
		if distance*d < 0 || math.Abs(float64(distance)) > 3 || math.Abs(float64(distance)) < 1 {
			if should_retry {
				copy1 := slices.Clone(numbers)
				copy1 = remove_index(copy1, i)
				copy2 := slices.Clone(numbers)
				copy2 = remove_index(copy2, i-1)

				if i == 2 {
					copy3 := slices.Clone(numbers)
					copy3 = remove_index(copy3, i-2)
					return all_increasing_or_decreasing(copy1, false) || all_increasing_or_decreasing(copy2, false) || all_increasing_or_decreasing(copy3, false)
				}
				return all_increasing_or_decreasing(copy1, false) || all_increasing_or_decreasing(copy2, false)
			}
			return false
		}

		d += distance
	}
	return true
}

func part1(s string) {
	count := 0

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		if all_increasing_or_decreasing(numbers, false) {
			count++
		}
	}
	fmt.Println("Part 1: ", count)

}

func part2(s string) {
	count := 0

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		fmt.Println(numbers)
		if all_increasing_or_decreasing(numbers, true) {
			count++
		}
	}
	fmt.Println("Part 2: ", count)
}
