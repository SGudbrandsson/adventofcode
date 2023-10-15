package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	i := 0
	for fs.Scan() {
		parseLine(fs.Text(), i)
		i++
	}
	sticky()

	for i := 0; i < 100; i++ {
		processLights()
	}
	fmt.Println("Number of lights on are", countLights())
}

var grid = [100][100]bool{}
var old = [100][100]bool{}

func countLights() int {
	on := 0
	for _, v := range grid {
		for _, v := range v {
			if v {
				on++
			}
		}
	}
	return on
}

func processLights() {
	old = grid
	for k1, v := range grid {
		for k2, _ := range v {
			on := getNeighb(k1, k2)
			if grid[k1][k2] {
				if on == 2 || on == 3 {
					continue
				}
				grid[k1][k2] = false
				continue
			}
			if on == 3 {
				grid[k1][k2] = true
			}
		}
	}
	sticky()
}

func sticky() {
	grid[0][0] = true
	grid[0][99] = true
	grid[99][0] = true
	grid[99][99] = true
}

func getNeighb(k1, k2 int) int {
	on := 0
	for i := k1 - 1; i <= k1+1; i++ {
		if i > 99 || i < 0 {
			continue
		}
		for j := k2 - 1; j <= k2+1; j++ {
			if j > 99 || j < 0 {
				continue
			}
			if i == k1 && j == k2 {
				continue
			}
			if old[i][j] {
				on++
			}
		}
	}
	return on
}

func parseLine(l string, i int) {
	for k, v := range l {
		if v == '#' {
			grid[i][k] = true
			continue
		}
		grid[i][k] = false
	}
}
