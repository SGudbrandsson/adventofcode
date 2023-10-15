package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	sc "strconv"
	s "strings"
)

type point struct {
	x, y int
}

type op struct {
	instruction string
	from, to    point
}

var grid = [1000][1000]bool{}
var grid2 = [1000][1000]int{}

func main() {
	fn := flag.String("i", "input.txt", "input file for day 6 2015")
	flag.Parse()

	fh, err := os.Open(*fn)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	fs := bufio.NewScanner(fh)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		op := getOperation(fs.Text())
		operateGrid(&op)
	}

	measureLights()
}

func measureLights() {
	c := 0
	i := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] {
				c++
			}
			i += grid2[x][y]
		}
	}
	fmt.Println("Light count:", c)
	fmt.Println("Light intensity:", i)
}

func operateGrid(op *op) {
	for x := op.from.x; x <= op.to.x; x++ {
		for y := op.from.y; y <= op.to.y; y++ {
			operateLight(&op.instruction, x, y)
		}
	}
}

func operateLight(a *string, x int, y int) {
	switch *a {
	case "on":
		grid[x][y] = true
		grid2[x][y] += 1
	case "off":
		grid[x][y] = false
		if grid2[x][y] > 0 {
			grid2[x][y] -= 1
		}
	case "toggle":
		grid[x][y] = !grid[x][y]
		grid2[x][y] += 2
	}
}

func getOperation(l string) op {
	getPoint(&l, 0)
	return op{
		instruction: getInstructions(&l),
		from:        getPoint(&l, 0),
		to:          getPoint(&l, 1),
	}
}

func getPoint(l *string, i int) point {
	re := regexp.MustCompile("([0-9]+),([0-9]+)")
	d := re.FindAllStringSubmatch(*l, 2)
	return point{
		x: atoi(d[i][1]),
		y: atoi(d[i][2]),
	}
}

func atoi(s string) int {
	ret, e := sc.Atoi(s)
	if e != nil {
		return 0
	}
	return ret
}

func getInstructions(l *string) string {
	if s.Contains(*l, "turn on") {
		return "on"
	}
	if s.Contains(*l, "turn off") {
		return "off"
	}
	if s.Contains(*l, "toggle") {
		return "toggle"
	}
	return ""
}
