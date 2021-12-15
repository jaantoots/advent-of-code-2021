package main

import (
	"fmt"
	"io"
)

type Cave struct {
	risk []int
	width int
}

func (c *Cave) height() int {
	if c.width > 0 {
		return len(c.risk)/c.width
	}
	return 0
}

func (c *Cave) at(x, y int) *int {
	if x >= 0 && x < c.width && y >= 0 && y < c.height() {
		return &c.risk[y*c.width + x]
	}
	return nil
}

func (c Cave) String() string {
	var str string
	for i := 0; i < c.height(); i++ {
		str += fmt.Sprintf("%v", c.risk[i*c.width:(i+1)*c.width]) + "\n"
	}
	return str
}

func (c *Cave) lowest_risk(lowest *Cave, pre_cost, x, y int) {
	cost := c.at(x, y)
	if cost == nil {
		return
	}
	total := pre_cost + *cost
	best := lowest.at(x, y)
	if *best > 0 && total >= *best {
		return
	}
	*best = total
	c.lowest_risk(lowest, total, x - 1, y)
	c.lowest_risk(lowest, total, x + 1, y)
	c.lowest_risk(lowest, total, x, y - 1)
	c.lowest_risk(lowest, total, x, y + 1)
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
	lowest := Cave{make([]int, len(cave.risk)), cave.width}
	cave.lowest_risk(&lowest, 0, 0, 0)
	//fmt.Println(lowest)
	fmt.Println(lowest.risk[len(lowest.risk)-1]-lowest.risk[0])
	//var risk int
	//for y := 0; y < len(grid)/width; y++ {
	//	for x := 0; x < width; x++ {
	//		risk += get_risk_level(grid, width, x, y)
	//	}
	//}
	//fmt.Println(risk)
}
