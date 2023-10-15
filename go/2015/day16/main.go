package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	name := flag.String("i", "input.txt", "Input file")
	flag.Parse()

	fh, err := os.Open(*name)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		parseLine(fs.Text())
	}
	needle := sue{}
	needle.akitas = 0
	needle.cars = 2
	needle.cats = 7
	needle.children = 3
	needle.goldfish = 5
	needle.perfumes = 1
	needle.pomeranians = 3
	needle.samoyeds = 2
	needle.trees = 3
	needle.vizslas = 0
	fmt.Println("Sue is number", findSue(needle))
}

func findSue(s sue) int {
	for i := 1; i <= len(sues); i++ {
		sc := sues[i]
		if !same(sc.akitas, s.akitas) {
			continue
		}
		if !same(sc.cars, s.cars) {
			continue
		}
		if !gt(sc.cats, s.cats) {
			continue
		}
		if !same(sc.children, s.children) {
			continue
		}
		if !lt(sc.goldfish, s.goldfish) {
			continue
		}
		if !same(sc.perfumes, s.perfumes) {
			continue
		}
		if !lt(sc.pomeranians, s.pomeranians) {
			continue
		}
		if !same(sc.samoyeds, s.samoyeds) {
			continue
		}
		if !gt(sc.trees, s.trees) {
			continue
		}
		if !same(sc.vizslas, s.vizslas) {
			continue
		}
		return i
	}
	return -1
}

func gt(a, b int) bool {
	if a == -1 {
		return true
	}
	if a > b {
		return true
	}
	return false
}

func lt(a, b int) bool {
	if a == -1 {
		return true
	}
	if a < b {
		return true
	}
	return false
}

func same(a, b int) bool {
	if a == -1 {
		return true
	}
	if a == b {
		return true
	}
	return false
}

type sue struct {
	akitas      int
	cars        int
	cats        int
	children    int
	goldfish    int
	perfumes    int
	pomeranians int
	samoyeds    int
	trees       int
	vizslas     int
}

var sues = make([]sue, 501)

// Sue 484: cats: 0, goldfish: 0, children: 9
var re = regexp.MustCompile("^Sue ([0-9]+): ([a-z]+): ([0-9]+), ([a-z]+): ([0-9]+), ([a-z]+): ([0-9]+)$")

func parseLine(l string) {
	res := re.FindStringSubmatch(l)
	s := sue{}
	s.children = -1
	s.cats = -1
	s.samoyeds = -1
	s.pomeranians = -1
	s.akitas = -1
	s.vizslas = -1
	s.goldfish = -1
	s.trees = -1
	s.cars = -1
	s.perfumes = -1
	addProp := func(in string, val int) {
		switch in {
		case "children":
			s.children = val
		case "cats":
			s.cats = val
		case "samoyeds":
			s.samoyeds = val
		case "pomeranians":
			s.pomeranians = val
		case "akitas":
			s.akitas = val
		case "vizslas":
			s.vizslas = val
		case "goldfish":
			s.goldfish = val
		case "trees":
			s.trees = val
		case "cars":
			s.cars = val
		case "perfumes":
			s.perfumes = val
		}
	}

	addProp(res[2], atoi(res[3]))
	addProp(res[4], atoi(res[5]))
	addProp(res[6], atoi(res[7]))
	sues[atoi(res[1])] = s
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
