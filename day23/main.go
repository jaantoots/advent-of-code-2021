package main

import (
	"fmt"
)

const depth = 4

var stepCosts = [4]int{1, 10, 100, 1000}

type State struct {
	rooms   [4][depth]int
	hallway [4*2 + 1 + 2]int
}

func (s *State) Copy() State {
	var ss State
	for i, x := range s.hallway {
		ss.hallway[i] = x
	}
	for i, room := range s.rooms {
		for j, x := range room {
			ss.rooms[i][j] = x
		}
	}
	return ss
}

func FinalState() State {
	var s State
	for i, room := range s.rooms {
		for j := range room {
			s.rooms[i][j] = i + 1
		}
	}
	return s
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (s *State) reachable() map[State]int {
	incr := make(map[State]int)
	// from rooms to hallway
	for i, room := range s.rooms {
		for j, x := range room {
			if x > 0 {
				ref := 2 + 2*i
				// left hallway
				var steps int
				steps = j
				for k := ref; k >= 0; k-- {
					steps++
					if k > 0 && k%2 == 0 {
						continue
					}
					if s.hallway[k] > 0 {
						break // someone in the way
					}
					next := s.Copy()
					next.rooms[i][j] = 0
					next.hallway[k] = x
					incr[next] = steps * stepCosts[x-1]
					//fmt.Printf("%v, %v, %v; cost %v %v\n", i, j, k, incr[next], next)
				}
				// right hallway
				steps = j
				for k := ref; k < len(s.hallway); k++ {
					steps++
					if k < len(s.hallway)-1 && k%2 == 0 {
						continue
					}
					if s.hallway[k] > 0 {
						break // someone in the way
					}
					next := s.Copy()
					next.rooms[i][j] = 0
					next.hallway[k] = x
					incr[next] = steps * stepCosts[x-1]
					//fmt.Printf("%v, %v, %v; cost %v %v\n", i, j, k, incr[next], next)
				}
				break // lower cannot move
			}
		}
	}
	// from hallway to room
Hallway:
	for i, x := range s.hallway {
		if x == 0 {
			continue
		}
		// will only move to its room
		for k := min(2*x, i); k <= max(2*x, i); k++ {
			if k == i {
				continue
			}
			if s.hallway[k] > 0 {
				continue Hallway // way is blocked
			}
		}
		for j := depth - 1; j >= 0; j-- {
			if s.rooms[x-1][j] == x {
				continue // already correctly occupied
			}
			if s.rooms[x-1][j] > 0 {
				break // will not move into room containing invalid
			}
			next := s.Copy()
			next.hallway[i] = 0
			next.rooms[x-1][j] = x
			incr[next] = (abs(i-2*x) + j + 1) * stepCosts[x-1]
			//fmt.Printf("%v, %v, %v; cost %v %v\n", x-1, j, i, incr[next], next)
			break // no empty spaces
		}
	}
	return incr
}

func (s *State) costs() map[State]int {
	costs := make(map[State]int)
	costs[*s] = 0
	stack := make(map[State]bool)
	stack[*s] = true
	for len(stack) > 0 {
		for cur := range stack {
			//fmt.Printf("considering from (cost %v) %v\n", costs[cur], cur)
			for r, c := range cur.reachable() {
				//fmt.Printf("(incr %v) %v\n", c, r)
				prev, ok := costs[r]
				if !ok || costs[cur]+c < prev {
					costs[r] = costs[cur] + c
					stack[r] = true
					//fmt.Printf("added to stack\n")
				}
			}
			delete(stack, cur)
		}
	}
	return costs
}

func (s State) String() string {
	var str string
	str += "\n"
	row := make([]byte, len(s.hallway))
	for i, x := range s.hallway {
		if x > 0 {
			row[i] = byte(x) + 64
		} else {
			row[i] = 46
		}
	}
	str += string(row) + "\n"
	for i := range row {
		row[i] = 32
	}
	for j := 0; j < depth; j++ {
		for i := 0; i < 4; i++ {
			if s.rooms[i][j] > 0 {
				row[2+2*i] = byte(s.rooms[i][j]) + 64
			} else {
				row[2+2*i] = 46
			}
		}
		str += string(row) + "\n"
	}
	return str
}

func main() {
	var initial State
	for i := 0; i < depth; i++ {
		var a, b, c, d int
		_, err := fmt.Scanf("%d %d %d %d\n", &a, &b, &c, &d)
		if err != nil {
			panic(err)
		}
		initial.rooms[0][i] = a
		initial.rooms[1][i] = b
		initial.rooms[2][i] = c
		initial.rooms[3][i] = d
	}
	fmt.Println(initial)
	fmt.Println(FinalState())
	costs := initial.costs()
	//fmt.Println(costs)
	fmt.Println(costs[FinalState()])
}
