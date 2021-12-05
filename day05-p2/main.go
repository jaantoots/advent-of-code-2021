package main

import (
	"fmt"
	"io"
)

type Line struct {
	x1, y1, x2, y2 int
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

func (g *Grid) mark(x, y int) int {
	if g.width < x+1 {
		// fmt.Println("extend x")
		arr := make([]int, 2*len(g.arr))
		for i := 0; i < g.height(); i++ {
			n := copy(arr[i*2*g.width:], g.arr[i*g.width:(i+1)*g.width])
			if n != g.width {
				panic("copy failed")
			}
		}
		g.arr = arr
		g.width *= 2
		return g.mark(x, y)
	}
	if g.height() < y+1 {
		// fmt.Println("extend y")
		arr := make([]int, 2*len(g.arr))
		n := copy(arr, g.arr)
		if n != len(g.arr) {
			panic("copy failed")
		}
		g.arr = arr
		return g.mark(x, y)
	}
	g.arr[y*g.width+x]++
	return 0
}

func (g *Grid) count_overlap() int {
	var count int
	for _, x := range g.arr {
		if x > 1 {
			count++
		}
	}
	return count
}

func (g *Grid) mark_line(line Line) {
	x1, x2 := line.x1, line.x2
	y1, y2 := line.y1, line.y2
	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	switch {
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			g.mark(x1, y)
		}
	case y1 == y2:
		for x := x1; x <= x2; x++ {
			g.mark(x, y1)
		}
	case x2-x1 == y2-y1:
		for i := 0; i <= x2-x1; i++ {
			g.mark(x1+i, y1+i)
		}
	case x2-x1 == y1-y2:
		for i := 0; i <= x2-x1; i++ {
			g.mark(x1+i, y1-i)
		}
	default:
		panic(fmt.Sprintf("invalid line: %v", line))
	}
}

func main() {
	grid := Grid{1, make([]int, 1)}
	for {
		var x1, y1, x2, y2 int
		_, err := fmt.Scanf("%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		line := Line{x1, y1, x2, y2}
		// fmt.Println(line)
		grid.mark_line(line)
		// fmt.Println(grid)
		// fmt.Println(len(grid.arr), grid.width)
	}
	fmt.Println(grid.count_overlap())
}
