package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	_, err := fmt.Scanln(&line)
	if err != nil {
		panic(err)
	}

	timers := strings.Split(line, ",")
	var counts [9]uint
	var num int
	for _, x := range timers {
		fmt.Sscanf(x, "%d", &num)
		counts[num]++
	}
	// fmt.Println(counts)

	day := 0
	for day < 80 {
		day++
		spawn := counts[0]
		for i := 0; i < 8; i++ {
			counts[i] = counts[i+1]
		}
		counts[6] += spawn
		counts[8] = spawn
		// fmt.Println(day, counts)
	}

	var sum uint
	for _, x := range counts {
		sum += x
	}
	fmt.Println(sum)
}
