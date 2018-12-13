package main

import (
	"fmt"
	"io/ioutil"
	"container/list"
)

func main() {
	dat, _ := ioutil.ReadFile("day5/inp.txt")
	strdat := string(dat)
	strdat = strdat[0:len(strdat)-1]
	part1(strdat)
	part2(strdat)
}

func part1(strdat string) {
	list := list.New()
	for _, c := range strdat {
		list.PushBack(c)
	}

	for e:= list.Front().Next(); e != nil; e = ifNil(e, list.Front()).Next(){
		//fmt.Println(string(e.Value.(int32)))
		if e.Prev().Value.(int32)-32 == e.Value.(int32) || e.Prev().Value.(int32) == e.Value.(int32)-32 {
			g := ifNil(e.Prev().Prev(), list.Front()).Prev()

			list.Remove(e.Prev())
			list.Remove(e)

			e = g
		}
	}
	fmt.Println("Length: ", list.Len())
}

func ifNil(e1 *list.Element, e2 *list.Element) *list.Element{
	if e1 == nil {
		return e2
	}
	return e1
}

func part2(strdat string) {

	listOrig := list.New()
	chrs := make(map[int32]struct{})
	for _, c := range strdat {
		listOrig.PushBack(c)
		if c > 'Z' {
			chrs[c] = struct{}{}
		} else {
			chrs[c+32] = struct{}{}
		}
	}

	minLen := 1000000
	for c, _ := range chrs{
		list := list.New()
		for _, c2 := range(strdat){
			if c2 != c && c2+32 != c {
				list.PushBack(c2)
			}
		}

		for e:= list.Front().Next(); e != nil; e = ifNil(e, list.Front()).Next(){
			//fmt.Println(string(e.Value.(int32)))
			if e.Prev().Value.(int32)-32 == e.Value.(int32) || e.Prev().Value.(int32) == e.Value.(int32)-32 {
				g := ifNil(e.Prev().Prev(), list.Front()).Prev()

				list.Remove(e.Prev())
				list.Remove(e)

				e = g
			}
		}

		if list.Len() < minLen {
			minLen = list.Len()
		}
	}

	fmt.Println("Length: ", minLen)
}

