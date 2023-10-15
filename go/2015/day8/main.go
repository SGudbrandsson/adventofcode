package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
)

func main() {
	fi := flag.String("i", "input.txt", "The input file")
	flag.Parse()

	fh, err := os.Open(*fi)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		parseLine(fs.Text())
	}
	println("Sum is", sum)
	println("New sum is", newsum)
}

var sum = 0
var newsum = 0

func parseLine(l string) {
	newsum += len(strconv.Quote(l)) - len(l)
	sum += len(l)
	s, err := strconv.Unquote(l)
	if err != nil {
		return
	}
	sum -= len(s)
}
