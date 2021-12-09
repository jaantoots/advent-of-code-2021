package main

import (
	"fmt"
	"io"
	"sort"
)

func explore_basin(grid []int, visited []bool, width, x, y int) int {
	if visited[y*width+x] || grid[y*width+x] == 9 {
		return 0
	}
	visited[y*width+x] = true
	size := 1
	if x > 0 {
		size += explore_basin(grid, visited, width, x-1, y)
	}
	if y > 0 {
		size += explore_basin(grid, visited, width, x, y-1)
	}
	if x < width-1 {
		size += explore_basin(grid, visited, width, x+1, y)
	}
	if y < len(grid)/width-1 {
		size += explore_basin(grid, visited, width, x, y+1)
	}
	// fmt.Println(size)
	return size
}

func main() {
	var line string
	var grid []int
	var width int
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if width > 0 && width != len(line) {
			panic("width mismatch")
		}
		width = len(line)
		row := make([]int, len(line))
		for i, x := range line {
			row[i] = int(x) - 48
		}
		// fmt.Println(row)
		grid = append(grid, row...)
	}
	// fmt.Println(grid)
	visited := make([]bool, len(grid))
	basins := make([]int, len(grid))
	for y := 0; y < len(grid)/width; y++ {
		for x := 0; x < width; x++ {
			basins[y*width+x] = explore_basin(grid, visited, width, x, y)
		}
	}
	// fmt.Println(visited, basins)
	sort.Ints(basins)
	// fmt.Println(basins)
	fmt.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}
