package main

import (
	"fmt"
	"io"
	"strings"
)

func scan_draw() []int {
	var line string
	_, err := fmt.Scanln(&line)
	if err != nil {
		panic(err)
	}

	numbers := strings.Split(line, ",")
	draw := make([]int, len(numbers))
	for i, s := range numbers {
		fmt.Sscanf("%d", s)
		_, err := fmt.Sscanf(s, "%d", &draw[i])
		if err != nil {
			panic(err)
		}
	}
	return draw
}

func scan_row(dest *[5]int) error {
	for i := range dest {
		_, err := fmt.Scan(&dest[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func scan_board() ([5][5]int, error) {
	var board [5][5]int

	_, err := fmt.Scanln()
	if err != nil {
		return board, err
	}

	for i := range board {
		err := scan_row(&board[i])
		if err != nil {
			return board, err
		}
		// fmt.Println(board[i])
	}
	return board, nil
}

func check_bingo(state *[5][5]bool) bool {
	// Check rows
	for i := range state {
		bingo := true
		for j := range state[i] {
			if !state[i][j] {
				bingo = false
				break
			}
		}
		if bingo {
			// fmt.Println("row", i)
			return true
		}
	}
	// Check columns
	for i := range state[0] {
		bingo := true
		for j := range state {
			if !state[j][i] {
				bingo = false
				break
			}
		}
		if bingo {
			// fmt.Println("col", i)
			return true
		}
	}
	return false
}

func update_board_state(board *[5][5]int, state *[5][5]bool, num int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == num {
				state[i][j] = true
			}
		}
	}
}

func play_bingo(board *[5][5]int, draw []int) (int, int) {
	var state [5][5]bool
	for i, num := range draw {
		update_board_state(board, &state, num)
		// fmt.Println(i, board, state, num)
		if check_bingo(&state) {
			var sum int
			for i := range board {
				for j := range board[i] {
					if !state[i][j] {
						sum += board[i][j]
					}
				}
			}
			// fmt.Println(sum, num)
			return i, sum * num
		}
	}
	return 0, 0
}

func main() {
	var best_idx int
	var best_score int

	draw := scan_draw()
	// fmt.Println(draw)
	for {
		board, err := scan_board()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// fmt.Println(board)

		idx, score := play_bingo(&board, draw)
		if best_idx == 0 || idx < best_idx {
			best_idx = idx
			best_score = score
		}
		// fmt.Println(idx, score, best_idx, best_score)
	}
	fmt.Println(best_score)
}
