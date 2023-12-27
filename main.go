package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

/*
Any live cell with fewer than two live neighbours dies, as if by underpopulation.
Any live cell with two or three live neighbours lives on to the next generation.
Any live cell with more than three live neighbours dies, as if by overpopulation.
Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/

func main() {
	boardY := 10
	boardX := 6
	board := make([][]string, boardY)
	for i := 0; i < boardY; i++ {
		board[i] = []string{".", "*", ".", "*", ".", "."}
	}

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		drawBoard(board)
		getNewBoard(board, boardY, boardX)
		time.Sleep(1 * time.Second)
	}
}

func getNewBoard(board [][]string, maxRow, maxCol int) [][]string {
	for row := range board {
		for col := range board[row] {
			cell := board[row][col]
			alive := getAliveNeighbours(board, row, col, maxRow, maxCol)
			if cell == "*" {
				if alive < 2 || alive > 3 {
					board[row][col] = "."
				} else if alive == 2 || alive == 3 {
					board[row][col] = "*"
				}
			} else {
				if alive == 3 {
					board[row][col] = "*"
				}
			}
		}
	}
	return board
}

func getAliveNeighbours(board [][]string, row, col, maxRow, maxCol int) int {
	neighbours := make([]string, 8)
	// Top
	if row-1 >= 0 {
		if col-1 >= 0 {
			neighbours = append(neighbours, board[row-1][col-1])
		}
		neighbours = append(neighbours, board[row-1][col])
		if col+1 < maxCol {
			neighbours = append(neighbours, board[row-1][col+1])
		}
	}
	// left
	if col-1 >= 0 {
		neighbours = append(neighbours, board[row][col-1])
	}
	// right
	if col+1 < maxCol {
		neighbours = append(neighbours, board[row][col+1])
	}
	// Bottom
	if row+1 < maxRow {
		if col-1 >= 0 {
			neighbours = append(neighbours, board[row+1][col-1])
		}
		neighbours = append(neighbours, board[row+1][col])
		if col+1 < maxCol {
			neighbours = append(neighbours, board[row+1][col+1])
		}
	}

	alive := 0
	for i := range neighbours {
		if neighbours[i] == "*" {
			alive++
		}
	}

	return alive
}

func drawBoard(board [][]string) {
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board[row]); column++ {
			fmt.Print(board[row][column], " ")
		}
		fmt.Print("\n")
	}
}
