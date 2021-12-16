package main

import (
	"fmt"
	"io"
)

func send(s *string, c chan<- uint64) {
	if len(*s)%16 != 0 {
		*s += "0000000000000000"[len(*s)%16:]
	}
	var buffer uint64
	for i := 0; i < len(*s); i += 16 {
		_, err := fmt.Sscanf((*s)[i:i+16], "%X", &buffer)
		if err != nil {
			panic(err)
		}
		c <- buffer
	}
	close(c)
}

type Receiver struct {
	channel    <-chan uint64
	buffer     uint64
	idx        int
	versionSum uint64
}

func NewReceiver(c <-chan uint64) Receiver {
	return Receiver{c, 0, 64, 0}
}

func (r *Receiver) recv(n int) (uint64, bool) {
	var mask uint64
	mask--
	mask >>= r.idx
	val := r.buffer & mask
	if 64-r.idx >= n {
		val >>= 64 - r.idx - n
		r.idx += n
		return val, true
	}
	n -= 64 - r.idx
	var ok bool
	r.buffer, ok = <-r.channel
	r.idx = 0
	if !ok {
		return val, false
	}
	//fmt.Printf("%08X\n", buf)
	//fmt.Printf("buffer: %064b\n", r.buffer)
	extra, ok := r.recv(n)
	return val<<n | extra, ok
}

func (r *Receiver) handleLiteral() (uint64, uint64) {
	var length uint64
	var val uint64
	for i := 0; ; i++ {
		length += 5
		group, ok := r.recv(5)
		if !ok {
			panic("expected literal to continue")
		}
		if i*4 >= 64 {
			panic("uint64 overflow")
		}
		val = val<<4 | (0xF & group)
		if group>>4 == 0 {
			break
		}
	}
	//fmt.Printf("literal: %064b\n", val)
	return length, val
}

func (r *Receiver) handleBits(f func(uint64)) uint64 {
	var length uint64 = 15
	subLength, ok := r.recv(15)
	if !ok {
		panic("expected length of sub-packet bits")
	}
	//fmt.Printf("sub-packets bits: %d\n", subLength)
	for length < subLength+15 {
		l, val := r.handlePacket()
		length += l
		f(val)
	}
	if length != subLength+15 {
		panic("consumed too many bits")
	}
	return length
}

func (r *Receiver) handleNum(f func(uint64)) uint64 {
	var length uint64 = 11
	subNum, ok := r.recv(11)
	if !ok {
		panic("expected number of sub-packets")
	}
	//fmt.Printf("sub-packets num: %d\n", subNum)
	var i uint64
	for i = 0; i < subNum; i++ {
		l, val := r.handlePacket()
		length += l
		f(val)
	}
	return length
}

func (r *Receiver) handleOperator(opType uint64) (uint64, uint64) {
	//fmt.Printf("operation %d\n", opType)
	var length uint64 = 1
	lengthTypeId, ok := r.recv(1)
	if !ok {
		panic("expected packet type")
	}
	var val uint64
	// Initialize
	switch opType {
	case 1:
		val = 1
	case 2:
		val--
	}
	var prev uint64
	f := func(x uint64) {
		switch opType {
		case 0:
			if x>>63 > 0 || val>>63 > 0 {
				fmt.Printf("%064b %064b\n", x, val)
				panic("overflow danger")
			}
			val += x
		case 1:
			if x>>32 > 0 || val>>32 > 0 {
				fmt.Printf("%064b %064b\n", x, val)
				panic("overflow danger")
			}
			val *= x
		case 2:
			if x < val {
				val = x
			}
		case 3:
			if x > val {
				val = x
			}
		case 5:
			if prev > x {
				val = 1
			} else {
				val = 0
			}
			prev = x
		case 6:
			if prev < x {
				val = 1
			} else {
				val = 0
			}
			prev = x
		case 7:
			if prev == x {
				val = 1
			} else {
				val = 0
			}
			prev = x
		}
	}
	switch lengthTypeId {
	case 0:
		length += r.handleBits(f)
	case 1:
		length += r.handleNum(f)
	}
	//fmt.Printf("operation %d: %064b\n", opType, val)
	return length, val
}

func (r *Receiver) handlePacket() (uint64, uint64) {
	var length uint64 = 6
	version, ok := r.recv(3)
	if !ok {
		return 0, 0
	}
	r.versionSum += version
	//fmt.Printf("version: %d\n", version)
	typeId, ok := r.recv(3)
	if !ok {
		panic("expected packet type")
	}
	var l, val uint64
	switch typeId {
	case 4:
		l, val = r.handleLiteral()
		length += l
	default:
		l, val = r.handleOperator(typeId)
		length += l
	}
	return length, val
}

func handle(s string) {
	c := make(chan uint64)
	go send(&s, c)

	r := NewReceiver(c)
	_, value := r.handlePacket()
	fmt.Println(r.versionSum)
	fmt.Println(value)
	//for {
	//	v, ok := r.recv(16)
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("%08X\n", v)
	//	fmt.Printf("%064b\n", v)
	//}
}

func main() {
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		handle(line)
	}
}
