package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	r "regexp"
	"sort"
	sc "strconv"
	s "strings"
)

type X interface{}

type box struct {
	h, w, l int
}

func main() {
	// Read the input file
	file := flag.String("i", "input.txt", "input file name")
	flag.Parse()

	f, e := os.Open(*file)
	check(e)
	defer f.Close()

	filescanner := bufio.NewScanner(f)
	filescanner.Split(bufio.ScanLines)

	totalPaper := 0
	totalRibbon := 0
	lines := 0

	for filescanner.Scan() {
		b, e := processLine(filescanner.Text())
		lines++
		if e != nil {
			continue
		}
		paper, ribbon := calcReq(b)
		totalPaper += paper
		totalRibbon += ribbon
	}
	fmt.Println("Total paper area is:", totalPaper, "sqft")
	fmt.Println("Total ribbon needed is:", totalRibbon, "ft")
	fmt.Println("Total lines:", lines)
}

func calcReq(b *box) (int, int) {
	bi := []int{b.h, b.w, b.l}
	bs := []int{b.h * b.w, b.h * b.l, b.w * b.l}
	sort.Ints(bs)
	sort.Ints(bi)
	sum := 0
	for _, s := range bs {
		sum += s
	}
	ribbon := (bi[0]+bi[1])*2 + bi[0]*bi[1]*bi[2]
	// Return a full box plus smallest side, and the ribbon length
	return (2 * sum) + bs[0], ribbon
}

func processLine(line string) (*box, error) {
	m, _ := r.MatchString("^[0-9]+x[0-9]+x[0-9]+$", line)
	if !m {
		return nil, errors.New("Invalid input")
	}
	sp := s.Split(line, "x")
	ret := box{h: atoi(sp[0]), w: atoi(sp[1]), l: atoi(sp[2])}
	return &ret, nil
}

func atoi(s string) int {
	ret, e := sc.Atoi(s)
	if e != nil {
		return 0
	}
	return ret
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
