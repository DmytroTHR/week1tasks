package main

import (
	"reflect"
	"testing"
)

func init() {
	initStructuresForLanguage("RU")
}

func Test_hundredsToNumber(t *testing.T) {
	testcases := map[int][]string{
		-1:  {},
		0:   {},
		1:   {"один"},
		11:  {"одиннадцать"},
		112: {"сто", "двенадцать"},
		547: {"пятьсот", "сорок", "семь"},
		700: {"семьсот"},
		808: {"восемьсот", "восемь"},
		999: {"девятьсот", "девяносто", "девять"},
	}

	for k, v := range testcases {
		curRes := hundredsToNumber(k)
		if !reflect.DeepEqual(curRes, v) {
			t.Errorf("Expected %v but got %v for parameter %v", v, curRes, k)
		}
	}
}

func Test_powThousandSuffix(t *testing.T) {
	testcases := []struct {
		twoDigits int
		powThous  ThousandToPows
		result    string
	}{
		{0, PowersOfThousand[0], ""},
		{1, PowersOfThousand[1], "тысяча"},
		{11, PowersOfThousand[2], "миллионов"},
		{4, PowersOfThousand[3], "миллиарда"},
		{21, PowersOfThousand[4], "триллион"},
		{22, PowersOfThousand[5], "квадриллиона"},
		{99, PowersOfThousand[6], "квинтиллионов"},
	}

	for _, val := range testcases {
		curRes := powThousandSuffix(val.twoDigits, val.powThous)
		if curRes != val.result {
			t.Errorf("Expected %v but got %v", val.result, curRes)
		}
	}
}

func Test_correctGender(t *testing.T) {
	one, two, three := "один", "два", "три"
	testcases := []struct {
		lastDig  *string
		powThous ThousandToPows
		result   string
	}{
		{&one, PowersOfThousand[1], "одна"},
		{&two, PowersOfThousand[1], "две"},
		{&three, PowersOfThousand[1], "три"},
	}

	for _, val := range testcases {
		correctGender(val.lastDig, val.powThous)
		if *val.lastDig != val.result {
			t.Errorf("Expected %v but got %v", val.result, *val.lastDig)
		}
	}
}

func Test_GetStringRepresentation(t *testing.T) {
	testcases := map[string]string{
		"-1000":   "МИНУС" + delim + "одна" + delim + "тысяча",
		"0":       "ноль",
		"1000":    "одна" + delim + "тысяча",
		"1001001": "один" + delim + "миллион" + delim + "одна" + delim + "тысяча" + delim + "один",
		"abc":     "",
	}

	for k, v := range testcases {
		curRes, _ := GetStringRepresentation(k, "RU")
		if curRes != v {
			t.Errorf("Expected %v but got %v for initial value %v", v, curRes, k)
		}
	}
}
