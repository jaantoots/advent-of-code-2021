package main

import (
	"fmt"
	"io"
)

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

type Grid struct {
	width int
	arr   []int
}

func (g *Grid) height() int {
	return len(g.arr) / g.width
}

func (g Grid) String() string {
	var str string
	for i := 0; i < g.height(); i++ {
		for _, x := range g.arr[i*g.width : (i+1)*g.width] {
			str += fmt.Sprintf("%2d", x)
		}
		str += "\n"
	}
	return str
}

func (g *Grid) incr_adj(x, y int) {
	for i := max(0, x-1); i <= min(g.width-1, x+1); i++ {
		for j := max(0, y-1); j <= min(g.height()-1, y+1); j++ {
			g.arr[j*g.width+i]++
			if g.arr[j*g.width+i] == 10 {
				g.incr_adj(i, j)
			}
		}
	}
}

func (g *Grid) step() int {
	// Increment energy levels
	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height(); j++ {
			g.arr[j*g.width+i]++
			if g.arr[j*g.width+i] == 10 {
				g.incr_adj(i, j)
			}
		}
	}
	//fmt.Println(g)
	// Count flashes
	var flashes int
	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height(); j++ {
			if g.arr[j*g.width+i] > 9 {
				flashes++
				g.arr[j*g.width+i] = 0
			}
		}
	}
	//fmt.Println(g)
	return flashes
}

func main() {
	var line string
	var arr []int
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
		//fmt.Println(row)
		arr = append(arr, row...)
	}
	//fmt.Println(arr)
	grid := Grid{width, arr}
	//fmt.Println(grid)
	var total int
	for i := 0; i < 100; i++ {
		total += grid.step()
		//fmt.Println(grid)
	}
	fmt.Println(total)
}
