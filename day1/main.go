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

func part1(s string) {
	array1, array2 := parse_numbers(s)
	slices.Sort(array1)
	slices.Sort(array2)

	sum := 0
	for i := range len(array1) {
		sum += int(math.Abs(float64(array1[i] - array2[i])))
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func part2(s string) {
	array1, array2 := parse_numbers(s)

	sum := 0
	for i := range len(array1) {
		num := array1[i]
		count := 0
		for _, element := range array2 {
			if element == num {
				count++
			}
		}
		sum += num * count
	}
	fmt.Printf("Part 2: %d\n", sum)
}

func parse_numbers(s string) ([]int, []int) {
	lines := strings.Split(s, "\n")
	var array1 []int
	var array2 []int

	for _, line := range lines {
		substrs := strings.Split(line, "   ")
		n1, _ := strconv.Atoi(substrs[0])
		n2, _ := strconv.Atoi(substrs[1])

		array1 = append(array1, n1)
		array2 = append(array2, n2)
	}
	return array1, array2
}
