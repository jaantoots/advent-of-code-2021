package main

import (
	"fmt"
	"io"
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
	LineParsing:
		for _, c := range line {
			switch c {
			case ')', ']', '}', '>':
				if pos > 0 && check_match(buffer[pos-1], c) {
					pos--
				} else {
					ill = c
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
	}
	fmt.Println(score)
}
