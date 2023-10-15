package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

	fmt.Println("Distinct molecules are ", processToken())
}

type posKey struct {
	pos int
	key string
}

func processToken() int {
	molecules := make(map[string]bool)
	for i := 0; i < len(token); i++ {
		for k, v := range subMap {
			if len(token) >= i+len(k) && token[i:i+len(k)] == k {
				for j := 0; j < len(v); j++ {
					mol := token[0:i] + v[j] + token[i+len(k):]
					molecules[mol] = true
				}
			}
		}
	}

	return len(molecules)
}

func synLen() int {
	for _, v := range subMap["e"] {
		length, res := synth(v, 0)
		if res == token {
			return length
		}
	}
	return 0
}

func synth(s string, pos int) (int, string) {
	if pos == len(s)+1 {
		return len(s), s
	}
	return 0, ""
}

var subMap = make(map[string][]string)
var token string

func parseLine(l string) {
	if len(l) == 0 {
		return
	}
	val := strings.Split(l, " ")
	if len(val) == 1 {
		token = val[0]
		return
	}
	if _, ok := subMap[val[0]]; !ok {
		subMap[val[0]] = []string{val[2]}
		return
	}
	subMap[val[0]] = append(subMap[val[0]], val[2])
}
