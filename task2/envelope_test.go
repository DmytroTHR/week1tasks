package main

import "testing"

type TestMinMax struct {
	env    Envelope
	result float64
}

type TestFit struct {
	env1, env2 Envelope
	result     bool
}

func Test_maxDimension(t *testing.T) {
	testcases := []TestMinMax{
		{Envelope{10, 10.1}, 10.1},
		{Envelope{0, 0}, 0},
		{Envelope{11, 10}, 11},
	}

	for _, test := range testcases {
		curResult := test.env.maxDimension()
		if curResult != test.result {
			t.Errorf("expected %v but got %v -> for input param %v", test.result, curResult, test.env)
		}
	}
}

func Test_minDimension(t *testing.T) {
	testcases := []TestMinMax{
		{Envelope{10, 10.1}, 10},
		{Envelope{0, 0}, 0},
		{Envelope{11, 10}, 10},
	}

	for _, test := range testcases {
		curResult := test.env.minDimension()
		if curResult != test.result {
			t.Errorf("expected %v but got %v -> for input param %v", test.result, curResult, test.env)
		}
	}
}

func Test_FitsInto(t *testing.T) {
	testcases := []TestFit{
		{Envelope{1, 1}, Envelope{1.1, 1.1}, true},
		{Envelope{10, 5}, Envelope{9.5, 5.5}, false},
		{Envelope{10, 5}, Envelope{9.5, 4.5}, false},
	}

	for _, test := range testcases {
		curResult := test.env1.FitsInto(test.env2)
		if curResult != test.result {
			t.Errorf("expected %v but got %v -> for input params %v as envelope 1 and %v as envelope 2", test.result, curResult, test.env1, test.env2)
		}
	}
}
