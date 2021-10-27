package main

import (
	"reflect"
	"testing"
)

type TestFibo struct{
	begin, end int64
	result []int64
}

func Test_getUnderlyingSequence(t *testing.T){
	testcases := []TestFibo{
		{0, 0, []int64{0}},
		{1, 7, []int64{1, 1, 2, 3, 5}},
		{0, 1, []int64{0, 1, 1}},
	}

	for _, test := range testcases {
		begin, end = test.begin, test.end
		curResult := getUnderlyingSequence()
		if !reflect.DeepEqual(curResult, test.result) {
			t.Errorf("expected %v but got %v -> for input params begin = %v, end = %v", test.result, curResult, test.begin, test.end)
		}
	}	
}