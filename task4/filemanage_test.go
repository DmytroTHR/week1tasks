package main

import (
	"os"
	"testing"
)

const (
	testFileName        = "testFile.txt"
	notExistingFileName = "notexistingfilename.notexist"
	testFileData        = `Some testing text for our test
		Our tests are the best.`
)

var (
	testFileTask = FileTask{testFileName}
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

func Test_MakeFile(t *testing.T) {
	_, result := MakeFile(testFileName)
	if !result {
		t.Errorf("cannot make file from %v", testFileName)
	}
}

func Test_Exists(t *testing.T) {
	testcases := map[string]bool{
		testFileName:        true,
		notExistingFileName: false,
	}

	for filenm, res := range testcases {
		curRes := FileTask{filenm}.Exists()
		if curRes != res {
			t.Errorf("expected %v but received %v for filename: %v", res, curRes, filenm)
		}
	}
}

func Test_ReadToString(t *testing.T) {
	testcases := map[string]string{
		testFileName:        testFileData,
		notExistingFileName: "",
	}

	for filenm, res := range testcases {
		curRes, _ := FileTask{filenm}.ReadToString()
		if curRes != res {
			t.Errorf("expected %v but received %v for filename: %v", res, curRes, filenm)
		}
	}
}

func Test_RewriteWithString(t *testing.T) {
	err := testFileTask.RewriteWithString(testFileData)
	if err != nil {
		t.Errorf("failed to rewrite: %v", err.Error())
	}
}
