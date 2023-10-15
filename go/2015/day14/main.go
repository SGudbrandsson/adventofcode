package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type props struct {
	speed, dur, rest, dist, xp int
}

var deers = map[string]props{}

func main() {
	name := flag.String("i", "input.txt", "input file with reindeer speeds")
	dur := flag.Int("d", 2503, "Duration of the race")
	flag.Parse()

	f, err := os.Open(*name)
	if err != nil {
		panic(err)
	}

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		parseLine(fs.Text())
	}
	race(*dur)

	fmt.Println(deers)
	fmt.Println("Max is", getMax())
	fmt.Println("Max XP is", getMaxXP())
}

var re = regexp.MustCompile("^([A-Za-z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.$")

func parseLine(l string) {
	// Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
	// 1             2           3                                 4
	vars := re.FindStringSubmatch(l)
	deer := props{}
	deer.speed, deer.dur, deer.rest = atoi(vars[2]), atoi(vars[3]), atoi(vars[4])
	deers[vars[1]] = deer
}

func atoi(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return r
}

func race(dur int) {
	for i := 1; i < dur; i++ {
		for k, v := range deers {
			if inRace(i, v.dur, v.rest) {
				v.dist += v.speed
				deers[k] = v
			}
		}
		giveXP()
	}
}

func giveXP() {
	lead := []string{}
	max := getMax()
	for k, v := range deers {
		if v.dist == max {
			lead = append(lead, k)
		}
	}
	for _, v := range lead {
		deer := deers[v]
		deer.xp++
		deers[v] = deer
	}
}

func inRace(t, dur, rest int) bool {
	tot := dur + rest
	rem := t % tot
	if rem > 0 && rem <= dur {
		return true
	}
	return false
}

func getMax() int {
	max := 0
	for _, v := range deers {
		if v.dist > max {
			max = v.dist
		}
	}
	return max
}

func getMaxXP() int {
	max := 0
	for _, v := range deers {
		if v.xp > max {
			max = v.xp
		}
	}
	return max
}
