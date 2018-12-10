package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day10/inp.txt")
	strdat := strings.Split(string(dat), "\n")
	pts = parseAll(strdat)
	part1()
	//part2 is 10124
}

type Point struct {
	PosX int
	PosY int
	VelX int
	VelY int
}

var pts []*Point

func part1() {
	i := 0
	for i < 10124 {
		for _, pt := range pts {
			pt.PosX += pt.VelX
			pt.PosY += pt.VelY
		}
		i++
	}
	dump()
}

func dump() {
	minX, maxX, minY, maxY := findBB()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if hasDot(x,y) == nil {
				fmt.Print(".")
			} else {
				fmt.Print("x")
			}
		}
		fmt.Println()
	} 
}

func hasDot(x int, y int) *Point{
	for _, pt := range pts{
		if pt.PosX == x && pt.PosY == y{
			return pt
		}
	}
	return nil
}

func findBB() (int, int, int, int) {
	minX, maxX, minY, maxY := 1000000, 0, 10000000, 0
	for _, pt := range pts{
		if minX > pt.PosX { minX = pt.PosX }
		if maxX < pt.PosX { maxX = pt.PosX }
		if minY > pt.PosY { minY = pt.PosY }
		if maxY < pt.PosX { maxY = pt.PosY }
	}
	return minX, maxX, minY, maxY
}

func parseAll(s []string) []*Point{
	ret:=make([]*Point, len(s)-1)
	for i:=0; i < len(s)-1;i++{
		ret[i] = parse(s[i])
	}
	return ret
}


func parse(s string) *Point {
	ret := Point{}
	fmt.Sscanf(s, "position=<%d,  %d> velocity=<%d, %d>", &ret.PosX, &ret.PosY, &ret.VelX, &ret.VelY)
	return &ret
}
