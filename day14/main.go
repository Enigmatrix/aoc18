package main

import (
	"fmt"
	"strconv"
	"container/ring"
)
var inp int
func main() {
	inp, _ = strconv.Atoi("047801")
	part1()
	part2()
}

func part1() {
	//inp = 
	current1 := newRing(3);
	current2 := newRing(7)
	front := current1
	current1.Link(current2)
	for current1.Len() < inp+10{
		digits := current1.Value.(int) + current2.Value.(int)
		if digits >= 10 {
			g1 := newRing(digits / 10)
			g2 := newRing(digits % 10)
			g1.Link(g2)
			front.Prev().Link(g1)
		} else {
			g := newRing(digits)
			front.Prev().Link(g)
		}
		c1, c2 := current1.Value.(int), current2.Value.(int)
		for i:=0; i < c1+1; i++{
			current1 = current1.Next()
		}
		for i:=0; i < c2+1; i++{
			current2 = current2.Next()
		}
	}


	for i := 0; i < inp; i++ {
		front = front.Next()
	}
	for i := 0; i < 10; i++ {
		fmt.Print(front.Value)
		front = front.Next()
	}
	fmt.Println()
}

func newRing(v int) *ring.Ring{
	r:= ring.New(1)
	r.Value = v
	return r
}

func match(d int, c int) int{
	ih := "047801"
	if s, _ := strconv.Atoi(string(ih[c+1])); s == d {
		if c+1 == 5 {
			return -2
		}
		return c+1
	}
	return -1
}

func part2() {

	current1 := newRing(3);
	current2 := newRing(7)
	front := current1
	curDigitMatch := -1
	current1.Link(current2)
	for{
		digits := current1.Value.(int) + current2.Value.(int)
		if digits >= 10 {
			c1 := digits/10
			c2 := digits%10
			g1 := newRing(c1)
			g2 := newRing(c2)
			curDigitMatch = match(c1, curDigitMatch)
			if curDigitMatch == -2 {
				fmt.Println(current1.Len()-5)
				return
			}
			curDigitMatch = match(c2, curDigitMatch)
			if curDigitMatch == -2 {
				fmt.Println(current1.Len()-4)
				return
			}
			g1.Link(g2)
			front.Prev().Link(g1)
		} else {
			curDigitMatch = match(digits, curDigitMatch)
			if curDigitMatch == -2 {
				fmt.Println(current1.Len()-5)
				return
			}
			g := newRing(digits)
			front.Prev().Link(g)
		}
		c1, c2 := current1.Value.(int), current2.Value.(int)
		for i:=0; i < c1+1; i++{
			current1 = current1.Next()
		}
		for i:=0; i < c2+1; i++{
			current2 = current2.Next()
		}
	}
}

