package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	dat, _ := ioutil.ReadFile("day1.txt")
	strdat := strings.Split(string(dat), "\n")
	part1(strdat)
	part2(strdat)
}

func part1(strdat []string){
	total := 0
	for i:= 0; i < len(strdat)-1; i++ {
		op := strdat[i][0]
		num, _ := strconv.Atoi(strdat[i][1:len(strdat[i])])
		if op == '+' {
			total += num
		} else {
			total -= num
		}
	}
	fmt.Println("Total is", total)
}

func part2(strdat []string){
	itotal := 0
	total := 0
	freq := make([]int, len(strdat))
	for {
		for i:= 0; i < len(strdat)-1; i++ {
			itotal++
			op := strdat[i][0]
			num, _ := strconv.Atoi(strdat[i][1:len(strdat[i])])
			if op == '+' {
				total += num
			} else {
				total -= num
			}
			for _, v := range freq[:itotal] {
				if total == v {
					fmt.Print("Seen ", total)
					return
				}
			}

			freq[itotal] = total
		}
		freq = append(freq, make([]int, len(strdat))...)
	}

}
