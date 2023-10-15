package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// OPs can be empty string (only one input), RSHIFT, LSHIFT, OR, AND, NOT => bitwise op
type in struct {
	op, a, b        string
	val, vala, valb uint16
}

var inputMap = make(map[string]in)

func main() {
	name := flag.String("i", "input.txt", "input file")
	flag.Parse()

	f, err := os.Open(*name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		parseLine(fs.Text())
	}
	val := processPath("a")
	fmt.Println("A is", val)
	resetVal()
	inputMap["b"] = in{a: strconv.Itoa(val)}
	val = processPath("a")
	fmt.Println("Part 2: A is", val)
}

func resetVal() {
	for k, v := range inputMap {
		v.val = 0
		inputMap[k] = v
	}
}

func processPath(k string) int {
	node := inputMap[k]
	if node.val > 0 {
		return int(node.val)
	}
	node.vala = getValue(node.a)
	if len(node.b) > 0 {
		node.valb = getValue(node.b)
	}
	switch node.op {
	case "RSHIFT":
		node.val = node.vala >> node.valb
	case "LSHIFT":
		node.val = node.vala << node.valb
	case "OR":
		node.val = node.vala | node.valb
	case "AND":
		node.val = node.vala & node.valb
	case "NOT":
		node.val = ^node.vala
	default:
		node.val = node.vala
	}
	inputMap[k] = node
	return int(node.val)
}

func getValue(node string) uint16 {
	val, err := strconv.Atoi(node)
	if err != nil {
		val = processPath(node)
	}
	return uint16(val)
}

func parseLine(l string) {
	tokens := strings.Split(l, " ")
	m := in{}
	m.a = tokens[0]
	switch len(tokens) {
	case 5:
		m.op = tokens[1]
		m.b = tokens[2]
	case 4:
		m.op = tokens[0]
		m.a = tokens[1]
	}
	inputMap[tokens[len(tokens)-1]] = m
}
