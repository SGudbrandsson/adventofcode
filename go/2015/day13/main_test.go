package main

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	set := []string{"a", "b", "c"}
	res := permute(set, len(set))
	if !reflect.DeepEqual(set, []string{"a", "b", "c"}) {
		t.Errorf("Input set results are incorrect. Expected [a b c], got: %s", set)
	}
	exp := [][]string{
		[]string{"a", "b", "c"},
		[]string{"b", "a", "c"},
		[]string{"c", "a", "b"},
		[]string{"a", "c", "b"},
		[]string{"b", "c", "a"},
		[]string{"c", "b", "a"}}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected results incorrect.\nExpected: %v\nGot:%v", exp, res)
	}
}
