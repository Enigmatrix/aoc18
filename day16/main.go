package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/deckarep/golang-set"
)

var initialState map[int]bool
var mappy map[[5]bool]bool

type Instruction struct {
	Opcode int
	A int
	B int
	C int
}

type InstructionEffect struct{
	Before [4]int
	Opcode int
	A int
	B int
	C int
	After [4]int
}

var effects []InstructionEffect
var program []Instruction

func main() {
	dat, _ := ioutil.ReadFile("day16/inp.txt")
	strdat := strings.Split(string(dat), "\n\n\n\n")
	p1 := strings.Split(strdat[0], "\n\n")
	effects = parseIE(p1)
	program = parseI(strings.Split(strdat[1], "\n"))
	part1()
	part2()
}

func part1() {
	c := 0
	for _, e := range effects {
		o := 0
		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A] + e.Before[e.B] == e.After[e.C] { o++ }
		if isReg(e.A) && e.Before[e.A] + e.B == e.After[e.C] { o++ }

		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A] * e.Before[e.B] == e.After[e.C] { o++ }
		if isReg(e.A) && e.Before[e.A] * e.B == e.After[e.C] { o++ }

		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A] & e.Before[e.B] == e.After[e.C] { o++ }
		if isReg(e.A) && e.Before[e.A] & e.B == e.After[e.C] { o++ }

		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A] | e.Before[e.B] == e.After[e.C] { o++ }
		if isReg(e.A) && e.Before[e.A] | e.B == e.After[e.C] { o++ }

		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A]  == e.After[e.C] { o++ }
		if e.A == e.After[e.C] { o++ }

		if isReg(e.B) && toInt(e.A > e.Before[e.B]) == e.After[e.C] { o++ }
		if isReg(e.A) && toInt(e.Before[e.A] > e.B) == e.After[e.C] { o++ }
		if isReg(e.A) && isReg(e.B) && toInt(e.Before[e.A] > e.Before[e.B]) == e.After[e.C] { o++ }

		if isReg(e.B) && toInt(e.A == e.Before[e.B]) == e.After[e.C] { o++ }
		if isReg(e.A) && toInt(e.Before[e.A] == e.B) == e.After[e.C] { o++ }
		if isReg(e.A) && isReg(e.B) && toInt(e.Before[e.A] == e.Before[e.B]) == e.After[e.C] { o++ }

		if o >= 3 { c++ }
	}
	fmt.Println(c)

}
func part2() {
	var instr [16]mapset.Set
	for i:= 0; i < 16; i++{
		instr[i] = mapset.NewSetWith(0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15)
	}
	for _, e := range effects {
		k := mapset.NewSet()
		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A] + e.Before[e.B] == e.After[e.C] { k.Add(0) }
		if isReg(e.A) && e.Before[e.A] + e.B == e.After[e.C] { k.Add(1) }

		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A] * e.Before[e.B] == e.After[e.C] { k.Add(2) }
		if isReg(e.A) && e.Before[e.A] * e.B == e.After[e.C] { k.Add(3) }

		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A] & e.Before[e.B] == e.After[e.C] { k.Add(4) }
		if isReg(e.A) && e.Before[e.A] & e.B == e.After[e.C] { k.Add(5) }

		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A] | e.Before[e.B] == e.After[e.C] { k.Add(6) }
		if isReg(e.A) && e.Before[e.A] | e.B == e.After[e.C] { k.Add(7) }

		if isReg(e.A) && isReg(e.B) &&  e.Before[e.A]  == e.After[e.C] { k.Add(8) }
		if e.A == e.After[e.C] { k.Add(9) }

		if isReg(e.B) && toInt(e.A > e.Before[e.B]) == e.After[e.C] { k.Add(10) }
		if isReg(e.A) && toInt(e.Before[e.A] > e.B) == e.After[e.C] { k.Add(11) }
		if isReg(e.A) && isReg(e.B) && toInt(e.Before[e.A] > e.Before[e.B]) == e.After[e.C] { k.Add(12) }

		if isReg(e.B) && toInt(e.A == e.Before[e.B]) == e.After[e.C] { k.Add(13) }
		if isReg(e.A) && toInt(e.Before[e.A] == e.B) == e.After[e.C] { k.Add(14) }
		if isReg(e.A) && isReg(e.B) && toInt(e.Before[e.A] == e.Before[e.B]) == e.After[e.C] { k.Add(15) }

		instr[e.Opcode] = instr[e.Opcode].Intersect(k)
	}
	var instrFinal [16]int
	for !allEmptySet(&instr) {
		for i := range instr{
			if len(instr[i].ToSlice()) == 1 {
				k := instr[i].Pop()
				instrFinal[i] = k.(int)
				for j := range instr{
					if i != j{
						instr[j].Remove(k)
					}
				}
			}
		}
	}

	var regs [4]int
	for _, inst := range program {
		switch instrFinal[inst.Opcode] {
		case 0:
			regs[inst.C] = regs[inst.A] + regs[inst.B]
		case 1:
			regs[inst.C] = regs[inst.A] + inst.B
		case 2:
			regs[inst.C] = regs[inst.A] * regs[inst.B]
		case 3:
			regs[inst.C] = regs[inst.A] * inst.B
		case 4:
			regs[inst.C] = regs[inst.A] & regs[inst.B]
		case 5:
			regs[inst.C] = regs[inst.A] & inst.B
		case 6:
			regs[inst.C] = regs[inst.A] | regs[inst.B]
		case 7:
			regs[inst.C] = regs[inst.A] | inst.B
		case 8:
			regs[inst.C] = regs[inst.A]
		case 9:
			regs[inst.C] = inst.A
		case 10:
			regs[inst.C] = toInt(inst.A > regs[inst.B])
		case 11:
			regs[inst.C] = toInt(regs[inst.A] > inst.B)
		case 12:
			regs[inst.C] = toInt(regs[inst.A] > regs[inst.B])
		case 13:
			regs[inst.C] = toInt(inst.A == regs[inst.B])
		case 14:
			regs[inst.C] = toInt(regs[inst.A] == inst.B)
		case 15:
			regs[inst.C] = toInt(regs[inst.A] == regs[inst.B])
		}
	}
	fmt.Println(regs[0])

}

func allEmptySet(g *[16]mapset.Set) bool{
	for i := range g{
		if len(g[i].ToSlice()) != 0{
			return false
		}
	}
	return true
}

func toInt(b bool) int{
	if b {
		return 1
	} else {
		return 0
	}
}

func isReg(o int) bool {
	return o >= 0 && o <4
}

func parseIE(s []string) []InstructionEffect {
	ret:= make([]InstructionEffect, len(s))
	for i, g := range s {
		fmt.Sscanf(g, "Before: [%d, %d, %d, %d]\n%d %d %d %d\nAfter: [%d, %d, %d, %d]",
			&ret[i].Before[0],
			&ret[i].Before[1],
			&ret[i].Before[2],
			&ret[i].Before[3],
			&ret[i].Opcode,
			&ret[i].A,
			&ret[i].B,
			&ret[i].C,
			&ret[i].After[0],
			&ret[i].After[1],
			&ret[i].After[2],
			&ret[i].After[3],
		)
	}
	return ret
}

func parseI(s []string) []Instruction{
	ret:=make([]Instruction, len(s)-1)
	for i:=0; i < len(s)-1;i++{
		fmt.Sscanf(s[i], "%d %d %d %d ",
			&ret[i].Opcode,
			&ret[i].A,
			&ret[i].B,
			&ret[i].C)
	}
	return ret
}
