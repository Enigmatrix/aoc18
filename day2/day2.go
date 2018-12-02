package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("day2/day2.txt")
	strdat := strings.Split(string(dat), "\n")
	part1(strdat)
	part2(strdat)
}

func part1(strdat []string) {
	c2, c3 := 0, 0
	for _, x := range strdat {
		if2, if3 := false, false
		for _, c := range x {
			count := strings.Count(x, string(c))
			if count == 2 {
				if2 = true
			} else if count == 3{
				if3 = true
			}
		}
		if if2 { c2++	}
		if if3 { c3++	}
	}
	fmt.Println(c2*c3)
}

func part2(strdat []string) {
	for i1, x1 := range strdat {
		for i2:= 0; i2 < i1; i2++{
			if f, i := eq1(x1, strdat[i2]); f {
				for d:=0; d < len(x1); d++ {
					if d != i {
						fmt.Print(string(x1[d]))
					}
				}
				return
			}
		}
	}
}

func eq1(s1 string, s2 string) (bool, int) {
	alr := false
	c := 0
	for i:= 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if alr {
				return false, 0
			} else {
				alr = true
				c = i
			}
		}
	}
	return true, c
}
