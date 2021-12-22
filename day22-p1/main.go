package main

import (
	"fmt"
	"io"
)

func main() {
	var state string
	var x0, x1, y0, y1, z0, z1 int
	var m [101][101][101]bool
	for {
		_, err := fmt.Scanf("%s x=%d..%d,y=%d..%d,z=%d..%d", &state, &x0, &x1, &y0, &y1, &z0, &z1)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(x0, x1, y0, y1, z0, z1)
		for x := x0; x <= x1; x++ {
			if x < -50 || x > 50 {
				continue
			}
			for y := y0; y <= y1; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for z := z0; z <= z1; z++ {
					if z < -50 || z > 50 {
						continue
					}
					m[50+x][50+y][50+z] = state == "on"
				}
			}
		}
	}
	var count int
	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100; y++ {
			for z := 0; z <= 100; z++ {
				if m[x][y][z] {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
