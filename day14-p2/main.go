package main

import (
	"fmt"
	"io"
)

type Rules map[[2]rune]rune

type Polymer map[[2]rune]int

func NewPolymer(template string) *Polymer {
	var key [2]rune
	polymer := Polymer{}
	for _, b := range template {
		key[0], key[1] = key[1], b
		if key[0] == 0 {
			continue
		}
		polymer[key]++
	}
	return &polymer
}

func (p *Polymer) insert(rules Rules) *Polymer {
	ext := make(Polymer, len(*p))
	var key [2]rune
	for pair, count := range *p {
		c, ok := rules[pair]
		if ok {
			key[0], key[1] = pair[0], c
			ext[key] += count
			key[0], key[1] = c, pair[1]
			ext[key] += count
		} else {
			ext[pair] += count
		}
	}
	return &ext
}

func (p *Polymer) counts() map[rune]int {
	counts := make(map[rune]int, 64)
	for pair, count := range *p {
		counts[pair[0]] += count
		counts[pair[1]] += count
	}
	for c := range counts {
		// only first and last will not have been double counted
		counts[c] = (counts[c] + 1) / 2
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
	polymer := NewPolymer(template)
	//fmt.Println(polymer)
	//fmt.Println(rules)
	for i := 0; i < 40; i++ {
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
