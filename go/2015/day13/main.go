package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var scores = make(map[string]map[string]int)
var optimal = make(map[string]map[string]int)

func main() {
	name := flag.String("i", "input.txt", "list")
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

	// Add myself
	// for k, _ := range scores {
	// 	scores["me"] = map[string]int{k: 0}
	// }
	findOptimal()
}

type tuple struct {
	name  string
	score int
}

func findOptimal() {
	// Create all permutations of seating arrangements
	init := []string{}
	for p, _ := range scores {
		init = append(init, p)
	}
	init = append(init, "me")
	perm := permute(init, len(init))
	res := make([]int, len(perm))
	for k, v := range perm {
		sum := 0
		for curr := 0; curr < len(v); curr++ {
			next := curr + 1
			if next == len(v) {
				next = 0
			}

			sum += scores[perm[k][curr]][perm[k][next]]
			sum += scores[perm[k][next]][perm[k][curr]]
		}
		res[k] = sum
	}

	fmt.Println(maxInt(res))
}

func permute(arr []string, size int) [][]string {
	r := [][]string{}
	if size == 1 {
		perm := make([]string, len(arr))
		copy(perm, arr)
		r = append(r, perm)
		return r
	}
	for i := 0; i < size; i++ {
		r = append(r, permute(arr, size-1)...)

		if size%2 == 1 {
			arr[0], arr[size-1] = arr[size-1], arr[0]
			continue
		}
		arr[i], arr[size-1] = arr[size-1], arr[i]
	}
	return r
}

func maxInt(arr []int) int {
	max := -99999
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}

func max(m map[string]int) tuple {
	r := tuple{score: -999999}
	for k, v := range m {
		if v > r.score {
			r.name = k
			r.score = v
		}
	}

	return r
}

var re = regexp.MustCompile(`(?P<from>[A-Za-z]+) would (?P<op>(?:gain|lose)) (?P<amt>[0-9]+) happiness units by sitting next to (?P<to>[A-Za-z]+)\.`)

func parseLine(l string) {
	vars := re.FindStringSubmatch(l)

	sc := string(vars[3])
	score, err := strconv.Atoi(sc)
	if err != nil {
		return
	}

	from := string(vars[1])
	to := string(vars[4])

	if string(vars[2]) == "lose" {
		score = score * -1
	}
	_, ok := scores[from]
	if !ok {
		scores[from] = make(map[string]int)
	}
	scores[from][to] = score
}
