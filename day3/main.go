package main

import (
	"fmt"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	total := 0
	for _, match := range re.FindAllString(s, -1) {
		clean_match := strings.Replace(strings.Replace(match, "mul(", "", -1), ")", "", -1)
		values := strings.Split(clean_match, ",")
		n1, _ := strconv.Atoi(values[0])
		n2, _ := strconv.Atoi(values[1])
		total += n1 * n2
	}

	fmt.Println("Part 1: ", total)
}

func part2(s string) {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)

	total := 0
	enabled := true
	for _, match := range re.FindAllString(s, -1) {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		}
		if enabled && strings.HasPrefix(match, "mul(") {
			clean_match := strings.Replace(strings.Replace(match, "mul(", "", -1), ")", "", -1)
			values := strings.Split(clean_match, ",")
			n1, _ := strconv.Atoi(values[0])
			n2, _ := strconv.Atoi(values[1])
			total += n1 * n2
		}
	}

	fmt.Println("Part 2: ", total)
}
