package main

import (
	"fmt"
	"os"
)

type Player struct {
	pos, score int
}

func (p *Player) Move(n int) Player {
	pp := Player{p.pos, p.score}
	pp.pos += n
	pp.pos %= 10
	pp.score += pp.pos + 1
	return pp
}

type PlayerUniverses map[Player]int

func (p *PlayerUniverses) turn() PlayerUniverses {
	next := make(PlayerUniverses, 2*len(*p))
	for player, count := range *p {
		// [1 1 1] [1 1 1] -> [1 2 3 2 1]
		// [1 2 3 2 1] [1 1 1] -> [1 3 6 7 6 3 1]
		next[player.Move(3)] += count
		next[player.Move(4)] += 3 * count
		next[player.Move(5)] += 6 * count
		next[player.Move(6)] += 7 * count
		next[player.Move(7)] += 6 * count
		next[player.Move(8)] += 3 * count
		next[player.Move(9)] += count
	}
	return next
}

func (p *PlayerUniverses) total() int {
	var t int
	for _, count := range *p {
		t += count
	}
	return t
}

func (p *PlayerUniverses) countWins() int {
	var c int
	for player, count := range *p {
		if player.score >= 21 {
			c += count
			delete(*p, player)
		}
	}
	return c
}

func main() {
	var p1, p2 int
	fmt.Sscanf(os.Args[1], "%d", &p1)
	fmt.Sscanf(os.Args[2], "%d", &p2)
	fmt.Println(p1, p2)

	p1u := PlayerUniverses{Player{p1 - 1, 0}: 1}
	p2u := PlayerUniverses{Player{p2 - 1, 0}: 1}
	//fmt.Println(p1u, p2u)
	var p1w, p2w int
	for i := 1; p1u.total() > 0 && p2u.total() > 0; i++ {
		p1u = p1u.turn()
		p1w += p1u.countWins() * p2u.total()
		fmt.Printf("After turn %d.1: player 1 wins %d, player 2 wins %d, player 1 left %d, player 2 left %d\n", i, p1w, p2w, p1u.total(), p2u.total())
		p2u = p2u.turn()
		p2w += p2u.countWins() * p1u.total()
		fmt.Printf("After turn %d.2: player 1 wins %d, player 2 wins %d, player 1 left %d, player 2 left %d\n", i, p1w, p2w, p1u.total(), p2u.total())
	}
	fmt.Println(p1w > p2w, p1w, p2w)
}
