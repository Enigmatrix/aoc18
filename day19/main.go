package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"time"
)

type Instruction struct {
	Opcode string
	A int
	B int
	C int
}

var ipr int
var program []Instruction

func main() {
	dat, _ := ioutil.ReadFile("day19/inp.txt")
	strdat := strings.Split(string(dat), "\n")
	fmt.Sscanf(strdat[0], "#ip %d", &ipr)
	strdat = strdat[1:]
	program = parseI(strdat)
	part1()
	part2()
}

func part1() {

	var regs [6]int
	for regs[ipr] = 0; regs[ipr] < len(program); regs[ipr]++ {
		inst := program[regs[ipr]]
		switch inst.Opcode {
		case "addr":
			regs[inst.C] = regs[inst.A] + regs[inst.B]
		case "addi":
			regs[inst.C] = regs[inst.A] + inst.B
		case "mulr":
			regs[inst.C] = regs[inst.A] * regs[inst.B]
		case "muli":
			regs[inst.C] = regs[inst.A] * inst.B
		case "banr":
			regs[inst.C] = regs[inst.A] & regs[inst.B]
		case "bani":
			regs[inst.C] = regs[inst.A] & inst.B
		case "borr":
			regs[inst.C] = regs[inst.A] | regs[inst.B]
		case "bori":
			regs[inst.C] = regs[inst.A] | inst.B
		case "setr":
			regs[inst.C] = regs[inst.A]
		case "seti":
			regs[inst.C] = inst.A
		case "gtir":
			regs[inst.C] = toInt(inst.A > regs[inst.B])
		case "gtri":
			regs[inst.C] = toInt(regs[inst.A] > inst.B)
		case "gtrr":
			regs[inst.C] = toInt(regs[inst.A] > regs[inst.B])
		case "eqir":
			regs[inst.C] = toInt(inst.A == regs[inst.B])
		case "eqri":
			regs[inst.C] = toInt(regs[inst.A] == inst.B)
		case "eqrr":
			regs[inst.C] = toInt(regs[inst.A] == regs[inst.B])
		}
	}
	fmt.Println(regs[0])
}
func part2() {


	var regs [6]int
	regs[0] = 1
	wut:=false
	for regs[ipr] = 0; regs[ipr] < len(program); regs[ipr]++ {
		inst := program[regs[ipr]]
		fmt.Println(regs)
		fmt.Println(inst)

		//replace results
		if inst.Opcode == "eqrr" && inst.A == 2 && inst.B == 5 && inst.C == 2  && regs[1] == 7300{
			regs[0] = 0
			regs[1] = 10551330
			regs[2]= 10551329
			regs[3]=1
			regs[4] = 4
			regs[5] = 10551329
		}
		if inst.Opcode == "seti" && inst.A == 1 && inst.B == 8 && inst.C == 4{
			regs[0] = 1 + 10551329 + 137 + 77017
			regs[1] = 10551329
			regs[2]= 1
			regs[3] = 10551329
			regs[4] = 15
			regs[5] = 10551329
			wut = true
			continue
		}
		if (wut){
		time.Sleep(1*time.Second)
		}
		switch inst.Opcode {
		case "addr":
			regs[inst.C] = regs[inst.A] + regs[inst.B]
		case "addi":
			regs[inst.C] = regs[inst.A] + inst.B
		case "mulr":
			regs[inst.C] = regs[inst.A] * regs[inst.B]
		case "muli":
			regs[inst.C] = regs[inst.A] * inst.B
		case "banr":
			regs[inst.C] = regs[inst.A] & regs[inst.B]
		case "bani":
			regs[inst.C] = regs[inst.A] & inst.B
		case "borr":
			regs[inst.C] = regs[inst.A] | regs[inst.B]
		case "bori":
			regs[inst.C] = regs[inst.A] | inst.B
		case "setr":
			regs[inst.C] = regs[inst.A]
		case "seti":
			regs[inst.C] = inst.A
		case "gtir":
			regs[inst.C] = toInt(inst.A > regs[inst.B])
		case "gtri":
			regs[inst.C] = toInt(regs[inst.A] > inst.B)
		case "gtrr":
			regs[inst.C] = toInt(regs[inst.A] > regs[inst.B])
		case "eqir":
			regs[inst.C] = toInt(inst.A == regs[inst.B])
		case "eqri":
			regs[inst.C] = toInt(regs[inst.A] == inst.B)
		case "eqrr":
			regs[inst.C] = toInt(regs[inst.A] == regs[inst.B])
		}
		fmt.Println(regs)
	}
	fmt.Println(regs[0])

}

func parseI(s []string) []Instruction{
	ret:=make([]Instruction, len(s)-1)
	for i:=0; i < len(s)-1;i++{
		fmt.Sscanf(s[i], "%s %d %d %d ",
			&ret[i].Opcode,
			&ret[i].A,
			&ret[i].B,
			&ret[i].C)
	}
	return ret
}

func toInt(b bool) int{
	if b {
		return 1
	} else {
		return 0
	}
}
