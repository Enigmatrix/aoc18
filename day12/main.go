package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

var initialState map[int]bool
var mappy map[[5]bool]bool

func main() {
	dat, _ := ioutil.ReadFile("day12/inp.txt")
	strdat := strings.Split(string(dat), "\n")
	wut := ""
	mappy = make(map[[5]bool]bool)
	initialState = make(map[int]bool)
	fmt.Sscanf(strdat[0], "initial state: %s\n", &wut)
	strdat = strdat[2:len(strdat)-1]

	for i, c := range wut  {
		initialState[i] = c == '#'
	}
	for _, s := range strdat {
		var c [5]rune
		var out rune
		fmt.Sscanf(s, "%c%c%c%c%c => %c", &c[0],&c[1],&c[2],&c[3],&c[4],&out)
		var inp [5]bool
		for i:=0; i<5; i++{
			inp[i] = c[i] == '#'
		}
		outb := out == '#'
		mappy[inp] = outb
	}
	part1()
	part2()
}

func part1() {
	for gen := 0; gen < 20; gen++ {

		min, max := minMax()
		for i:= 0; i<2; i++ {
			initialState[min-i-1] = false
			initialState[max+i+1] = false
		}

		newState := make(map[int]bool)
		for i, v := range initialState{
			var key [5]bool
			key[0] = initialState[i-2]
			key[1] = initialState[i-1]
			key[2] = v
			key[3] = initialState[i+1]
			key[4] = initialState[i+2]
			newState[i] = mappy[key]
		}
		initialState = newState
	}

	fmt.Println(getCount())
}

func minMax() (int, int){
	min, max := 10000, -10000
	for i, _:= range initialState{
		if i < min {min=i}
		if i > max {max=i}
	}
	return min, max
}

func minMax2() (int, int){
	min, max := 10000, -10000
	for i, v:= range initialState{
		if i < min && v {min=i}
		if i > max && v {max=i}
	}
	return min, max
}


func part2() {
	x := 50000000000 - 3065
	fmt.Println(189521 + 62*x)
}

func getCount() int {
	c:= 0
	for i, v := range initialState {
		if v { c+=i }
	}
	return c
}
