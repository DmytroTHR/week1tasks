package main

import (
	"os"
	"reflect"
	"testing"
)

type TestFibo struct {
	begin, end uint64
	result     []uint64
}

func Test_dataInput(t *testing.T) {
	osArgsWere := os.Args
	defer func() { os.Args = osArgsWere }()

	testcases := map[bool][2]string{
		true:  {"-b=1", "-e=100"},
		false: {"-b=-100", "-e=1"},
	}

	for key, val := range testcases {
		os.Args[1] = val[0]
		os.Args[2] = val[1]
		got := dataInput()
		if got != key {
			t.Errorf("Expected %v, but received %v", key, got)
		}
	}
}

func Test_getUnderlyingSequence(t *testing.T) {
	testcases := []TestFibo{
		{0, 0, []uint64{0}},
		{1, 7, []uint64{1, 1, 2, 3, 5}},
	}

	for _, test := range testcases {
		begin, end = test.begin, test.end
		curResult := getUnderlyingSequence()
		if !reflect.DeepEqual(curResult, test.result) {
			t.Errorf("expected %v but got %v -> for input params begin = %v, end = %v", test.result, curResult, test.begin, test.end)
		}
	}
}

func Test_main(t *testing.T) {
	osArgsWere := os.Args
	defer func() { os.Args = osArgsWere }()

	os.Args[1], os.Args[2] = "-b=1", "-e=100"
	main()

	os.Args[1], os.Args[2] = "-b=-100", "-e=1"
	main()
}
