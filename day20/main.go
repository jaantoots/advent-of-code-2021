package main

import (
	"fmt"
	"io"
)

type Image [][]bool

func (img *Image) height() int {
	return len(*img)
}

func (img *Image) width() int {
	if len(*img) == 0 {
		return 0
	}
	return len((*img)[0])
}

func (img *Image) Validate() bool {
	width := img.width()
	for i, row := range *img {
		if len(row) != width {
			fmt.Printf("line %d has length %d, expected %d\n", i, len(row), width)
			return false
		}
	}
	return true
}

func (img Image) String() string {
	var str string
	for _, row := range img {
		rs := make([]rune, len(row)+1)
		for j, p := range row {
			if p {
				rs[j] = '#'
			} else {
				rs[j] = '.'
			}
		}
		rs[len(row)] = '\n'
		str += string(rs)
	}
	return str
}

func (img *Image) at(x, y int, bg bool) bool {
	if 0 <= x && x <= img.width() && 0 <= y && y <= img.height() {
		return (*img)[y][x]
	}
	return bg
}

func (img *Image) extend(n int, bg bool) *Image {
	height, width := img.height()+2*n, img.width()+2*n
	big := make([]bool, height*width)
	for i := range big {
		big[i] = bg
	}
	out := make(Image, height)
	for i := range out {
		out[i] = big[i*width : (i+1)*width]
	}
	for i, row := range *img {
		for j, p := range row {
			out[i+2][j+2] = p
		}
	}
	return &out
}

func (img *Image) get3x3at(x, y int) int {
	var val int
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			val <<= 1
			if (*img)[j][i] {
				val |= 1
			}
		}
	}
	return val
}

func (img *Image) Enhance(algo []bool, bg bool) (Image, bool) {
	ext := img.extend(2, bg)
	height, width := img.height()+2, img.width()+2
	out := make(Image, height)
	for j := range out {
		row := make([]bool, width)
		for i := range row {
			row[i] = algo[ext.get3x3at(i+1, j+1)]
		}
		out[j] = row
	}
	newBg := algo[0]
	if bg {
		newBg = algo[511]
	}
	return out, newBg
}

func (img *Image) Lit() int {
	var count int
	for _, row := range *img {
		for _, p := range row {
			if p {
				count++
			}
		}
	}
	return count
}

func main() {
	var line string
	_, err := fmt.Scanln(&line)
	if err != nil {
		panic(err)
	}
	algo := make([]bool, len(line))
	for i, c := range line {
		if c == '#' {
			algo[i] = true
		}
	}
	fmt.Println(len(algo))
	fmt.Scanln(&line)
	image := make(Image, 0, 100)
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		row := make([]bool, len(line))
		for i, c := range line {
			if c == '#' {
				row[i] = true
			}
		}
		image = append(image, row)
	}
	if !image.Validate() {
		panic("invalid image")
	}
	//fmt.Println(image)
	img := image
	bg := false
	for i := 0; i < 2; i++ {
		img, bg = img.Enhance(algo, bg)
		fmt.Println(i+1, img.Lit())
	}

	img = image
	bg = false
	for i := 0; i < 50; i++ {
		img, bg = img.Enhance(algo, bg)
		fmt.Println(i+1, img.Lit())
	}
	fmt.Println(img.Lit())
}
