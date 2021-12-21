package main

import (
	"fmt"
	"os"
)

type Die int

func (d *Die) Roll() int {
	defer func() { (*d)++ }()
	return int(*d)%100 + 1
}

type Player struct {
	pos, score int
}

func (p *Player) Turn(d *Die) {
	p.pos += d.Roll() + d.Roll() + d.Roll()
	p.pos %= 10
	p.score += p.pos + 1
}

func playFrom(p1, p2 int) (int, int) {
	var die Die
	pl1, pl2 := Player{p1 - 1, 0}, Player{p2 - 1, 0}
	for {
		pl1.Turn(&die)
		if pl1.score >= 1000 {
			fmt.Printf("Player 1 has won: %d\n", pl1.score)
			return pl2.score, int(die)
		}
		pl2.Turn(&die)
		if pl2.score >= 1000 {
			fmt.Printf("Player 2 has won: %d\n", pl2.score)
			return pl1.score, int(die)
		}
	}
}

func main() {
	var p1, p2 int
	fmt.Sscanf(os.Args[1], "%d", &p1)
	fmt.Sscanf(os.Args[2], "%d", &p2)
	fmt.Println(p1, p2)
	p, d := playFrom(p1, p2)
	fmt.Println(p, d, p*d)
}
