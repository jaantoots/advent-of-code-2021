package main

import (
	"fmt"
)

type Target struct {
	x_min, x_max, y_min, y_max int
}

func (t *Target) whenHorizontalTarget(v int) []int {
	var x int
	t_in := make([]int, 0, 4)
	for i := 0; x <= t.x_max; i++ {
		if x <= t.x_max && x >= t.x_min {
			t_in = append(t_in, i)
			if v == 0 {
				// Will stay in target
				t_in = append(t_in, -1)
			}
		}
		if v == 0 {
			break
		}
		x += v
		v--
	}
	return t_in
}

func (t *Target) whenVerticalTarget(u int) []int {
	var y int
	t_in := make([]int, 0, 4)
	for i := 0; y >= t.y_min; i++ {
		if y <= t.y_max && y >= t.y_min {
			t_in = append(t_in, i)
		}
		y += u
		u--
	}
	return t_in
}

func (t *Target) findHighestTrajectory() int {
	if t.y_max > 0 {
		panic("y_max > 0")
	}
	if t.x_min < 0 {
		panic("x_min < 0")
	}
	// All horizontally legal timesteps
	t_att := -1
	v_max := t.x_max
	t_leg := make(map[int]bool, 256)
	for v := 0; v <= v_max; v++ {
		t_in := t.whenHorizontalTarget(v)
		var last int
		for _, tt := range t_in {
			if tt == -1 {
				if last < t_att || t_att < 0 {
					t_att = last
				}
			} else {
				t_leg[tt] = true
			}
			last = tt
		}
		//fmt.Println(v, t_in, t_att, t_leg)
	}
	fmt.Println(t_att, t_leg)
	// Otherwise will always skip over
	u_max := ((-8*t.y_min-1)/2 - 1) / 2
	fmt.Println(u_max)
	var u_att int
	for u := 0; u <= u_max; u++ {
		t_in := t.whenVerticalTarget(u)
		//fmt.Println(u, t_in)
		for _, tt := range t_in {
			if (tt >= t_att && t_att >= 0) || t_leg[tt] {
				//fmt.Println(u, tt)
				u_att = u
				break
			}
		}
	}
	return u_att * (u_att + 1) / 2
}

func main() {
	var x_min, x_max, y_min, y_max int
	for {
		_, err := fmt.Scanf("target area: x=%d..%d, y=%d..%d", &x_min, &x_max, &y_min, &y_max)
		if err != nil {
			break
		}
		target := Target{x_min, x_max, y_min, y_max}
		fmt.Println(target)
		fmt.Println(target.findHighestTrajectory())
	}
}
