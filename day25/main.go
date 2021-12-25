package main

import (
	"fmt"
	"io"
)

type CucumberMap [][]int

func (c *CucumberMap) height() int {
	return len(*c)
}

func (c *CucumberMap) width() int {
	if len(*c) > 0 {
		return len((*c)[0])
	}
	return 0
}

func (c *CucumberMap) Valid() bool {
	width := c.width()
	for _, row := range *c {
		if len(row) != width {
			return false
		}
	}
	return true
}

func (c *CucumberMap) Step() int {
	var count int
	// East-facing
	for j := 0; j < c.height(); j++ {
		fwd := (*c)[j][0]
		// Do not look at 0 if it was empty
		end := 0
		if fwd == 0 {
			end = 1
		}
		for i := c.width() - 1; i >= end; i-- {
			if fwd == 0 && (*c)[j][i] == 1 {
				fwd = (*c)[j][i]
				(*c)[j][(i+1)%c.width()] = (*c)[j][i]
				(*c)[j][i] = 0
				count++
			} else {
				fwd = (*c)[j][i]
			}
		}
	}
	//fmt.Println(*c)
	// South-facing
	for i := 0; i < c.width(); i++ {
		fwd := (*c)[0][i]
		// Do not look at 0 if it was empty
		end := 0
		if fwd == 0 {
			end = 1
		}
		for j := c.height() - 1; j >= end; j-- {
			if fwd == 0 && (*c)[j][i] == 2 {
				fwd = (*c)[j][i]
				(*c)[(j+1)%c.height()][i] = (*c)[j][i]
				(*c)[j][i] = 0
				count++
			} else {
				fwd = (*c)[j][i]
			}
		}
	}
	return count
}

func (c CucumberMap) String() string {
	var str string
	for _, row := range c {
		p := make([]rune, len(row))
		for i, x := range row {
			switch x {
			case 0:
				p[i] = '.'
			case 1:
				p[i] = '>'
			case 2:
				p[i] = 'v'
			default:
				p[i] = '?'
			}
		}
		str += string(p) + "\n"
	}
	return str
}

func main() {
	var line string
	var cucumbers CucumberMap
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		row := make([]int, len(line))
		for i, c := range line {
			switch c {
			case '>':
				row[i] = 1
			case 'v':
				row[i] = 2
			}
		}
		cucumbers = append(cucumbers, row)
	}
	if !cucumbers.Valid() {
		panic("invalid map")
	}
	//fmt.Println(cucumbers)

	steps := 0
	for {
		steps++
		count := cucumbers.Step()
		//fmt.Println(steps)
		//fmt.Println(cucumbers)
		if count == 0 {
			break
		}
	}
	fmt.Println(steps)
}
