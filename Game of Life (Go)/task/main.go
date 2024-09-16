package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Print("Enter the size of the universe: ")
	fmt.Scan(&n)

	rand.Seed(time.Now().UnixNano())
	current := make([][]bool, n)
	next := make([][]bool, n)

	for i := range current {
		current[i] = make([]bool, n)
		next[i] = make([]bool, n)
		for j := range current[i] {
			current[i][j] = rand.Intn(2) == 1
		}
	}

	for gen := 1; gen <= 10; gen++ {
		alive := countAlive(current, n)
		fmt.Printf("Generation #%d\nAlive: %d\n", gen, alive)
		printGrid(current, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				liveNeighbors := countLiveNeighbors(current, i, j, n)
				next[i][j] = liveNeighbors == 3 || (current[i][j] && liveNeighbors == 2)
			}
		}
		current, next = next, current
	}
}

func countLiveNeighbors(grid [][]bool, x, y, size int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				nx, ny := (x+dx+size)%size, (y+dy+size)%size
				if grid[nx][ny] {
					count++
				}
			}
		}
	}
	return count
}

func printGrid(grid [][]bool, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func countAlive(grid [][]bool, size int) int {
	alive := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] {
				alive++
			}
		}
	}
	return alive
}
