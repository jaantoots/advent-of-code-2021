package main

import (
	"fmt"
	"io"
)

func get_risk_level(grid []int, width, x, y int) int {
	p := grid[y*width+x]
	switch {
	case x > 0 && grid[y*width+x-1] <= p:
		return 0
	case y > 0 && grid[(y-1)*width+x] <= p:
		return 0
	case x < width-1 && grid[y*width+x+1] <= p:
		return 0
	case y < len(grid)/width-1 && grid[(y+1)*width+x] <= p:
		return 0
	}
	// fmt.Println(x, y)
	return p + 1
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
	var risk int
	for y := 0; y < len(grid)/width; y++ {
		for x := 0; x < width; x++ {
			risk += get_risk_level(grid, width, x, y)
		}
	}
	fmt.Println(risk)
}
