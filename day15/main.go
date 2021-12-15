package main

import (
	"fmt"
	"io"
)

type Cave struct {
	risk  []int
	width int
}

func (c *Cave) height() int {
	if c.width > 0 {
		return len(c.risk) / c.width
	}
	return 0
}

func (c *Cave) extend(n int) Cave {
	big := Cave{make([]int, len(c.risk)*n*n), c.width * n}
	h, w := c.height(), c.width
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					big.risk[j*h*n*w+i*w+y*n*w+x] = (c.risk[y*w+x]+i+j-1)%9 + 1
				}
			}
		}
	}
	return big
}

func (c Cave) String() string {
	var str string
	for i := 0; i < c.height(); i++ {
		str += fmt.Sprintf("%v", c.risk[i*c.width:(i+1)*c.width]) + "\n"
	}
	return str
}

type Point [2]int

func (c *Cave) try_update(lowest *Cave, stack *map[Point]bool, pre_cost int, p Point) {
	x, y := p[0], p[1]
	if x >= 0 && x < c.width && y >= 0 && y < c.height() {
		if pre_cost+c.risk[y*c.width+x] < lowest.risk[y*c.width+x] || lowest.risk[y*c.width+x] == 0 {
			lowest.risk[y*c.width+x] = pre_cost + c.risk[y*c.width+x]
			(*stack)[p] = true
		}
	}
}

func (c *Cave) lowest_risk() Cave {
	lowest := Cave{make([]int, len(c.risk)), c.width}
	stack := make(map[Point]bool, len(c.risk))
	stack[Point{0, 0}] = true
	for len(stack) > 0 {
		for p := range stack {
			cost := lowest.risk[p[1]*c.width+p[0]]
			c.try_update(&lowest, &stack, cost, Point{p[0] - 1, p[1]})
			c.try_update(&lowest, &stack, cost, Point{p[0] + 1, p[1]})
			c.try_update(&lowest, &stack, cost, Point{p[0], p[1] - 1})
			c.try_update(&lowest, &stack, cost, Point{p[0], p[1] + 1})
			delete(stack, p)
		}
	}
	return lowest
}

func main() {
	var line string
	var risk []int
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
		risk = append(risk, row...)
	}
	cave := Cave{risk, width}
	//fmt.Println(cave)
	lowest := cave.lowest_risk()
	//fmt.Println(lowest)
	fmt.Println(lowest.risk[len(lowest.risk)-1])

	big := cave.extend(5)
	//fmt.Println(big)
	bowest := big.lowest_risk()
	//fmt.Println(bowest)
	fmt.Println(bowest.risk[len(bowest.risk)-1])
}
