package main

import (
	"fmt"
	"io"
	"strconv"
)

type Input struct {
	raw *[14]int
	idx int
}

func (i *Input) Next() int {
	defer func() { i.idx++ }()
	return i.raw[i.idx]
}

type State struct {
	vars [5]int // w, x, y, z, tmp
	input *Input
}

func NewState(input *Input) State {
	return State{[5]int{0, 0, 0, 0, 0}, input}
}

func parseInstruction(ins, aa, bb string) func(*State) {
	var a, b, value int
	switch aa {
	case "w":
		a = 0
	case "x":
		a = 1
	case "y":
		a = 2
	case "z":
		a = 3
	default:
		panic(fmt.Sprintf("invalid value for a: %s", aa))
	}
	if ins != "inp" {
		switch bb {
		case "w":
			b = 0
		case "x":
			b = 1
		case "y":
			b = 2
		case "z":
			b = 3
		default:
			b = 4
			c, err := strconv.Atoi(bb)
			if err != nil {
				panic(fmt.Sprintf("invalid value for b: %s", bb))
			}
			value = c
		}
	}
	//fmt.Printf("parsed: %v, %v, %v, %v\n", ins, a, b, value)
	switch ins {
	case "inp":
		return func(s *State) {
			s.vars[a] = s.input.Next()
		}
	case "add":
		return func(s *State) {
			s.vars[4] = value
			s.vars[a] += s.vars[b]
		}
	case "mul":
		return func(s *State) {
			s.vars[4] = value
			s.vars[a] *= s.vars[b]
		}
	case "div":
		return func(s *State) {
			s.vars[4] = value
			s.vars[a] /= s.vars[b]
		}
	case "mod":
		return func(s *State) {
			s.vars[4] = value
			s.vars[a] %= s.vars[b]
		}
	case "eql":
		return func(s *State) {
			s.vars[4] = value
			if s.vars[a] == s.vars[b] {
				s.vars[a] = 1
			} else {
				s.vars[a] = 0
			}
		}
	default:
		panic(fmt.Sprintf("invalid instruction: %s", ins))
	}
}

func runProgram(program []func(*State), num *[14]int) bool {
	input := Input{num, 0}
	state := NewState(&input)
	for _, f := range program {
		f(&state)
	}
	//fmt.Println(state.vars, *state.input.raw)
	if state.vars[3] == 0 {
		return true
	}
	return false
}

func findLargest(program []func(*State)) [14]int {
	var serial [14]int
	serial[0] = 10
	for {
		serial[len(serial)-1]--
		for i := len(serial) - 1; i >= 0; i-- {
			if serial[i] <= 0 {
				if i == 0 {
					return [14]int{}
				}
				serial[i] = 9
				serial[i-1]--
				if i < 8 {
					fmt.Println(serial)
				}
			}
		}
		if runProgram(program, &serial) {
			return serial
		}
	}
}

func main() {
	var program []func(*State)
	var ins, a, b string
	for {
		n, err := fmt.Scanln(&ins, &a, &b)
		if err == io.EOF {
			break
		}
		if n < 2 || (n < 3 && ins != "inp") {
			panic(err)
		}
		program = append(program, parseInstruction(ins, a, b))
	}

	//fmt.Println(runProgram(program, []int{1,3,5,7,9,2,4,6,8,9,9,9,9,9}))
	//fmt.Println(runProgram(program, &[14]int{9,9,9,9,9,9,9,9,9,9,9,8,3,2}))
	//fmt.Println(runProgram(program, &[14]int{8,9,9,9,9,9,9,9,9,9,9,8,3,2}))
	//fmt.Println(runProgram(program, &[14]int{7,9,9,9,9,9,9,9,9,9,9,8,3,2}))
	//fmt.Println(runProgram(program, &[14]int{6,9,9,9,9,9,9,9,9,9,9,8,3,2}))
	//fmt.Println(runProgram(program, &[14]int{5,9,9,9,9,9,9,9,9,9,9,8,3,2}))
	//fmt.Println(runProgram(program, &[14]int{4,9,9,9,9,9,9,9,9,9,9,8,3,2}))
	//fmt.Println(runProgram(program, &[14]int{3,9,9,9,9,9,9,9,9,9,9,8,3,2}))
	//fmt.Println(runProgram(program, &[14]int{2,9,9,9,9,9,9,9,9,9,9,8,3,2}))
	//fmt.Println(runProgram(program, &[14]int{1,9,9,9,9,9,9,9,9,9,9,8,3,2}))

	fmt.Println(findLargest(program))
}
