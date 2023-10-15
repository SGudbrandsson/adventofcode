package main

import (
	"flag"
	"fmt"
	s "strings"
	"time"
)

func main() {
	start := time.Now()
	pw := flag.String("i", "cqjxxyzz", "Password to rotate")
	flag.Parse()

	fmt.Println("New password is", generate(*pw))
	fmt.Println("Time taken:", time.Since(start))
}

func generate(pw string) string {
	for {
		pw = newPw(pw)
		if isSecure(pw) {
			return pw
		}
	}
}

func newPw(pw string) string {
	// Increment last letter
	l, rotated := increment(pw[len(pw)-1])
	pw = pw[:len(pw)-1]
	if rotated && len(pw) > 0 {
		pw = newPw(pw)
	}
	return pw + string(l)
}

var a = byte(97)
var z = byte(122)

func increment(l byte) (byte, bool) {
	if l == z {
		return a, true
	}
	for {
		l++
		if !s.ContainsAny(string(l), "iol") {
			break
		}
	}
	return l, false
}

func isSecure(pw string) bool {
	if s.ContainsAny(pw, "iol") {
		return false
	}

	hasSeq := false
	pairs := 0
	pairPos := 0
	for i := 1; i < len(pw); i++ {
		// Check for pairs
		if pairPos < i-1 && pw[i] == pw[i-1] {
			pairs++
			pairPos = i
		}
		// Check for sequence
		if i > 1 &&
			pw[i-2]+1 == pw[i-1] &&
			pw[i-1]+1 == pw[i] {
			hasSeq = true
		}
	}

	return hasSeq && pairs >= 2
}
