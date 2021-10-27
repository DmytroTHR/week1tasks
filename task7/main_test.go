package main

import (
	"reflect"
	"testing"
)

type TestSqrts struct {
	input  int
	result []int
}

func Test_findNeededNumbers(t *testing.T) {
	testcases := []TestSqrts{
		{0, []int{}},
		{1, []int{0}},
		{100, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, test := range testcases {
		num = test.input
		curResult := findNeededNumbers()
		if !reflect.DeepEqual(curResult, test.result) {
			t.Errorf("expected %v but got %v -> for input param num = %v", test.result, curResult, test.input)
		}
	}
}
