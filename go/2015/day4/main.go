package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main() {
	seed := "bgvyzdsv"
	if len(os.Args) > 1 {
		seed = os.Args[1]
	}

	i := int64(0)
	h5 := int64(0)
	for {
		s := strconv.AppendInt([]byte(seed), i, 10)
		sum := md5.Sum(s)
		hash := hex.EncodeToString(sum[:])
		if string(hash[:5]) == "00000" && h5 == 0 {
			h5 = i
		}
		if string(hash[:6]) == "000000" {
			break
		}
		i++
	}
	fmt.Println("5 zero number is:", h5)
	fmt.Println("6 zero number is:", i)
}
