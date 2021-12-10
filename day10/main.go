package main

import (
	"fmt"
	"io"
	"sort"
)

func check_match(a, b rune) bool {
	switch {
	case a == '(' && b == ')':
		return true
	case a == '[' && b == ']':
		return true
	case a == '{' && b == '}':
		return true
	case a == '<' && b == '>':
		return true
	}
	return false
}

func main() {
	var score int
	comps := make([]int, 0)
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		buffer := make([]rune, len(line))
		var pos int
		var ill rune
		var corrupt bool
	LineParsing:
		for _, c := range line {
			switch c {
			case ')', ']', '}', '>':
				if pos > 0 && check_match(buffer[pos-1], c) {
					pos--
				} else {
					ill = c
					corrupt = true
					break LineParsing
				}
			case '(', '[', '{', '<':
				buffer[pos] = c
				pos++
			default:
				panic(fmt.Sprintf("unexpected character: %s", c))
			}
		}
		// fmt.Println(buffer)
		switch ill {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		}

		if corrupt {
			continue
		}
		var comp int
		for i := pos - 1; i >= 0; i-- {
			comp *= 5
			switch buffer[i] {
			case '(':
				comp += 1
			case '[':
				comp += 2
			case '{':
				comp += 3
			case '<':
				comp += 4
			}
		}
		// fmt.Println(line, comp)
		comps = append(comps, comp)
	}
	fmt.Println(score)

	sort.Ints(comps)
	fmt.Println(comps[len(comps)/2])
}
