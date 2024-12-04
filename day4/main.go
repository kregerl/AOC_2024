package main

import (
	"fmt"
	"os"
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
	var board [][]string

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}
	count := 0
	for row_index, row := range board {
		for col_index := range row {
			if board[row_index][col_index] == "X" {
				count += find_word(board, "XMAS", row_index, col_index)
			}
		}
	}

	fmt.Println("Part 1: ", count)
}

func find_word(board [][]string, word string, row_index int, col_index int) int {
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	count := 0

	for _, dir := range directions {
		current_row, current_col := row_index, col_index
		found := true

		for _, char := range word[1:] {
			current_row += dir[0]
			current_col += dir[1]

			if current_row < 0 || current_row >= len(board) || current_col < 0 || current_col >= len(board[0]) || board[current_row][current_col] != string(char) {
				found = false
				break
			}
			found = true
		}
		if found {
			count++
		}
	}
	return count
}

func find_word_2(board [][]string, word string, row_index int, col_index int) bool {
	directions := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	found := true
	count := 0

	for _, dir := range directions {
		half_len := len(word) / 2
		current_row := row_index - (dir[0] * half_len)
		current_col := col_index - (dir[1] * half_len)
		for _, char := range word {
			if current_row < 0 || current_row >= len(board) || current_col < 0 || current_col >= len(board[0]) || board[current_row][current_col] != string(char) {
				found = false
				break
			}
			found = true
			current_row += dir[0]
			current_col += dir[1]
		}
		if found {
			count += 1
		}
	}
	return count == 2
}

func part2(s string) {
	var board [][]string

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		board = append(board, strings.Split(line, ""))
	}
	count := 0
	for row_index, row := range board {
		for col_index := range row {
			if board[row_index][col_index] == "A" && find_word_2(board, "MAS", row_index, col_index) {
				count++
			}
		}
	}
	fmt.Println("Part 2: ", count)
}
