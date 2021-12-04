package main

import (
	"fmt"
	"io"
)

func main() {
	var line string
	var ones []uint
	var count uint
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if len(line) != len(ones) {
			if len(ones) > 0 {
				panic(fmt.Sprintf("bad string length: %s", line))
			}
			ones = make([]uint, len(line))
		}
		for i, s := range line {
			if s != '0' {
				ones[i]++
			}
		}
		count++
		// fmt.Println(line, count, ones)
	}

	var gamma uint
	for _, x := range ones {
		gamma <<= 1
		if x > count/2 {
			gamma |= 1
		}
		// fmt.Println(gamma)
	}
	epsilon := gamma ^ (1<<len(ones) - 1)
	// fmt.Println(gamma, epsilon)
	fmt.Println(gamma * epsilon)
}
