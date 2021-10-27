package main

import (
	"errors"
	"math"
	"testing"
)

type TestCanBe struct {
	a, b, c float64
	result  bool
}

type TestParse struct {
	input            string
	resA, resB, resC float64
	resName          string
	resErr           error
}

type TestArea struct {
	triangle Triangle
	result   float64
}

func Test_canBeTriangle(t *testing.T) {
	testcases := []TestCanBe{
		{2, 3, 4, true},
		{1, 2, 3, false},
	}

	for _, test := range testcases {
		curResult := canBeTriangle(test.a, test.b, test.c)
		if curResult != test.result {
			t.Errorf("expected %v but got %v -> for input params num = %v, %v, %v", test.result, curResult, test.a, test.b, test.c)
		}
	}
}

func Test_parseTriangleInput(t *testing.T) {
	testcases := []TestParse{
		{"tr1,2,3,4", 2.0, 3.0, 4.0, "tr1", nil},
		{"2,3,4", 0.0, 0.0, 0.0, "", errors.New(NotEnoughParams)},
		{"tr1,2,3,ab", 0.0, 0.0, 0.0, "", errors.New(CannotParse)},
	}

	for _, test := range testcases {
		curName, curA, curB, curC, curErr := parseTriangleInput(test.input)
		if curA != test.resA || curB != test.resB || curC != test.resC || curName != test.resName || (curErr != test.resErr && curErr.Error() != test.resErr.Error()) {
			t.Errorf("parsing error on params %v", test)
		}
	}
}

func Test_trimTabsAndSpaces(t *testing.T) {
	testcases := [][2]string{
		{" ", ""},
		{"\t", ""},
		{"\t  \t", ""},
		{" \tabc \t", "abc"},
	}

	for _, test := range testcases {
		curResult := trimTabsAndSpaces(test[0])
		if curResult != test[1] {
			t.Errorf("expected <%v> but received <%v> -> for param %v", test[1], curResult, test[0])
		}
	}
}

func Test_Area(t *testing.T) {
	testcases := []TestArea{
		{Triangle{2, 3, 4, "tr1"}, 2.905},
		{Triangle{11, 5, 13, "tr1"}, 26.892},
	}

	for _, test := range testcases{
		curResult := test.triangle.Area()
		if math.Abs(curResult - test.result) >= 0.001 {
			t.Errorf("expected %v but received %v -> for triangle %v", test.result, curResult, test.triangle)
		}
	}
}

func Test_TriangleInitManually(t *testing.T) {
	TriangleInitManually()
}
