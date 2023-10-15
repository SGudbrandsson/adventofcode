package main

import "fmt"

var input = [20]int{11, 30, 47, 31, 32, 36, 3, 1, 5, 3, 32, 36, 15, 11, 46, 26, 28, 1, 19, 3}
var sol = make(map[int]int)

func main() {
	res := recurse(0, 0, 0, 150)
	fmt.Println("Solution for 150 is", res)
	fmt.Println("Solution for the lowest is", getLowest())
}

func getLowest() int {
	lowest := 100
	score := 0
	for k, v := range sol {
		if k < lowest {
			lowest = k
			score = v
		}
	}
	return score
}

func recurse(it, curr, jugs, goal int) int {
	ret := 0
	for i := it; i < len(input); i++ {
		if input[i]+curr == goal {
			ret++
			sol[jugs]++
		}
		ret += recurse(i+1, input[i]+curr, jugs+1, goal)
	}
	return ret
}
