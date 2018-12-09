package main

import (
	"fmt"
	"container/ring"
)

func main() {
	strdat := "411 players; last marble is worth 71058 points"
	a, c := 0, 0
	fmt.Sscanf(strdat, "%d players; last marble is worth %d points", &a, &c)
	part1(a, c)
	part2(a, c)
}


func part1(players int, marble int) {
	scores := make([]int, players)
	circle := ring.New(1)
	lastPut := 0
	cur := circle
	turn := 0
	for lastPut != marble {
		now := lastPut + 1
		if now % 23 == 0 {
			scores[turn] += now
			for i:=0;i<8;i++{
				cur = cur.Prev()
			}
			scores[turn] += cur.Unlink(1).Value.(int)
			cur = cur.Next()
		} else {
			newCur := newRing(now)
			cur.Next().Link(newCur)
			cur = newCur
		}
		lastPut = now
		turn = (turn+1)%players
	}
	max := 0
	for _, x := range scores {
		if max < x {
			max = x
		}
	}
	fmt.Println(max)
}

func newRing(x int) *ring.Ring {
	ret := ring.New(1)
	ret.Value = x
	return ret
}

func part2(players int, marble int){
	part1(players, marble*100)
}
