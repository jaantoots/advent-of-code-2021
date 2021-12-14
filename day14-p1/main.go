package main

import (
	"fmt"
	"io"
)

type Rules map[[2]rune]rune

type Polymer []rune

func (p Polymer) String() string {
	return string(p)
}

func (p Polymer) insert(rules Rules) Polymer {
	ext := make(Polymer, 0, 2*len(p))
	var key [2]rune
	for _, b := range p {
		key[0], key[1] = key[1], b
		c, ok := rules[key]
		if ok {
			ext = append(ext, c)
		}
		ext = append(ext, b)
	}
	return ext
}

func (p Polymer) counts() map[rune]int {
	counts := make(map[rune]int, 64)
	for _, c := range p {
		counts[c]++
	}
	return counts
}

func main() {
	var template string
	_, err := fmt.Scanln(&template)
	if err != nil {
		panic(err)
	}
	var line string
	fmt.Scanln(&line)
	var a, b, c rune
	rules := make(Rules, 64)
	for {
		_, err := fmt.Scanf("%c%c -> %c\n", &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		rules[[2]rune{a, b}] = c
	}
	polymer := Polymer(template)
	//fmt.Println(polymer)
	//fmt.Println(rules)
	for i := 0; i < 10; i++ {
		polymer = polymer.insert(rules)
		//fmt.Println(polymer)
	}
	counts := polymer.counts()
	//fmt.Println(counts)
	var min, max int
	for _, x := range counts {
		if x > max {
			max = x
		}
		if x < min || min == 0 {
			min = x
		}
	}
	//fmt.Println(min, max)
	fmt.Println(max - min)
}
