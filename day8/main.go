package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type Node struct {
	ChildCount int
	MetaCount int
	Children []*Node
	Meta []int
}

func main() {
	dat, _ := ioutil.ReadFile("day8/inp.txt")
	strdat := strings.Split(string(dat), "\n")
	ns := parse(strdat[0])
	part1(ns)
	part2(ns)
}

func part1(ns []int) {
	n, _ := getNode(ns, 0)
	fmt.Println(getSum(&n))
}

func getSum(n *Node) int {
	sum := 0
	for i := 0; i < n.MetaCount; i++ {
		sum += n.Meta[i]
	}
	for i := 0; i < n.ChildCount; i++ {
		sum += getSum(n.Children[i])
	}
	return sum
}

func getNode(ns []int, start int) (Node, int){
	cur := Node{
		ChildCount: ns[start],
		MetaCount: ns[start+1],
	}
	childStart := start+2
	for i:=0; i < cur.ChildCount; i++{
		child, newStart := getNode(ns, childStart)
		cur.Children = append(cur.Children, &child)
		childStart = newStart
	}
	cur.Meta = ns[childStart: childStart+cur.MetaCount]
	
	return cur, childStart+cur.MetaCount
}

func part2(ns []int) {
	n, _ := getNode(ns, 0)
	fmt.Println(getSum2(&n))
}

func getSum2(n *Node) int {
	sum := 0
	if n.ChildCount == 0 {
		for i := 0; i < n.MetaCount; i++ {
			sum += n.Meta[i]
		}
		return sum
	}

	for i := 0; i < n.MetaCount; i++ {
		m := n.Meta[i]-1
		if m < 0 || m >= n.ChildCount { continue }
		sum += getSum2(n.Children[m])
	}
	return sum
}


func parse(s string) []int {
	g := strings.Split(s, " ")
	ret := make([]int, len(g))
	for i,x:=range g{
		h, _ := strconv.Atoi(x)
		ret[i] = h;
	}
	return ret
}
