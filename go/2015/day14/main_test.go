package main

import (
	"reflect"
	"testing"
)

func TestRace(t *testing.T) {
	res := []bool{}
	exp := []bool{true, true, false, false, false, false, true, true, false}

	for i := 1; i < 10; i++ {
		res = append(res, inRace(i, 2, 4))
	}
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
