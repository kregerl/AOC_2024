package main

import (
	"fmt"
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
	lines := strings.Split(s, "\n")

	order_map := make(map[int][]int)
	is_rules := true
	sum := 0
	for _, line := range lines {
		if line == "" {
			is_rules = false
			continue
		}

		if is_rules {
			nums := strings.Split(line, "|")
			before, _ := strconv.Atoi(nums[0])
			after, _ := strconv.Atoi(nums[1])

			if _, ok := order_map[after]; !ok {
				order_map[after] = []int{before}
			} else {
				order_map[after] = append(order_map[after], before)
			}
			continue
		}

		updates := strings.Split(line, ",")
		midpoint := len(updates) / 2
		is_valid := true

		for i := 0; i < len(updates)-1; i++ {
			j := i + 1
			page1, _ := strconv.Atoi(updates[i])
			page2, _ := strconv.Atoi(updates[j])

			page1_values, page1_ok := order_map[page1]
			if page1_ok && slices.Contains(page1_values, page2) {
				is_valid = false
				break
			}
		}

		if is_valid {
			v, _ := strconv.Atoi(updates[midpoint])
			sum += v
		}
	}

	fmt.Println("Part 1: ", sum)
}

func find_error(order_map map[int][]int, updates []string, midpoint int) int {
	for i := 0; i < len(updates)-1; i++ {
		j := i + 1
		page1, _ := strconv.Atoi(updates[i])
		page2, _ := strconv.Atoi(updates[j])

		page1_values, page1_ok := order_map[page1]
		if page1_ok && slices.Contains(page1_values, page2) {
			return i
		}
	}
	return -1
}

func part2(s string) {
	lines := strings.Split(s, "\n")

	order_map := make(map[int][]int)
	is_rules := true
	sum := 0
	for _, line := range lines {
		if line == "" {
			is_rules = false
			continue
		}

		if is_rules {
			nums := strings.Split(line, "|")
			before, _ := strconv.Atoi(nums[0])
			after, _ := strconv.Atoi(nums[1])

			if _, ok := order_map[after]; !ok {
				order_map[after] = []int{before}
			} else {
				order_map[after] = append(order_map[after], before)
			}
			continue
		}

		updates := strings.Split(line, ",")
		midpoint := len(updates) / 2

		err := find_error(order_map, updates, midpoint)

		if err == -1 {
			continue
		}

		for ok := true; ok; ok = (err != -1) {
			k := updates[err]
			l := updates[err+1]
			updates[err] = l
			updates[err+1] = k
			err = find_error(order_map, updates, midpoint)
		}
		v, _ := strconv.Atoi(updates[midpoint])
		sum += v
	}

	fmt.Println("Part 2: ", sum)
}
