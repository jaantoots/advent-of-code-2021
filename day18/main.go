package main

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Element struct {
	num  int
	pair *Pair
}

type Pair struct {
	first, second *Element
}

func (e Element) String() string {
	if e.pair == nil {
		return fmt.Sprintf("%d", e.num)
	}
	return fmt.Sprintf("[%v %v]", e.pair.first, e.pair.second)
}

var re = regexp.MustCompile(`^\d*`)

func ParseElement(s string) (*Element, string) {
	num := re.FindString(s)
	if len(num) > 0 {
		i, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		return &Element{i, nil}, s[len(num):]
	}
	var first, second *Element
	if s[0] != '[' {
		panic("parsing error")
	}
	first, s = ParseElement(s[1:])
	if s[0] != ',' {
		panic("parsing error")
	}
	second, s = ParseElement(s[1:])
	if s[0] != ']' {
		panic("parsing error")
	}
	return &Element{0, &Pair{first, second}}, s[1:]
}

func (a *Element) add(b *Element) *Element {
	result := Element{0, &Pair{a, b}}
	result.reduce()
	return &result
}

func (e *Element) reduce() {
	for {
		_, _, flag := e.explode(0)
		if flag {
			continue
		}
		if e.split() {
			continue
		}
		break
	}
}

func (e *Element) add_left(x int) {
	if e.pair == nil {
		e.num += x
		return
	}
	e.pair.first.add_left(x)
}

func (e *Element) add_right(x int) {
	if e.pair == nil {
		e.num += x
		return
	}
	e.pair.second.add_right(x)
}

func (e *Element) explode(depth int) (int, int, bool) {
	if e.pair == nil {
		return 0, 0, false
	}
	if depth > 4 {
		fmt.Printf("warning: too deep: %d\n", depth)
	}
	var left, right int
	var flag bool
	left, right, flag = e.pair.first.explode(depth + 1)
	if flag {
		e.pair.second.add_left(right)
		return left, 0, true
	}
	left, right, flag = e.pair.second.explode(depth + 1)
	if flag {
		e.pair.first.add_right(left)
		return 0, right, true
	}
	if depth >= 4 {
		// Both must be numbers, otherwise would have returned above
		left, right = e.pair.first.num, e.pair.second.num
		e.num = 0
		e.pair = nil
		return left, right, true
	}
	return 0, 0, false
}

func (e *Element) split() bool {
	if e.pair != nil {
		return e.pair.first.split() || e.pair.second.split()
	}
	if e.num >= 10 {
		first := e.num / 2
		e.pair = &Pair{&Element{first, nil}, &Element{e.num - first, nil}}
		e.num = 0
		return true
	}
	return false
}

func (e *Element) magnitude() int {
	if e.pair == nil {
		return e.num
	}
	return 3*e.pair.first.magnitude() + 2*e.pair.second.magnitude()
}

func main() {
	var line string
	var sum, elem *Element
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		elem, _ = ParseElement(line)
		//fmt.Printf("elem: %v\n", elem)
		if sum == nil {
			sum = elem
		} else {
			sum = sum.add(elem)
		}
		//fmt.Printf("sum: %v\n", sum)
	}
	fmt.Println(sum)
	fmt.Println(sum.magnitude())
}
