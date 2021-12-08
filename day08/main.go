package main

import (
	"fmt"
	"io"
	"strings"
)

func contains(x string, y string) bool {
	for _, r := range y {
		if !strings.ContainsRune(x, r) {
			return false
		}
	}
	return true
}

func main() {
	var count int
	var sum int
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

		var key [10]string
		for _, in := range ins {
			switch len(in) {
			case 2:
				key[1] = in
			case 4:
				key[4] = in
			case 3:
				key[7] = in
			case 7:
				key[8] = in
			}
		}
		// len 6: 0, 6, 9
		for _, in := range ins {
			if len(in) != 6 {
				continue
			}
			switch {
			case contains(in, key[4]):
				key[9] = in
			case contains(in, key[1]):
				key[0] = in
			default:
				key[6] = in
			}
		}
		// len 5: 2, 3, 5
		for _, in := range ins {
			if len(in) != 5 {
				continue
			}
			switch {
			case contains(in, key[1]):
				key[3] = in
			case contains(key[6], in):
				key[5] = in
			default:
				key[2] = in
			}
		}
		// fmt.Println(key)

		var num int
		for _, out := range outs {
			switch len(out) {
			// 1, 4, 7, 8
			case 2, 4, 3, 7:
				count++
			}

			num *= 10
			for i, k := range key {
				if contains(out, k) && contains(k, out) {
					num += i
				}
			}
		}
		sum += num
		// fmt.Println(num)
	}
	fmt.Println(count)
	fmt.Println(sum)
}
