package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/jupp0r/go-priority-queue"
)

func main() {
	dat, _ := ioutil.ReadFile("day7/inp.txt")
	strdat := strings.Split(string(dat), "\n")
	steps := parseSteps(strdat)
	part1(steps)
	part2(steps)
}

type Step struct {
	Name rune
	Pre rune
}

type Node struct {
	Label rune
	Pre []rune
	Seq []rune
	Processed bool
}

func part1(steps []Step) {
	nodeMap := make(map[rune]*Node)
	for  _, s := range steps {
		pre := getNode(s.Pre, nodeMap)
		this := getNode(s.Name, nodeMap)
		this.Pre = append(this.Pre, s.Pre)
		pre.Seq = append(pre.Seq, s.Name)
	}
	process := pq.New()
	for _, node := range nodeMap {
		if len(node.Pre) == 0 {
			process.Insert(node.Label, float64(node.Label))
		}
	}

	var fina []rune

	for process.Len() > 0 {
		label, _ := process.Pop()
		fina = append(fina, label.(rune))
		node := nodeMap[label.(rune)]
		node.Processed = true
		for _, p := range node.Seq {
			if nodeMap[p].Processed { continue }
			ins := true
			for _, pre := range nodeMap[p].Pre {
				if !nodeMap[pre].Processed {
					ins = false
					break
				}
			}
			if ins {
				process.Insert(p, float64(p))
			}
		}
	}

	for _, c:= range fina {
		fmt.Print(string(c))
	}
}

func getNode(idx rune,ma map[rune]*Node) *Node {
	_, prs := ma[idx]
	if !prs {
		ma[idx] = &Node{
			Label: idx,
		}
	}
	return ma[idx]
}

type Process struct {
	Step rune
	TimeEnd int
}

func part2(steps []Step) {

	nodeMap := make(map[rune]*Node)
	for  _, s := range steps {
		pre := getNode(s.Pre, nodeMap)
		this := getNode(s.Name, nodeMap)
		this.Pre = append(this.Pre, s.Pre)
		pre.Seq = append(pre.Seq, s.Name)
	}
	process := pq.New()
	for _, node := range nodeMap {
		if len(node.Pre) == 0 {
			process.Insert(node.Label, float64(node.Label))
		}
	}

	var workers [5]*Process
	timeNow := 0

	for process.Len() > 0 {
		label, _ := process.Pop()

		var pIdx int
		for i, w := range workers {
			if w == nil {
				pIdx = i
				break
			} else if workers[pIdx].TimeEnd > w.TimeEnd {
				pIdx = i
			}
		}

		if workers[pIdx] != nil {
			node := nodeMap[workers[pIdx].Step]
			node.Processed = true
			timeNow = workers[pIdx].TimeEnd
			for _, p := range node.Seq {
				if nodeMap[p].Processed { continue }
				ins := true
				for _, pre := range nodeMap[p].Pre {
					if !nodeMap[pre].Processed {
						ins = false
						break
					}
				}
				if ins {
					process.Insert(p, float64(p))
				}
			}
		}

		workers[pIdx] = &Process {
			Step: label.(rune),
			TimeEnd: timeNow + 61 + int(label.(rune)-'A'),
		}


		for ;process.Len() == 0; {
			var pIdx int
			for i, w := range workers {
				if workers[pIdx] == nil && w != nil {
					pIdx = i
				} else if w != nil && workers[pIdx].TimeEnd > w.TimeEnd {
					pIdx = i
				}
			}

			if workers[pIdx] == nil {
				fmt.Println(timeNow)
				return
			}

			node := nodeMap[workers[pIdx].Step]
			node.Processed = true
			timeNow = workers[pIdx].TimeEnd
			for _, p := range node.Seq {
				if nodeMap[p].Processed { continue }
				ins := true
				for _, pre := range nodeMap[p].Pre {
					if !nodeMap[pre].Processed {
						ins = false
						break
					}
				}
				if ins {
					process.Insert(p, float64(p))
				}
			}
			workers[pIdx] = nil

		}
	}
}

func noNullWorkers(workers []*Process) bool{
	for _, w := range workers {
		if w == nil {
			return false
		}
	}
	return true
}

func parseSteps(s []string) []Step{
	ret:=make([]Step, len(s)-1)
	for i:=0; i < len(s)-1;i++{
		ret[i] = parseStep(s[i])
	}
	return ret
}


func parseStep(s string) Step {
	g :=strings.Split(s, " must be finished before step ")
	s1 := g[0][5]
	s2 := g[1][0]
	return Step {
		Name: rune(s2),
		Pre: rune(s1),
	}
}
