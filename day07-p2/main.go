package main

import (
	"fmt"
	"strings"
)

func Mean(ints []int) int {
	var sum int
	for _, x := range ints {
		sum += x
	}
	fmt.Println(sum)
	return sum / len(ints)
}

func Max(ints []int) int {
	var max int
	for _, x := range ints {
		if x > max {
			max = x
		}
	}
	return max
}

func Min(ints []int) int {
	var min int
	if len(ints) != 0 {
		min = ints[0]
	}
	for _, x := range ints {
		if x < min {
			min = x
		}
	}
	return min
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

	costs := make([]int, Max(pos)+1)
	for _, x := range pos {
		cost := 0
		for i := x; i >= 0; i-- {
			cost += (x - i)
			costs[i] += cost
		}
		cost = 0
		for i := x; i < len(costs); i++ {
			cost += (i - x)
			costs[i] += cost
		}
	}
	// fmt.Println(costs)
	fmt.Println(Min(costs))
}
