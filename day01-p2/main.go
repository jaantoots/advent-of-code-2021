package main

import (
	"fmt"
	"io"
)

const window = 3

func main() {
	var buffer [window]int
	var cur, idx, count int
	for {
		_, err := fmt.Scanln(&cur)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if buffer[idx] > 0 && cur > buffer[idx] {
			count += 1
		}
		// fmt.Printf("%v %d %d %d\n", buffer, idx, cur, count)
		buffer[idx] = cur
		idx = (idx + 1) % window
	}
	fmt.Println(count)
}
