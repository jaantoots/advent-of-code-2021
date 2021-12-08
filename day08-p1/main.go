package main

import (
	"fmt"
	"io"
)

func main() {
	var count int
	var ins [10]string
	var outs [4]string
	for {
		_, err := fmt.Scanf("%s %s %s %s %s %s %s %s %s %s | %s %s %s %s", &ins[0], &ins[1], &ins[2], &ins[3], &ins[4], &ins[5], &ins[6], &ins[7], &ins[8], &ins[9], &outs[0], &outs[1], &outs[2], &outs[3])
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(outs)
		for _, out := range outs {
			switch len(out) {
			// 1, 4, 7, 8
			case 2, 4, 3, 7:
				count++
			}
		}
		fmt.Println(count)
	}
	fmt.Println(count)
}
