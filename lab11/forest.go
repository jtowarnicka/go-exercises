package main

import (
	"fmt"
	"math/rand"
)

var trees int = 0
var burnt int = 0

func main() {
	var forest [4][25]int
	generateForest(&forest)
	printer(forest)
	setThunder(&forest)

	defer calcStatistic(forest)
}

func generateForest(forest *[4][25]int) {
    for i := range forest {
        for j := range forest[i] {
            if rand.Float64() < 0.5 {
                forest[i][j] = 1
				trees++
            } else {
                forest[i][j] = 0
            }
        }
    }
}

func treeList(forest *[4][25]int) [][]int {
    var treeList [][]int
    for i := range forest {
        for j := range forest[i] {
            if forest[i][j] == 1 {
                treeList = append(treeList, []int{i, j})
            }
        }
    }
    return treeList
}

func setThunder(forest *[4][25]int) {
    t := treeList(forest)
    tree := t[rand.Intn(len(t))]
	forest[tree[0]][tree[1]] = 2
	printer(*forest)
	setOnFire(forest, tree[0], tree[1])
}

func setOnFire(forest *[4][25]int, x int, y int) {
	forest[x][y] = 3
	burnt++
	printer(*forest)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newX := x + i
			newY := y + j
			if newX < 0 || newX >= 4 || newY < 0 || newY >= 25 {
				continue
			}
			if forest[newX][newY] == 1 {
				setOnFire(forest, newX, newY)
			}
		}
	}
}

func printer(forest [4][25]int) {
	fmt.Println("-------------------------------------------------")
	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] == 0 {
				fmt.Print("\033[90m⍋ \033[0m")
			} else if forest[i][j] == 1 {
				fmt.Print("\033[32m⍋ \033[0m")
			} else if forest[i][j] == 2 {
				fmt.Print("\033[33m⍋ \033[0m")
			} else if forest[i][j] == 3 {
				fmt.Print("\033[31m⍋ \033[0m")
			} else if (forest[i][j] == 4) {
				fmt.Print("\033[38;5;208m⍋ \033[0m")
			}
		}
		fmt.Println()
	}
}

func calcStatistic(forest [4][25]int) {
	fmt.Println("-------------------------------------------------")
	fmt.Println("Trees: ", trees)
	fmt.Println("Burnt: ", burnt)
	fmt.Println("Density: ", float64(trees) / float64(100))
	result := float64(burnt) / float64(trees) * 100
	fmt.Printf("Burnt trees: %.2f%%\n", result)
}
