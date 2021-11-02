package main

import (
	"os"
	"reflect"
	"testing"
)

const (
	piterFile       = "P.txt"
	moscowFile      = "M.txt"
	ticketFile      = "testtickets.txt"
	notexistingFile = "not.exist"
)

type TicketsCheck struct {
	numArray  []int
	strResult []string
}

func Test_dataInput(t *testing.T) {
	osArgsWere := os.Args
	defer func() { os.Args = osArgsWere }()

	os.Args = []string{"main.go", piterFile, ticketFile}
	_, _, errNil := dataInput()

	os.Args = []string{"main.go", moscowFile}
	_, _, errNotNil := dataInput()

	if errNil != nil || errNotNil == nil {
		t.Errorf("Unexpected error received %v should be nil, %v shouldn't be nil", errNil, errNotNil)
	}

}

func Test_readFileToString(t *testing.T) {
	testcases := map[string]string{
		piterFile:       "Piter",
		notexistingFile: "",
	}

	for k, v := range testcases {
		resStr, _ := readFileToString(k)
		if resStr != v {
			t.Errorf("Expected %v but received %v", v, resStr)
		}
	}
}

func Test_getNumbersFromFile(t *testing.T) {
	testcases := map[string][]int{
		ticketFile:      {123321, 201010, 123456, 555881},
		notexistingFile: nil,
	}

	for k, v := range testcases {
		curRes, _ := getNumbersFromFile(k)
		if !reflect.DeepEqual(curRes, v) {
			t.Errorf("Expected %v but received %v", v, curRes)
		}
	}
}

func Test_chooseAlgorithm(t *testing.T) {
	testcases := map[string]int{
		piter:     AlgPiter,
		moscow:    AlgMoscow,
		piterFile: 0,
	}

	for k, v := range testcases {
		curRes := chooseAlgorithm(k)
		if curRes != v {
			t.Errorf("Expected %v but received %v", v, curRes)
		}

	}
}

func Test_getDigitsArrayFromNumber(t *testing.T) {
	testcases := map[int][]int{
		123456:  {1, 2, 3, 4, 5, 6},
		1234:    {0, 0, 1, 2, 3, 4},
		1234567: {},
	}

	for k, v := range testcases {
		curRes, _ := getDigitsArrayFromNumber(k)
		if !reflect.DeepEqual(curRes, v) {
			t.Errorf("Expected %v but received %v", v, curRes)
		}

	}
}

func Test_addZerosIfNeeded(t *testing.T) {
	testcases := map[int]string{
		123456:  "123456",
		1234:    "001234",
		1234567: "123456",
	}

	for k, v := range testcases {
		curRes := addZerosIfNeeded(k)
		if curRes != v {
			t.Errorf("Expected %v but received %v", v, curRes)
		}

	}
}

func Test_findHappyMoscowNumbers(t *testing.T) {
	testcases := []TicketsCheck{
		{[]int{123456, 111003, 13220}, []string{"111003", "013220"}},
		{[]int{1, 2, 3}, []string{}},
		{[]int{1234567}, []string{}},
	}

	for _, v := range testcases {
		curRes := findHappyMoscowNumbers(v.numArray)
		if !reflect.DeepEqual(curRes, v.strResult) {
			t.Errorf("Expected %v but received %v", v.strResult, curRes)
		}
	}
}

func Test_findHappyPiterNumbers(t *testing.T) {
	testcases := []TicketsCheck{
		{[]int{123456, 111003, 13220, 338558}, []string{"013220", "338558"}},
		{[]int{1, 2, 3}, []string{}},
		{[]int{1234567}, []string{}},
	}

	for _, v := range testcases {
		curRes := findHappyPiterNumbers(v.numArray)
		if !reflect.DeepEqual(curRes, v.strResult) {
			t.Errorf("Expected %v but received %v", v.strResult, curRes)
		}
	}
}

func Test_main(t *testing.T){
	osArgsWere := os.Args
	defer func() { os.Args = osArgsWere }()

	os.Args = []string{"main.go", piterFile, ticketFile}
	main()

	os.Args = []string{"main.go", moscowFile, ticketFile}
	main()
	
	os.Args = []string{"main.go"}
	main()

	os.Args = []string{"main.go", notexistingFile, ticketFile}
	main()

	os.Args = []string{"main.go", moscowFile, notexistingFile}
	main()
}