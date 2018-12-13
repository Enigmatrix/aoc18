package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/jupp0r/go-priority-queue"
)

var rideMapOrig [][]rune

func main() {
	dat, _ := ioutil.ReadFile("day13/inp.txt")
	strdat:= string(dat)
	/*strdat =  `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`*/
	rideMapOrig = getMap(string(strdat))
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
	carts := pq.New()
	cartMap := make(map[int]struct{})
	rideMap := make([][]rune, len(rideMapOrig))
	for i := range rideMapOrig {
			rideMap[i] = make([]rune, len(rideMapOrig[i]))
			copy(rideMap[i], rideMapOrig[i])
	}
	dirMap := map[rune][3]rune{
		'>': [3]rune{'^', '>', 'v'},
		'<': [3]rune{'v', '<', '^'},
		'v': [3]rune{'>', 'v', '<'},
		'^': [3]rune{'<', '^', '>'},
	}
	for x, _ := range rideMap {
		for y, c := range rideMap[x]{
			if c == '^' || c == 'v' || c == '>' || c == '<' {
				if c == '^' || c == 'v' {
					rideMap[x][y] = '|'
				} else {
					rideMap[x][y] = '-'
				}
				cartMap[x*1000+y] = struct{}{}
				carts.Insert(&Cart{
					LocX: x,
					LocY: y,
					CurrentDir: c,
					IntersectionTimes: 0,
				}, float64(x*1000+y))
			}
		}
	}
	for {
		toInsertCarts := make([]*Cart, 0)
		for c, p := carts.Pop(); p == nil; c, p = carts.Pop(){
			cart := c.(*Cart)
			newX, newY := cart.LocX, cart.LocY
			switch cart.CurrentDir {
			case '^':
				newX--
			case 'v':
				newX++
			case '<':
				newY--
			case '>':
				newY++
			}

			if _, p := cartMap[newX*1000+newY]; p {
				fmt.Println("Collision:", newY, newX)
				return
			}

			switch rideMap[newX][newY]{
			case '/':
				switch cart.CurrentDir {
				case '>':
					cart.CurrentDir = '^'
				case '^':
					cart.CurrentDir = '>'
				case 'v':
					cart.CurrentDir = '<'
				case '<':
					cart.CurrentDir = 'v'
				}
			case '\\':
				switch cart.CurrentDir {
				case '>':
					cart.CurrentDir = 'v'
				case '^':
					cart.CurrentDir = '<'
				case 'v':
					cart.CurrentDir = '>'
				case '<':
					cart.CurrentDir = '^'
				}
			case '+':
				cart.CurrentDir = dirMap[cart.CurrentDir][cart.IntersectionTimes % 3]
				cart.IntersectionTimes++
			}

			delete(cartMap, 1000*cart.LocX+cart.LocY)
			cart.LocX = newX
			cart.LocY = newY
			cartMap[1000*cart.LocX+cart.LocY] = struct{}{}
			toInsertCarts = append(toInsertCarts, cart)

		}
		for _, cart := range toInsertCarts{
			//fmt.Println(cart)
			carts.Insert(cart, float64(cart.LocX*1000+cart.LocY))
		}
	}
}
func part2(){
	carts := pq.New()
	cartMap := make(map[int]*Cart)
	rideMap := make([][]rune, len(rideMapOrig))
	for i := range rideMapOrig {
			rideMap[i] = make([]rune, len(rideMapOrig[i]))
			copy(rideMap[i], rideMapOrig[i])
	}
	dirMap := map[rune][3]rune{
		'>': [3]rune{'^', '>', 'v'},
		'<': [3]rune{'v', '<', '^'},
		'v': [3]rune{'>', 'v', '<'},
		'^': [3]rune{'<', '^', '>'},
	}
	for x, _ := range rideMap {
		for y, c := range rideMap[x]{
			if c == '^' || c == 'v' || c == '>' || c == '<' {
				if c == '^' || c == 'v' {
					rideMap[x][y] = '|'
				} else {
					rideMap[x][y] = '-'
				}
				cart := &Cart{
					LocX: x,
					LocY: y,
					CurrentDir: c,
					IntersectionTimes: 0,
				}
				cartMap[x*1000+y] = cart
				carts.Insert(cart, float64(x*1000+y))
			}
		}
	}
	for {
		toInsertCarts := make([]*Cart, 0)
		for c, p := carts.Pop(); p == nil; c, p = carts.Pop(){
			cart := c.(*Cart)
			if cart.IsDeleted { continue }
			newX, newY := cart.LocX, cart.LocY
			switch cart.CurrentDir {
			case '^':
				newX--
			case 'v':
				newX++
			case '<':
				newY--
			case '>':
				newY++
			}

			switch rideMap[newX][newY]{
			case '/':
				switch cart.CurrentDir {
				case '>':
					cart.CurrentDir = '^'
				case '^':
					cart.CurrentDir = '>'
				case 'v':
					cart.CurrentDir = '<'
				case '<':
					cart.CurrentDir = 'v'
				}
			case '\\':
				switch cart.CurrentDir {
				case '>':
					cart.CurrentDir = 'v'
				case '^':
					cart.CurrentDir = '<'
				case 'v':
					cart.CurrentDir = '>'
				case '<':
					cart.CurrentDir = '^'
				}
			case '+':
				cart.CurrentDir = dirMap[cart.CurrentDir][cart.IntersectionTimes % 3]
				cart.IntersectionTimes++
			}

			delete(cartMap, 1000*cart.LocX+cart.LocY)
			cart.LocX = newX
			cart.LocY = newY
			if other, p := cartMap[newX*1000+newY]; p {
				other.IsDeleted = true
				cart.IsDeleted = true
				delete(cartMap, 1000*newX+newY)
				toInsertCarts = append(toInsertCarts, cart)
			} else {
				cartMap[1000*cart.LocX+cart.LocY] = cart
				toInsertCarts = append(toInsertCarts, cart)
			}


		}

		for _, cart := range toInsertCarts{
			if cart.IsDeleted {	continue }
			carts.Insert(cart, float64(cart.LocX*1000+cart.LocY))
		}
		if carts.Len() == 1 {
			c, _ := carts.Pop()
			f := c.(*Cart)
			fmt.Println("End: ", f.LocY, f.LocX)
			return
		}
	}
}

func getMap(s string) [][]rune {
	rows := strings.Split(string(s), "\n")
	ret := make([][]rune, len(rows))
	for ri, _ := range rows {
		ret[ri] = make([]rune, len(rows[ri]))
		for i, c := range rows[ri]{
			ret[ri][i] = c
		}
	}
	return ret
} 
