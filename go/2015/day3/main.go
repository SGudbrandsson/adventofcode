package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type coords struct {
	x int
	y int
}

var counter = make(map[coords]int)
var santa = coords{0, 0}
var rsanta = coords{0, 0}

func main() {
	file := flag.String("i", "input.txt", "input file")
	flag.Parse()
	count(&santa)
	count(&rsanta)
	m := parseFile(file)
	fmt.Println("Number of moves:", m)
	fmt.Println("Number of houses visited:", len(counter))
}

// I'm choosing to use a bufio reader instead of reading the
// whole file into a byte slice for practice purposes.
// This example is a lot simpler by reading the whole file
// into memory, then ranging over the slice for parsing.
// In the real world, you should always watch out for unknown
// inputs such as large files, large streams or unknowns like that.
func parseFile(file *string) int {
	fh, e := os.Open(*file)
	check(e)
	defer fh.Close()

	moves := 0
	fr := bufio.NewReader(fh)
	flip := false
	for {
		b, err := fr.ReadByte()
		if err != nil {
			if err != io.EOF {
				continue
			}
			break
		}
		mover := &santa
		if flip {
			mover = &rsanta
		}
		if move(&b, mover) {
			moves++
		}
		flip = !flip
	}
	return moves
}

func move(dir *byte, mover *coords) bool {
	switch string(*dir) {
	case "<":
		mover.x--
	case ">":
		mover.x++
	case "^":
		mover.y++
	case "v":
		mover.y--
	default:
		return false
	}
	count(mover)
	return true
}

func count(mover *coords) {
	counter[*mover] += 1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
