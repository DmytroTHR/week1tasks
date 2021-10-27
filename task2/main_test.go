package main

import "testing"

type TestContinue struct {
	param  string
	result bool
}

type TestFits struct {
	env1, env2 Envelope
	result     string
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

func Test_resultFits(t *testing.T) {
	testcases := []TestFits{
		{Envelope{10, 5}, Envelope{9.5, 4.5}, SecondIntoFirst},
		{Envelope{1, 1}, Envelope{1.1, 1.1}, FirstIntoSecond},
		{Envelope{10, 5}, Envelope{9.5, 5.5}, Neither},
	}

	for _, test := range testcases {
		curResult := resultFits(test.env1, test.env2)
		if curResult != test.result {
			t.Errorf("expected %v but got %v -> for input params %v, %v", test.result, curResult, test.env1, test.env2)
		}
	}

}
