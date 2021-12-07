package main

import (
	"fmt"
	"sort"
	"strings"
)

func Median(ints []int) int {
	sort.Ints(ints)
	// fmt.Println(ints)

	switch l := len(ints); {
	case l == 0:
		return 0
	case l%2 == 0:
		return ints[l/2]
	default:
		return (ints[(l-1)/2] + ints[(l+1)/2]) / 2
	}
}

func main() {
	var line string
	_, err := fmt.Scanln(&line)
	if err != nil {
		panic(err)
	}

	poss := strings.Split(line, ",")
	pos := make([]int, len(poss))
	for i, x := range poss {
		fmt.Sscanf(x, "%d", &pos[i])
	}
	// fmt.Println(pos)
	median := Median(pos)
	// fmt.Println(median)

	var fuel int
	for _, x := range pos {
		if median > x {
			fuel += median - x
		} else {
			fuel += x - median
		}
	}
	fmt.Println(fuel)
}
