package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type conn struct {
	node   *node
	length uint32
}

type node struct {
	name string
	conn []conn
}

type item struct {
	name  string
	dist  uint32
	count int
}

type path struct {
	from, to string
	dist     uint32
}

type group struct {
	places []string
}

var graph = make(map[string]node)
var shortest = []uint32{}
var visited = make(map[string]bool)
var queue = []item{}
var pq = []path{}
var cost = make(map[string]map[string]uint32)
var paths = []path{}

func main() {
	fi := flag.String("i", "input.txt", "Input file")
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

	for k, _ := range graph {
		// Method 1 - trying to implement Djikstras Shortest path algorithm
		setup(k)
		queue = append(queue, item{k, 0, 0})
		for {
			walk()
			queue = queue[1:]
			if len(queue) == 0 {
				break
			}
		}
		fmt.Println("Shortest length is", shortest[len(shortest)-1])
	}

	// Method 2 - pick a starting node, then brute force paths
	// (assuming all connect together in a mesh - doesn't work in a proper graph)
	for name, _ := range cost {
		walk2(name, []string{}, 0)
	}
	fmt.Println("Shortest2 length is", shortest2)

	// Method 3 - Sort by shortest distance and create groups, then connect them
	getShort3()

	// Next up - assume no backtracking, and generate all permutations.
	getShort4()
}

func getShort4() {
	names := []string{}
	for k, _ := range cost {
		names = append(names, k)
	}
	perm := permute(names, len(names))

	dist := uint32(4_294_967_295)
	long := uint32(0)
	for i := 0; i < len(perm); i++ {
		sum := uint32(0)
		for j := 0; j < len(perm[i])-1; j++ {
			sum += cost[perm[i][j]][perm[i][j+1]]
		}
		if sum < dist {
			dist = sum
		}
		if sum > long {
			long = sum
		}
	}
	fmt.Println("Shortest Distance 4 is:", dist)
	fmt.Println("Longest Distance 4 is:", long)
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

func getShort3() {
	// Sort by distance
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].dist < paths[j].dist
	})

	visited := make(map[string]bool, len(cost))
	groups := []group{}
	m := []path{}
	dist := uint32(0)
	for _, v := range paths {
		gf := findinGroup(v.from, groups)
		gt := findinGroup(v.to, groups)
		if visited[v.from] && visited[v.to] && gf == gt {
			continue
		}
		dist += v.dist
		m = append(m, v)
		if !visited[v.from] && !visited[v.to] {
			groups = append(groups, group{[]string{v.from, v.to}})
			visited[v.from] = true
			visited[v.to] = true
			continue
		}
		if !visited[v.from] {
			groups[gt].places = append(groups[gt].places, v.from)
			visited[v.from] = true
			continue
		}
		if !visited[v.to] {
			groups[gf].places = append(groups[gf].places, v.to)
			visited[v.to] = true
			continue
		}
		// Merge groups
		groups[gf].places = append(groups[gf].places, groups[gt].places...)
		groups[gt] = groups[len(groups)-1]
		groups = groups[:len(groups)-1]
	}
	fmt.Println("Shortest 3 length is", dist)
	// Not correct - we only have a connecting map of shortest distances, but we need to walk back and forth to visit all nodes.
}

func findinGroup(needle string, groups []group) int {
	for k, v := range groups {
		for _, v := range v.places {
			if v == needle {
				return k
			}
		}
	}
	return -1
}

var shortest2 = uint32(10000000)

func walk2(n string, visited []string, d uint32) {
	visited = append(visited, n)
	for k, v := range cost[n] {
		if contains(visited, k) {
			continue
		}
		d += v
		walk2(k, visited, d)
	}
	if len(visited) != len(cost) {
		return
	}
	if shortest2 > d {
		shortest2 = d
	}
}

func contains(stack []string, n string) bool {
	for _, v := range stack {
		if v == n {
			return true
		}
	}
	return false
}

func setup(k string) {
	shortest = make([]uint32, len(graph))
	i := 0
	for k, _ := range graph {
		shortest[i] = 4_294_967_295
		visited[k] = false
		i++
	}
	shortest[0] = 0
	visited[k] = true
}

func walk() {
	step := queue[0]
	if step.count == len(shortest) {
		return
	}
	for _, v := range graph[step.name].conn {
		if visited[v.node.name] {
			continue
		}
		if shortest[step.count] > step.dist+v.length {
			shortest[step.count] = step.dist + v.length
		}
		queue = append(queue, item{v.node.name, step.dist + v.length, step.count + 1})
	}
}

func firstKey() string {
	r := ""
	for _, v := range graph {
		r = v.name
		break
	}
	return r
}

func parseLine(l string) {
	t1 := strings.Split(l, " = ")
	t2 := strings.Split(t1[0], " to ")
	n1 := getNode(t2[0])
	n2 := getNode(t2[1])
	dist, err := strconv.Atoi(t1[1])
	if err != nil {
		panic(err)
	}
	n1.conn = append(n1.conn, conn{node: n2, length: uint32(dist)})
	n2.conn = append(n2.conn, conn{node: n1, length: uint32(dist)})
	graph[t2[0]] = *n1
	graph[t2[1]] = *n2
	if _, ok := cost[n1.name]; !ok {
		cost[n1.name] = make(map[string]uint32)
	}
	cost[n1.name][n2.name] = uint32(dist)
	if _, ok := cost[n2.name]; !ok {
		cost[n2.name] = make(map[string]uint32)
	}
	cost[n2.name][n1.name] = uint32(dist)
	pq = append(pq, path{from: n1.name, to: n2.name})
	paths = append(paths, path{from: n1.name, to: n2.name, dist: uint32(dist)})
}

func getNode(s string) *node {
	n, ok := graph[s]
	if !ok {
		n = node{name: s}
	}
	return &n
}
