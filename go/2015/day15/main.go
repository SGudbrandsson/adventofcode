package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ingredient struct {
	capacity, durability, flavor, texture, calories, sum int
}

var ingredients = []ingredient{}

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
	fmt.Println("Score", getScore(100))
}

type spAmounts struct {
	a, b, c, d int
}

func getScore(spoons int) int {
	r := 0
	for i := 0; i < spoons; i++ {

		for _, v := range ingredients {

		}
	}
	return r
}

func iterate(ingCnt, total int) {

}

func createCombinations(tot, ingCnt, pos int) []ingredient {
	ret := []int{}
	last := []int{}
	// setup
	for i := 0; i < ingCnt; i++ {
		if i == pos {
			last[i] = tot
			continue
		}
		last[i] = 0
	}
	ret = append(ret, last)
	for j := 0; j < ingCnt; j++ {
		for i := 0; ; i++ {
			if last[pos] == 0 {
				break
			}
		}
	}
	return ret
}

// Chocolate: capacity 0, durability 0, flavor -2, texture 2, calories 8
var re = regexp.MustCompile("^([a-zA-Z]+): capacity (-?[0-9]+), durability (-?[0-9]+), flavor (-?[0-9]+), texture (-?[0-9]+), calories (-?[0-9]+)$")

func parseLine(l string) {
	res := re.FindStringSubmatch(l)
	ingr := ingredient{}
	ingr.capacity, ingr.durability = atoi(res[2]), atoi(res[3])
	ingr.flavor, ingr.texture, ingr.calories = atoi(res[4]), atoi(res[5]), atoi(res[6])
	ingredients = append(ingredients, ingr)
}

func atoi(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}
