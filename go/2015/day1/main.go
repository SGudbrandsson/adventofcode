package main

import (
	"fmt"
	"os"
	s "strings"
)

func main() {
	args := get_args()
	fmt.Println("The floor is:", sumbrackets(args))
	fmt.Println("Santa enters the basement at:", find_basement(args))
}

func find_basement(input string) int {
	floor := 0
	for i := 0; i < len(input); i++ {
		if string(input[i]) == "(" {
			floor++
		}
		if string(input[i]) == ")" {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return 0
}

func sumbrackets(input string) int {
	return s.Count(input, "(") - s.Count(input, ")")
}

func get_args() string {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	arg := os.Args[1]

	if !s.Contains(arg, "(") && !s.Contains(arg, ")") {
		usage()
		os.Exit(1)
	}
	return arg
}

func usage() {
	fmt.Println("Usage: include ( or ) in the input arguments")
}
