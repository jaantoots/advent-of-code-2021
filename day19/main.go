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

func (p Point) Cycle(n int) Point {
	pp := p
	for j := 0; j < n; j++ {
		pp = pp.cycle()
	}
	return pp
}

func (p Point) RotateXY(n int) Point {
	pp := p
	for j := 0; j < n; j++ {
		pp = pp.rotateXY()
	}
	return pp
}

func (p Point) RotateXZ(n int) Point {
	pp := p
	for j := 0; j < n; j++ {
		pp = pp.rotateXZ()
	}
	return pp
}

type Transform struct {
	c, r, s int
	delta   Point
}

func (p Point) Transform(t Transform) Point {
	return p.Cycle(t.c).RotateXY(t.r).RotateXZ(t.s).Add(t.delta)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (a Point) L1Dist(b Point) int {
	delta := a.Subtract(b)
	return abs(delta.x) + abs(delta.y) + abs(delta.z)
}

func MaxL1(pos map[int]Point) (int, int, int) {
	var max_i, max_j, max int
	for i, p := range pos {
		for j, r := range pos {
			if i >= j {
				continue
			}
			dist := p.L1Dist(r)
			if dist > max {
				max_i, max_j, max = i, j, dist
			}
		}
	}
	return max_i, max_j, max
}

type Scanner []Point

type Matches map[int]int

func BuildMap(scanners *map[int]*Scanner, start int) (*Scanner, map[int]Point) {
	positions := make(map[int]Point, len(*scanners))
	positions[start] = Point{0, 0, 0}

	s := (*scanners)[start]
	delete(*scanners, start)
	//fmt.Println(start, *s)
	big := s
	for i, o := range *scanners {
		matches, t := s.Match(o)
		if matches != nil {
			subMap, subPositions := BuildMap(scanners, i)
			subMap = subMap.Transform(t)
			for j, p := range subPositions {
				positions[j] = p.Transform(t)
			}
			m := big.match(subMap)
			big = big.extend(m, subMap)
			fmt.Printf("add to %d (%d beacons) from %d (%d beacons): %d beacons\n", start, len(*big), i, len(*subMap), len(*big))
		}
	}
	return big, positions
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

func (s *Scanner) Transform(t Transform) *Scanner {
	r := make(Scanner, len(*s))
	for i, p := range *s {
		r[i] = p.Transform(t)
	}
	return &r
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
		for rot := 0; rot < 4; rot++ {
			for sgn := 0; sgn < 2; sgn++ {
				//n := s.cycle(cyc).rotateXY(rot).rotateXZ(sgn)
				r := Transform{cyc, rot, sgn, Point{0, 0, 0}}
				n := o.Transform(r)
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
	big, positions := BuildMap(&scanners, 0)
	//fmt.Println(*big)
	fmt.Println(len(*big))
	fmt.Println(MaxL1(positions))
}
