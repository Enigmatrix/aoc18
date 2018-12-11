package main

import (
	"fmt"
	"strconv"
)

var Serial int

func main() {
	serial, _ := strconv.Atoi("2694")
	Serial = serial
	part1()
	part2()
}

func part1() {
	var stuff [301][301]int
	for x := 1; x <= 300; x++{
		for y := 1; y <= 300; y++{
			stuff[x][y] = powah(x, y)
		}
	}
	max := -1000
	xmax := 0
	ymax := 0
	for x := 1; x <= 300-3; x++{
		for y := 1; y <= 300-3; y++{

			s := 0
			for x1 := x; x1 < x+3; x1++{
				for y1 := y; y1 < y+3; y1++{
					s += stuff[x1][y1]
				}
			}
			if max < s {
				max = s
				xmax = x
				ymax = y
			}

		}
	}
	fmt.Println(" Coords", xmax, ",", ymax)
}

func powah(x int, y int) int {
	rack := x+10
	hDigit := (((rack*y+Serial)*rack) % 1000) / 100
	return hDigit - 5
}

func part2() {
	var stuff [301][301]int
	for x := 1; x <= 300; x++{
		for y := 1; y <= 300; y++{
			stuff[x][y] = powah(x, y)
		}
	}
	max := -1000
	xmax := 0
	ymax := 0
	szmax := 0
	cursum := 0
	prevxsum := 0
	for sz := 1; sz <= 300; sz++{
		cursum = 0
		for x1 := 1; x1 < sz; x1++{
			for y1 := 1; y1 < sz; y1++{
				cursum += stuff[x1][y1]
			}
		}
		prevxsum = cursum
		for x := 1; x <= 300-sz; x++{
			if x != 1{
				for y1 := 1; y1 <= sz; y1++ {
					prevxsum -= stuff[x-1][y1]
					prevxsum += stuff[x+sz-1][y1]
				}
				cursum = prevxsum
			}
			for y := 1; y <= 300-sz; y++{

				if y != 1{
					for x1 := x; x1 < x+sz; x1++ {
						cursum -= stuff[x1][y-1]
						cursum += stuff[x1][y+sz-1]
					}
				}

				if max < cursum {
					max = cursum
					xmax = x
					ymax = y
					szmax = sz
				}
			}
		}
	}
	fmt.Println(" Coords", xmax, ",", ymax, ",",szmax)
}
