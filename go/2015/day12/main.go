package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type data struct {
	data []interface{}
}

func main() {
	fn := flag.String("i", "input.txt", "Input string")
	flag.Parse()

	data, err := os.ReadFile(*fn)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("-?[0-9]+")
	nums := re.FindAllString(string(data), -1)
	if nums == nil {
		fmt.Println("Oh no .. no numbers")
		os.Exit(1)
	}

	parseJson(data)
	fmt.Println("Sum is", sumNums(nums))
}

func parseJson(j []byte) {
	data := map[string]any{}
	json.Unmarshal(j, &data)
	nums := validateData(data, false)
	sum := float64(0)
	for _, v := range nums {
		sum += v
	}
	fmt.Println("Value 2 is:", sum)
}

func validateData(data map[string]any, isArr bool) []float64 {
	ret := make([]float64, 1000)
	skip := false
	for k, v := range data {
		switch vv := v.(type) {
		case map[string]interface{}:
			d := validateData(vv, false)
			ret = append(ret, d...)
		case []interface{}:
			for i := 0; i < len(vv); i++ {
				val := map[string]interface{}{"a": vv[i]}
				d := validateData(val, true)
				ret = append(ret, d...)
			}
		case string:
			if vv == "red" && !isArr {
				skip = true
				break
			}
		case float64:
			ret = append(ret, vv)
		default:
			fmt.Println(k)
		}
	}
	if skip {
		return nil
	}
	return ret
}

func sumNums(nums []string) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		n, err := strconv.Atoi(nums[i])
		if err != nil {
			continue
		}
		sum += n
	}
	return sum
}
