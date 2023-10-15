package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func main() {
	input := "1113122113"
	start := time.Now()
	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
		if i == 39 {
			fmt.Println("Length is", len(input))
		}
	}
	fmt.Println("Length is", len(input))
	fmt.Println("Time taken:", time.Since(start))
	// fmt.Println("String is", input)
}

func lookAndSay(s string) string {
	buff := bytes.NewBufferString("")
	buff.Grow(len(s) * 14 / 10)
	count := 0
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] != s[i-1] {
			buff.WriteString(strconv.Itoa(count))
			buff.WriteByte(s[i-1])
			count = 1
			continue
		}
		count++
	}
	buff.WriteString(strconv.Itoa(count))
	buff.WriteByte(s[len(s)-1])
	return buff.String()
}
