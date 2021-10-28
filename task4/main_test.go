package main

import (
	"os"
	"testing"
)

func init() {
	_, err := os.Create(testFileName)
	if err != nil {
		panic("cannot init testFile for testing\n" + err.Error())
	}
	err = os.WriteFile(testFileName, []byte(testFileData), 0666)
	if err != nil {
		panic("cannot init data for testFile\n" + err.Error())
	}
}

func Test_getFile(t *testing.T) {
	testcases := map[string]bool{
		testFileName:        true,
		notExistingFileName: false,
	}

	for filenm, res := range testcases {
		_, curRes := getFile(filenm)
		if curRes != res {
			t.Errorf("expected %v but received %v for filename: %v", res, curRes, filenm)
		}
	}
}

func Test_getStringFromFile(t *testing.T) {
	testcases := map[string]bool{
		testFileName:        true,
		notExistingFileName: false,
	}

	for filenm, res := range testcases {
		_, curRes := getStringFromFile(FileTask{filenm})
		if curRes != res {
			t.Errorf("expected %v but received %v for filename: %v", res, curRes, filenm)
		}
	}
}

func Test_countSubstringsInFile(t *testing.T) {
	type resFunc struct {
		filename string
		count    int
		success  bool
	}
	testcases := map[string]resFunc{
		"test":         {testFileName, 3, true},
		"cannot find":  {testFileName, 0, true},
		"no such file": {notExistingFileName, 0, false},
	}

	for substr, res := range testcases {
		curResInt, curResBool := countSubstringsInFile(res.filename, substr)
		if curResInt != res.count || curResBool != res.success {
			t.Errorf("expected %v, %v but received %v, %v for filename: %v", res.count, res.success, curResInt, curResBool, res.filename)
		}
	}
}

func Test_replaceSubstringInFile(t *testing.T) {
	type resFunc struct {
		filename string
		success  bool
	}
	testcases := map[string]resFunc{
		"test":         {testFileName, true},
		"cannot find":  {testFileName, true},
		"no such file": {notExistingFileName, false},
	}

	for substr, res := range testcases {
		newSubstr := "NEW" + substr
		curRes := replaceSubstringInFile(res.filename, substr, newSubstr)
		if curRes != res.success {
			t.Errorf("expected %v but received %v for filename: %v", res.success, curRes, res.filename)
		}
	}
}

func Test_main(t *testing.T) {
	osArgsWere := os.Args
	defer func() { os.Args = osArgsWere }()

	os.Args = []string{"main.go", testFileName, "test"}
	main()
	os.Args = []string{"main.go", notExistingFileName, "test"}
	main()	

	os.Args = []string{"main.go", testFileName, "test", "NEWtest"}
	main()
	os.Args = []string{"main.go", notExistingFileName, "test", "NEWtest"}
	main()

	os.Args = []string{"main.go", testFileName} //not enough params
	main()
}
