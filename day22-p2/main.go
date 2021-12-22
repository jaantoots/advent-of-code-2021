package main

import (
	"container/list"
	"fmt"
	"io"
)

type Segment struct {
	min, max int
}

type Cuboid struct {
	state bool
	axes  [3]Segment
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func (c *Cuboid) valid() bool {
	for _, seg := range c.axes {
		if seg.min > seg.max {
			return false
		}
	}
	return true
}

func (c *Cuboid) Copy() *Cuboid {
	axes := [3]Segment{}
	for i, seg := range c.axes {
		axes[i] = Segment{seg.min, seg.max}
	}
	return &Cuboid{c.state, axes}
}

func (c *Cuboid) Volume() int {
	vol := 1
	for _, seg := range c.axes {
		vol *= seg.max - seg.min + 1
	}
	return vol
}

func (a *Cuboid) Intersect(b *Cuboid) *Cuboid {
	axes := [3]Segment{}
	for i := range axes {
		axes[i] = Segment{max(a.axes[i].min, b.axes[i].min), min(a.axes[i].max, b.axes[i].max)}
	}
	c := Cuboid{a.state, axes}
	if c.valid() {
		return &c
	}
	return nil
}

func (c *Cuboid) split(ax, x int) (*Cuboid, *Cuboid) {
	left, right := c.Copy(), c.Copy()
	left.axes[ax].max = x - 1
	right.axes[ax].min = x
	return left, right
}

func (a *Cuboid) fragment(sub *Cuboid) []*Cuboid {
	all := []*Cuboid{}
	next := a
	var left, middle, right *Cuboid
	for i, seg := range sub.axes {
		left, middle = next.split(i, seg.min)
		middle, right = middle.split(i, seg.max+1)
		if left.valid() {
			all = append(all, left)
		}
		if right.valid() {
			all = append(all, right)
		}
		next = middle
	}
	all = append(all, next)
	return all
}

func (a *Cuboid) Combine(b *Cuboid) ([]*Cuboid, []*Cuboid) {
	inter := a.Intersect(b)
	if inter == nil {
		//return []*Cuboid{a}, []*Cuboid{b}
		return nil, nil
	}
	//fmt.Printf("cuboids %v and %v intersect\n", *a, *b)
	var frags, extra []*Cuboid
	for _, f := range a.fragment(inter) {
		if *f == *inter {
			f.state = b.state
			//fmt.Printf("    inter %v\n", *f)
		}
		frags = append(frags, f)
		//fmt.Printf("  fragment %v\n", *f)
	}
	inter.state = b.state
	for _, f := range b.fragment(inter) {
		if *f == *inter {
			//fmt.Printf("    inter %v\n", *f)
			continue
		}
		extra = append(extra, f)
		//fmt.Printf("  extra %v\n", *f)
	}
	return frags, extra
}

func addCuboid(cuboids *list.List, cuboid *Cuboid) {
	stack := list.New()
	stack.PushBack(cuboid)
	var a, b *Cuboid
	for e := stack.Front(); e != nil; e = e.Next() {
		b = e.Value.(*Cuboid)
		var match bool
		for d := cuboids.Front(); d != nil; d = d.Next() {
			a = d.Value.(*Cuboid)
			frags, extra := a.Combine(b)
			for _, f := range frags {
				cuboids.PushBack(f)
			}
			for _, f := range extra {
				stack.PushBack(f)
			}
			if frags != nil {
				cuboids.Remove(d)
				match = true
				break
			}
		}
		if !match && b.state {
			cuboids.PushBack(b)
		}
	}
}

func onCount(cuboids *list.List) int {
	var count int
	var c *Cuboid
	for e := cuboids.Front(); e != nil; e = e.Next() {
		c = e.Value.(*Cuboid)
		if c.state {
			count += c.Volume()
		}
	}
	return count
}

func main() {
	var state string
	var x0, x1, y0, y1, z0, z1 int
	cuboids := list.New()
	for {
		_, err := fmt.Scanf("%s x=%d..%d,y=%d..%d,z=%d..%d", &state, &x0, &x1, &y0, &y1, &z0, &z1)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		cuboid := Cuboid{state == "on", [3]Segment{Segment{x0, x1}, Segment{y0, y1}, Segment{z0, z1}}}
		fmt.Println(cuboid)
		addCuboid(cuboids, &cuboid)
		//fmt.Println(onCount(cuboids))
	}
	fmt.Println(onCount(cuboids))
}
