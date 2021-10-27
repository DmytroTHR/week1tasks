package main

import (
	"testing"
)

type TestContinue struct {
	param  string
	result bool
}

func Test_weContinue(t *testing.T) {
	testcases := []TestContinue{
		{"yes", true},
		{"y", true},
		{"YES ", true},
		{" Y", true},
		{"", false},
		{"No", false},
	}

	for _, test := range testcases {
		curResult := weContinue(test.param)
		if curResult != test.result {
			t.Errorf("expected %v but got %v -> for input param %v", test.result, curResult, test.param)
		}
	}
}

func Test_main(t *testing.T) {
	trianglesArray = append(trianglesArray, Triangle{2, 3, 4, "tr1"})
	main()
}
