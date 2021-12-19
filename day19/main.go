package main

import (
	"fmt"
)

const RequiredMatches = 12

type Point struct {
	x, y, z int
}

func (a Point) Subtract(b Point) Point {
	return Point{a.x - b.x, a.y - b.y, a.z - b.z}
}

func (a Point) Add(b Point) Point {
	return Point{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (p Point) cycle() Point {
	return Point{p.y, p.z, p.x}
}

func (p Point) rotateXY() Point {
	return Point{-p.y, p.x, p.z}
}

func (p Point) rotateXZ() Point {
	return Point{-p.x, p.y, -p.z}
}

type Scanner []Point

type Matches map[int]int

func BuildMap(scanners *map[int]*Scanner, start int) *Scanner {
	s := (*scanners)[start]
	delete(*scanners, start)
	//fmt.Println(start, *s)
	big := s
	for i, o := range *scanners {
		matches, t := s.Match(o)
		if matches != nil {
			subMap := BuildMap(scanners, i).Transform(t)
			m := big.match(subMap)
			big = big.extend(m, subMap)
			fmt.Printf("add to %d (%d beacons) from %d (%d beacons): %d beacons\n", start, len(*big), i, len(*subMap), len(*big))
		}
	}
	return big
}

func (s *Scanner) extend(matches Matches, o *Scanner) *Scanner {
	big := make(Scanner, len(*s), len(*s)+len(*o)-len(matches))
	for i, p := range *s {
		big[i] = p
	}
	for j, p := range *o {
		_, ok := matches[j]
		if ok {
			continue
		}
		big = append(big, p)
	}
	return &big
}

type Transform struct {
	c, r, s int
	delta   Point
}

func (s *Scanner) Transform(t Transform) *Scanner {
	return s.cycle(t.c).rotateXY(t.r).rotateXZ(t.s).Translate(t.delta)
}

func (s *Scanner) Equals(o *Scanner) bool {
	if len(*s) != len(*o) {
		return false
	}
	for i, p := range *s {
		if p != (*o)[i] {
			return false
		}
	}
	return true
}

func (s *Scanner) Match(o *Scanner) (Matches, Transform) {
	// consider every relative orientation
	//fmt.Printf("\n\nconsider rotations of: %v\n", *o)
	for cyc := 0; cyc < 3; cyc++ {
		c := o.cycle(cyc)
		for rot := 0; rot < 4; rot++ {
			r := c.rotateXY(rot)
			for sgn := 0; sgn < 2; sgn++ {
				//n := s.cycle(cyc).rotateXY(rot).rotateXZ(sgn)
				n := r.rotateXZ(sgn)
				//fmt.Printf("(%d, %d, %d): %v\n", cyc, rot, sgn, *n)
				matches, translation := s.matchTranslate(n)
				if matches != nil {
					return matches, Transform{cyc, rot, sgn, translation}
				}
			}
		}
	}
	return nil, Transform{}
}

func (s *Scanner) cycle(n int) *Scanner {
	if n == 0 {
		return s
	}
	t := make(Scanner, len(*s))
	for i, p := range *s {
		t[i] = p
		for j := 0; j < n; j++ {
			t[i] = t[i].cycle()
		}
	}
	return &t
}

func (s *Scanner) rotateXY(n int) *Scanner {
	if n == 0 {
		return s
	}
	t := make(Scanner, len(*s))
	for i, p := range *s {
		t[i] = p
		for j := 0; j < n; j++ {
			t[i] = t[i].rotateXY()
		}
	}
	return &t
}

func (s *Scanner) rotateXZ(n int) *Scanner {
	if n == 0 {
		return s
	}
	t := make(Scanner, len(*s))
	for i, p := range *s {
		t[i] = p.rotateXZ()
	}
	return &t
}

func (s *Scanner) matchTranslate(o *Scanner) (Matches, Point) {
	// consider translations that map at least one beacon from both scanners
	for _, first := range *s {
		for _, second := range *o {
			translation := first.Subtract(second)
			t := o.Translate(translation)
			matches := s.match(t)
			//fmt.Printf("%v, %v, %v: %v\n", translation, first, second, len(matches))
			// should have only one
			if len(matches) >= RequiredMatches {
				return matches, translation
			}
		}
	}
	return nil, Point{}
}

func (s *Scanner) Translate(delta Point) *Scanner {
	t := make(Scanner, len(*s))
	for i, p := range *s {
		t[i] = p.Add(delta)
	}
	return &t
}

func (s *Scanner) match(o *Scanner) Matches {
	// find which beacons have matching positions
	matches := make(Matches, 16)
	for i, first := range *s {
		for j, second := range *o {
			if first == second {
				matches[j] = i
			}
		}
	}
	return matches
}

func readScanner() (int, *Scanner) {
	points := make(Scanner, 0, 64)
	var num int
	_, err := fmt.Scanf("--- scanner %d ---\n", &num)
	if err != nil {
		return -1, &points
	}
	//fmt.Println(num)
	var x, y, z int
	for {
		_, err := fmt.Scanf("%d,%d,%d\n", &x, &y, &z)
		if err != nil {
			break
		}
		points = append(points, Point{x, y, z})
	}
	//fmt.Println(num, points)
	return num, &points
}

func main() {
	scanners := make(map[int]*Scanner, 64)
	for {
		i, s := readScanner()
		if len(*s) == 0 {
			break
		}
		scanners[i] = s
	}
	//fmt.Println(scanners)
	//for i, s := range scanners {
	//	//fmt.Println(i, *s)
	//	for j, o := range scanners {
	//		if i >= j {
	//			continue
	//		}
	//		matches, _ := s.Match(o)
	//		if matches != nil {
	//			fmt.Println(i, j, matches)
	//		}
	//	}
	//}
	big := BuildMap(&scanners, 0)
	//fmt.Println(*big)
	fmt.Println(len(*big))
}
