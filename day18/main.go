package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

var mapOrig [][]rune

func main() {
	dat, _ := ioutil.ReadFile("day18/inp.txt")
	strdat:= string(dat)
	mapOrig = getMap(string(strdat))
	part1()
	part2()
}

type Cart struct{
	LocX int
	LocY int
	CurrentDir rune
	IntersectionTimes int
	IsDeleted bool
}

func part1() {
	dupMap := make([][]rune, len(mapOrig))
	for i:= range mapOrig {
		dupMap[i] = make([]rune, len(mapOrig[i]))
		copy(dupMap[i], mapOrig[i])
	}
	dupMap2 := make([][]rune, len(dupMap))
	for i:=range dupMap {
		dupMap2[i] = make([]rune, len(dupMap[i]))
	}
	for j := 0; j < 10; j++{
		for x := range dupMap {
			for y := range dupMap[x] {
				_, tree, ly := count(x,y, dupMap)
				switch dupMap[x][y]{
				case '.':
					if tree >= 3 {
						dupMap2[x][y] = '|'
					} else{
						dupMap2[x][y] = '.'
					}
				case '|':
					if ly >= 3 {
						dupMap2[x][y] = '#'
					} else{
						dupMap2[x][y] = '|'
					}
				case '#':
					if ly < 1 || tree < 1 {
						dupMap2[x][y] = '.'
					} else{
						dupMap2[x][y] = '#'
					}
				}
			}
		}
		x := dupMap2
		dupMap2 = dupMap
		dupMap = x

	}


	tree, ly := 0,0
	for x := range dupMap {
		for y := range dupMap[x] {
			switch dupMap[x][y]{
			case '|':
				tree++
			case '#':
				ly++
			}
		}
	}
	fmt.Println(tree*ly)



}
func part2(){
	pattern := [28]int { 207088, 203211, 205542, 200466, 201264, 196144, 196020, 192075, 193452, 191425, 194884, 194110, 196803, 193428, 199593, 201348, 208603, 210672, 217634, 217462, 223728, 224553, 226806, 220528, 219849, 213057, 213462, 205700}
	fmt.Println(pattern[(1000000000-900-1)%28])
}

func count(x int, y int, dupMap [][]rune) (int, int, int){
	rx := [8]int {0,0,1,1,1,-1,-1,-1}
	ry := [8]int {1,-1,1,0,-1,1,0,-1}
	c1, c2, c3 := 0,0,0
	for i := range rx {
		nx := rx[i]+x
		ny := ry[i]+y
		if nx >= len(dupMap) || nx < 0 || ny < 0 || ny >= len(dupMap[i]) {continue}
		switch dupMap[nx][ny]{
		case '.':
			c1++
		case '|':
			c2++
		case '#':
			c3++
		}
	}
	
	return c1,c2,c3
}

func getMap(s string) [][]rune {
	rows := strings.Split(string(s), "\n")
	ret := make([][]rune, len(rows)-1)
	rows = rows[:len(rows)-1]
	for ri, _ := range rows {
		ret[ri] = make([]rune, len(rows[ri]))
		for i, c := range rows[ri]{
			ret[ri][i] = c
		}
	}
	return ret
} 
