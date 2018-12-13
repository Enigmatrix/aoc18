package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day6/inp.txt")
	strdat := strings.Split(string(dat), "\n")
	coords := parseCoords(strdat)
	part1(coords)
	part2(coords)
}

const MAX = 400

type Coord struct{
	X int
	Y int
	Label rune
	DistMatrix *[MAX][MAX]int
}

type Match struct {
	CoordIndex int
	Tie bool
	Initialized bool
}

func abs(a int) int{
	if a < 0 {
		return -a
	}
	return a
}

func part1(coords []Coord) {

	var match [MAX][MAX]Match
	for x:=0; x<MAX;x++{
		for y:=0; y<MAX;y++{
			for i, c := range coords {
				if !match[x][y].Initialized || coords[match[x][y].CoordIndex].DistMatrix[x][y] > c.DistMatrix[x][y] {
					match[x][y].CoordIndex = i
					match[x][y].Tie = false
					match[x][y].Initialized = true
				} else if coords[match[x][y].CoordIndex].DistMatrix[x][y] == c.DistMatrix[x][y] {
					//fmt.Println("Tie ", x, ", ", y, "(", *match[x][y].Coord, ", ", c, ")")
					match[x][y].Tie = true
				}
			}
		}
	}

	exceptSet := make(map[int]struct{})
	for x:=0; x<MAX;x++{
		exceptSet[match[x][0].CoordIndex] = struct{}{}
		exceptSet[match[0][x].CoordIndex] = struct{}{}
		exceptSet[match[x][MAX-1].CoordIndex] = struct{}{}
		exceptSet[match[MAX-1][x].CoordIndex] = struct{}{}
	}

	/*
	for x:=0; x<MAX;x++{
		for y:=0; y<MAX;y++{
			m:=match[x][y]
			if m.Tie {
				fmt.Print(".")
			} else {
				fmt.Print(m.Coord.Label)
			}
		}
		fmt.Println()
	}*/

	nuni := make(map[Coord]int)
	for x:=0; x<MAX;x++{
		for y:=0; y<MAX;y++{
			m:=match[x][y]
			c := m.CoordIndex
			if _, prs := exceptSet[c]; prs {
				continue
			}
			if !m.Tie {
				_, prs := nuni[coords[c]]
				if !prs {
					nuni[coords[c]] = 0
				}
				nuni[coords[c]]++
			}
		}
	}

	max := 0
	for _, v := range nuni {
		if max < v {
			max = v
		}
	}
	fmt.Println(max)

}

func part2(coords []Coord) {

	var dist [MAX][MAX]int
	for x:=0; x<MAX;x++{
		for y:=0; y<MAX;y++{
			for _, c := range coords {
				dist[x][y] += c.DistMatrix[x][y]
			}
		}
	}

	m := 0
	for x:=0; x<MAX;x++{
		for y:=0; y<MAX;y++{
			if dist[x][y] < 10000 {
				m++
			}
		}
	}
	fmt.Println(m)

	
}

func parseCoords(s []string) []Coord{
	ret := make([]Coord, len(s)-1)
	for i:= 0; i < len(s)-1; i++{
		ret[i] = parseCoord(s[i])
		ret[i].Label = 'a'+rune(i)
	}
	return ret
}

func parseCoord(s string) Coord{
	xy := strings.Split(s, ", ")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	
	c := Coord{
		X: x,
		Y: y,
		DistMatrix: &[MAX][MAX]int{},
	}

	for x:=0; x<MAX;x++{
		for y:=0; y<MAX;y++{
			c.DistMatrix[x][y] = abs(x-c.X)+abs(y-c.Y)
		}
	}

	return c
}
