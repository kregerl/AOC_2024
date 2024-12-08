package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

func FetchSliceOfIntsInString(line string) []int {
	nums := []int{}
	var build strings.Builder
	isNegative := false
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}

		if char == '-' {
			isNegative = true
		}

		if (char == ' ' || char == ',' || char == '~' || char == '|') && build.Len() != 0 {
			localNum, err := strconv.ParseInt(build.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			if isNegative {
				localNum *= -1
			}
			nums = append(nums, int(localNum))
			build.Reset()
			isNegative = false
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		if isNegative {
			localNum *= -1
		}
		nums = append(nums, int(localNum))
		build.Reset()
	}
	return nums
}

func main() {
	input := ReadFileLineByLine("input.txt")
	ans1, ans2 := ans(input)
	fmt.Println("Part 1: ", ans1)
	fmt.Println("Part 2: ", ans2)
}

func ans(input []string) (int, int) {
	ans1, ans2 := 0, 0
	for _, row := range input {
		nums := FetchSliceOfIntsInString(row)

		if isSumAMatch(nums[0], 0, nums[1:], false) {
			ans1 += nums[0]
		}

		if isSumAMatch(nums[0], 0, nums[1:], true) {
			ans2 += nums[0]
		}
	}
	return ans1, ans2
}

func calculate(a, b int, operation byte) int {
	calculation := 0
	switch operation {
	case '+':
		calculation = a + b
	case '*':
		calculation = a * b
	case '|':
		mul, q := 10, 10
		for q != 0 {
			q = b / mul
			if q > 0 {
				mul *= 10
			}
		}
		calculation = (a * mul) + b
	}
	return calculation
}

func isSumAMatch(expectedSum, sum int, input []int, isPart2 bool) bool {
	if len(input) == 0 {
		return sum == expectedSum
	}

	if sum > expectedSum {
		return false
	}

	if isSumAMatch(expectedSum, calculate(sum, input[0], '+'), input[1:], isPart2) {
		return true
	}

	if isPart2 && isSumAMatch(expectedSum, calculate(sum, input[0], '|'), input[1:], isPart2) {
		return true
	}
	return isSumAMatch(expectedSum, calculate(sum, input[0], '*'), input[1:], isPart2)
}
