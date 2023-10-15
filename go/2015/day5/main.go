package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	s "strings"
)

var naughty = 0
var nice = 0
var naughty2 = 0
var nice2 = 0

func main() {
	input := flag.String("i", "input.txt", "Input string")
	flag.Parse()

	file, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		checkLine(fs.Text())
		newCheckLine(fs.Text())
	}
	fmt.Println("Old naughty lines:", naughty)
	fmt.Println("Old nice lines:", nice)
	fmt.Println("New naughty lines:", naughty2)
	fmt.Println("New nice lines:", nice2)
}

func newCheckLine(l string) {
	var pairs = make(map[string][]int, 0)
	pr := false
	rep := false
	for i := 0; i < len(l); i++ {
		if !rep && charRepeats(i, &l) {
			rep = true
		}
		if i == 0 {
			continue
		}
		pair := l[i-1 : i+1]
		set, ok := pairs[pair]
		if ok && pairValid(set, i) {
			pr = true
		}
		pairs[pair] = append(pairs[pair], i)
	}
	if !pr || !rep {
		naughty2++
		return
	}
	nice2++
}

func pairValid(set []int, i int) bool {
	for _, v := range set {
		if v < i-1 {
			return true
		}
	}
	return false
}

func charRepeats(i int, l *string) bool {
	if i < 2 {
		return false
	}
	if (*l)[i] != (*l)[i-2] {
		return false
	}
	return true
}

func checkLine(l string) {
	badSets := []string{"ab", "cd", "pq", "xy"}
	for _, bs := range badSets {
		if s.Contains(l, bs) {
			naughty++
			return
		}
	}

	hasTwo := false
	vowels := 0
	for k, v := range l {
		switch string(v) {
		case "a", "e", "i", "o", "u":
			vowels++
		}
		if k > 0 && l[k-1] == byte(v) {
			hasTwo = true
		}
	}
	if !hasTwo || vowels < 3 {
		naughty++
		return
	}
	nice++
}
