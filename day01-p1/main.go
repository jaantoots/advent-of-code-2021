package main

import (
	"fmt"
	"io"
)

func main() {
	var depth, prev, count int
	for {
		_, err := fmt.Scanln(&depth)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if prev > 0 && depth > prev {
			count += 1
		}
		// fmt.Printf("%d %d %t %d\n", prev, depth, depth > prev, count)
		prev = depth
	}
	fmt.Println(count)
}
