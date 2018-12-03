package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Claim struct {
	Id string
	Left int
	Top int
	Width int
	Height int
}

func main() {
	dat, _ := ioutil.ReadFile("day3/inp.txt")
	strdat := strings.Split(string(dat), "\n")
	claims := parseClaims(strdat)
	part1(claims)
	part2(claims)
}

func part1(claims []Claim) {
	var fabric  [1000][1000]int
	for _, c := range claims {
		
		for x := c.Left; x < c.Width+c.Left; x++ {
			for y := c.Top; y < c.Height+c.Top; y++ {
				fabric[x][y]++
			}
		}

	}

	total := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if fabric[x][y] >= 2 {
				total++
			}
		}
	}
	fmt.Println(total)

}

func part2(claims []Claim) {
	var fabric  [1000][1000][]string
	for _, c := range claims {
		
		for x := c.Left; x < c.Width+c.Left; x++ {
			for y := c.Top; y < c.Height+c.Top; y++ {
				fabric[x][y] = append(fabric[x][y], c.Id)
			}
		}

	}


	for _, c := range claims {

		good := true
		Claim:
		for x := c.Left; x < c.Width+c.Left; x++ {
			for y := c.Top; y < c.Height+c.Top; y++ {
				if len(fabric[x][y]) > 1 {
					good = false
					break Claim
				}
			}
		}
		if good {
			fmt.Print(c.Id)
			return
		}
	
	}
}

func parseClaims(s []string) []Claim {
	ret := make([]Claim, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		ret[i] = parseClaim(s[i])
	}
	return ret
}

func parseClaim(s string) Claim {
	idAndRest := strings.Split(s, " @ ")
	id := idAndRest[0][1: len(idAndRest[0])]
	tlAndwh := strings.Split(idAndRest[1], ": ")
	tl := strings.Split(tlAndwh[0], ",")
	wh := strings.Split(tlAndwh[1], "x")
	left, _ := strconv.Atoi(tl[0])
	top, _ := strconv.Atoi(tl[1])
	width, _ := strconv.Atoi(wh[0])
	height, _ := strconv.Atoi(wh[1])
	return Claim {
		Id: id,
		Left: left,
		Top: top,
		Width: width,
		Height: height,
	}
}
