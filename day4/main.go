package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"sort"
)

type Event struct {
	Time time.Time
	Guard int
	Type int
}

const Sleep = 0
const Awake = 1
const Guard = 2

func main() {
	dat, _ := ioutil.ReadFile("day4/inp.txt")
	strdat := strings.Split(string(dat), "\n")
	events := parseEvents(strdat)
	sort.Slice(events[:], func(i, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})
	part1(events)
	part2(events)
}

type Duty struct {
	SleepTime int
	LastSleepTime time.Time
	Awake bool
	SleepMinutes [60]int
}

func part1(events []Event) {
	guards := make(map[int]*Duty)
	lastGuard := -1
	for _, event := range events {
		switch event.Type {
		case Guard:
			if lastGuard != -1 {
				last := guards[lastGuard]
				if !last.Awake {
					dur := 60 - last.LastSleepTime.Minute()
					last.SleepTime += dur
					for i := 0; i < dur; i++ {
						last.SleepMinutes[i + last.LastSleepTime.Minute()]++
					}
				}
			}
			lastGuard = event.Guard
			last, prs := guards[lastGuard]
			if !prs {
				last = &Duty{}
				guards[lastGuard] = last
			}
			last.Awake = true
			break
		case Awake:
			last := guards[lastGuard]
			dur := int(event.Time.Sub(last.LastSleepTime).Minutes())
			last.SleepTime += dur
			for i := 0; i < dur; i++ {
				last.SleepMinutes[i + last.LastSleepTime.Minute()]++
			}
			last.Awake = true
			break
		case Sleep:
			last := guards[lastGuard]
			last.LastSleepTime = event.Time
			last.Awake = false
			break
		}
	}
	max := 0
	maxMin:=0
	maxMinT := 0
	idMax := 0
	for k, v := range guards {
		if v.SleepTime > max {
			max = v.SleepTime
			idMax = k
		}
	}
	for i, x := range guards[idMax].SleepMinutes {
		if x > maxMin {
			maxMin = x
			maxMinT = i
		}
	}
	fmt.Println( idMax ," x ", maxMinT, " = ", maxMinT * idMax)
}

func part2(events []Event) {
	guards := make(map[int]*Duty)
	lastGuard := -1
	for _, event := range events {
		switch event.Type {
		case Guard:
			if lastGuard != -1 {
				last := guards[lastGuard]
				if !last.Awake {
					dur := 60 - last.LastSleepTime.Minute()
					last.SleepTime += dur
					for i := 0; i < dur; i++ {
						last.SleepMinutes[i + last.LastSleepTime.Minute()]++
					}
				}
			}
			lastGuard = event.Guard
			last, prs := guards[lastGuard]
			if !prs {
				last = &Duty{}
				guards[lastGuard] = last
			}
			last.Awake = true
			break
		case Awake:
			last := guards[lastGuard]
			dur := int(event.Time.Sub(last.LastSleepTime).Minutes())
			last.SleepTime += dur
			for i := 0; i < dur; i++ {
				last.SleepMinutes[i + last.LastSleepTime.Minute()]++
			}
			last.Awake = true
			break
		case Sleep:
			last := guards[lastGuard]
			last.LastSleepTime = event.Time
			last.Awake = false
			break
		}
	}
	maxMinT := 0
	min:=0
	idMax := 0
	for k, v := range guards {
		for i, x := range v.SleepMinutes {
			if x > maxMinT {
				maxMinT = x
				idMax = k
				min = i
			}
		}
	}
	fmt.Println( idMax ," x ", min, " = ", min * idMax)
}

func parseEvents(s []string) []Event {
	ret := make([]Event, len(s)-1)
	for i:=0; i< len(s)-1; i++ {
		ret[i] = parseEvent(s[i])
	}
	return ret
}

func parseEvent(s string) Event {
	timeAndRest := strings.Split(s, "] ")
	time, _ := time.Parse("2006-01-02 15:04", timeAndRest[0][1:len(timeAndRest[0])])
	guard, Type := 0, 0
	if strings.Contains(timeAndRest[1], "falls asleep") {
		Type = Sleep
	} else if strings.Contains(timeAndRest[1], "wakes up") {
		Type = Awake
	} else {
		stuff := strings.Split(timeAndRest[1], " ")
		num, _ := strconv.Atoi(stuff[1][1: len(stuff[1])])
		Type = Guard
		guard = num
	}
	return Event {
		Time: time,
		Type: Type,
		Guard: guard,
	}
}
