package main

import (
	"fmt"
)

type Fold struct {
	axis  rune
	value int
}

type Paper struct {
	dots [][2]int
}

func (p *Paper) grid() ([]bool, int) {
	var min_x, min_y, max_x, max_y int
	if len(p.dots) > 0 {
		min_x, min_y = p.dots[0][0], p.dots[0][1]
		max_x, max_y = min_x, min_y
	}
	for _, dot := range p.dots {
		x, y := dot[0], dot[1]
		if x > max_x {
			max_x = x
		}
		if y > max_y {
			max_y = y
		}
		if x < min_x {
			min_x = x
		}
		if y < min_y {
			min_y = y
		}
	}
	width := (max_x - min_x) + 1
	height := (max_y - min_y) + 1
	arr := make([]bool, width*height)
	for _, dot := range p.dots {
		x, y := dot[0]-min_x, dot[1]-min_y
		arr[y*width+x] = true
	}
	return arr, width
}

func (p Paper) String() string {
	arr, width := p.grid()
	runes := make([]rune, len(arr)+len(arr)/width)
	for i := 0; i < len(arr)/width; i++ {
		for j := 0; j < width; j++ {
			if arr[i*width+j] {
				runes[i*(width+1)+j] = '#'
			} else {
				runes[i*(width+1)+j] = '.'
			}
		}
		runes[i*(width+1)+width] = '\n'
	}
	return string(runes)
}

func contains(dots [][2]int, dot [2]int) bool {
	for _, comp := range dots {
		if dot == comp {
			return true
		}
	}
	return false
}

func (p *Paper) apply_fold(fold Fold) Paper {
	dots := make([][2]int, 0, len(p.dots))
	var axis int
	switch fold.axis {
	case 'x':
		axis = 0
	case 'y':
		axis = 1
	}
	for _, dot := range p.dots {
		if dot[axis] > fold.value {
			dot = [2]int{dot[0], dot[1]}
			dot[axis] = 2*fold.value - dot[axis]
		}
		if !contains(dots, dot) {
			dots = append(dots, dot)
		}
	}
	return Paper{dots}
}

func main() {
	var x, y int
	dots := make([][2]int, 0, 256)
	for {
		_, err := fmt.Scanf("%d,%d", &x, &y)
		if err != nil {
			break
		}
		dots = append(dots, [2]int{x, y})
	}
	//fmt.Println(dots)
	var axis rune
	var val int
	folds := make([]Fold, 0, 64)
	for {
		_, err := fmt.Scanf("fold along %c=%d", &axis, &val)
		if err != nil {
			break
		}
		folds = append(folds, Fold{axis, val})
	}
	//fmt.Println(folds)
	paper := Paper{dots}
	//fmt.Println(paper)
	folded_once := paper.apply_fold(folds[0])
	//fmt.Println(folded_once)
	fmt.Println(len(folded_once.dots))
	p := paper
	for _, fold := range folds {
		p = p.apply_fold(fold)
	}
}
