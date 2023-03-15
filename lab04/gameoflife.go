package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func generateGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	return grid
}

func rouseCells(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			// grid[i][j] = rand.Intn(2)
			if rand.Float64() < 0.5 {
                grid[i][j] = 1
            } else {
                grid[i][j] = 0
            }
		}
	}
}

func display(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				fmt.Print("\033[38;5;205m\u25AA\033[0m")
				// fmt.Print("\033[38;5;205m▒\033[0m")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func neighbors(grid [][]int, a, b, X, Y int) int {
	aliveNeighbours := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			x := (a + i)
			y := (b + j)

			if x < 0 || x >= X || y < 0 || y >= Y {
				continue
			}
			if grid[x][y] == 1 {
				aliveNeighbours++
			}
			
		}
	}

	aliveNeighbours -= grid[a][b]

	return aliveNeighbours
}

func next(grid [][]int, X int, Y int) [][]int {
	next := generateGrid(X, Y)

	for i := range grid {
		for j := range grid[i] {
			aliveNeighbours := neighbors(grid, i, j, X, Y)

            // 1. cell is lonely and dies
			// 2. cell is overcrowded and dies
			// 3. new cell is born
			// 4. cell is happy and stays alive
			if grid[i][j] == 1 && aliveNeighbours < 2 {
				next[i][j] = 0
			} else if grid[i][j] == 1 && aliveNeighbours > 3 {
				next[i][j] = 0
			} else if grid[i][j] == 0 && aliveNeighbours == 3 {
				next[i][j] = 1
			} else {
				next[i][j] = grid[i][j]
			}
		}
	}

	return next;
}

func main() {
	lines := flag.Int("lines", 20, "number of rows in the grid")
    cols := flag.Int("columns", 20, "number of columns in the grid")
    flag.Parse()

    grid := generateGrid(*lines, *cols)
	rouseCells(grid)
	for {
		fmt.Println("\033[H\033[2J")
		display(grid)
		grid = next(grid, *lines, *cols)
		time.Sleep(time.Second/8)
	}

	// command:
	// go run gameoflife.go -lines $LINES -columns $COLUMNS
}
