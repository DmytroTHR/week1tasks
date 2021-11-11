package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	ErrorMsgTooBig     = "cannot represent - number is too big"
	ErrorMsgNotANumber = "provided string is not a number"
)

var (
	OneToTwenty = map[int]string{
		1:  "один",
		2:  "два",
		3:  "три",
		4:  "четыре",
		5:  "пять",
		6:  "шесть",
		7:  "семь",
		8:  "восемь",
		9:  "девять",
		10: "десять",
		11: "одиннадцать",
		12: "двенадцать",
		13: "тринадцать",
		14: "четырнадцать",
		15: "пятнадцать",
		16: "шестнадцать",
		17: "семнадцать",
		18: "восемнадцать",
		19: "девятнадцать",
	}

	TwentyToNinety = map[int]string{
		20: "двадцать",
		30: "тридцать",
		40: "сорок",
		50: "пятьдесят",
		60: "шестьдесят",
		70: "семьдесят",
		80: "восемьдесят",
		90: "девяносто",
	}

	Hundreds = map[int]string{
		100: "сто",
		200: "двести",
		300: "триста",
		400: "четыреста",
		500: "пятьсот",
		600: "шестьсот",
		700: "семьсот",
		800: "восемьсот",
		900: "девятьсот",
	}

	PowersOfThousand = map[int]ThousandToPows{
		0: {false, "", "", ""},
		1: {true, "тысяча", "тысячи", "тысяч"},
		2: {false, "миллион", "миллиона", "миллионов"},
		3: {false, "миллиард", "миллиарда", "миллиардов"},
		4: {false, "триллион", "триллиона", "триллионов"},
		5: {false, "квадриллион", "квадриллиона", "квадриллионов"},
		6: {false, "квинтиллион", "квинтиллиона", "квинтиллионов"},
	}
)

const (
	delim          = " "
	minusRepresent = "МИНУС"
	zeroRepresent  = "ноль"
)

type ThousandToPows struct {
	IsFemale     bool
	One          string
	TwoThreeFour string
	Others       string
}

func hundredsToNumber(hundreds int) []string {
	result := make([]string, 0, 3)

	ones := hundreds % 10
	tens := (hundreds - ones) % 100
	hunds := (hundreds - tens - ones)

	if hunds > 0 {
		result = append(result, Hundreds[hunds])
	}

	if tens > 10 {
		result = append(result, TwentyToNinety[tens])
	}

	if ones > 0 || tens == 10 {
		if tens == 10 {
			result = append(result, OneToTwenty[tens+ones])
		} else {
			result = append(result, OneToTwenty[ones])
		}
	}

	return result
}

func powThousandSuffix(twoLastDigits int, tenToPow ThousandToPows) string {
	firstDigit := twoLastDigits / 10
	lastDigit := twoLastDigits % 10
	switch {
	case lastDigit == 1 && firstDigit != 1:
		return tenToPow.One
	case (lastDigit == 2 || lastDigit == 3 || lastDigit == 4) && firstDigit != 1:
		return tenToPow.TwoThreeFour
	default:
		return tenToPow.Others
	}
}

func correctGender(lastDigit *string, thousandToPow ThousandToPows) {
	if thousandToPow.IsFemale {
		switch {
		case *lastDigit == "один":
			*lastDigit = "одна"
		case *lastDigit == "два":
			*lastDigit = "две"
		}
	}
}

func GetStringRepresentation(number string) (string, error) {

	regex, err := regexp.Compile("-?[0-9]+")
	if err != nil {
		return "", fmt.Errorf(ErrorMsgNotANumber)
	}
	matching := regex.FindString(number)
	if matching != number {
		return "", fmt.Errorf(ErrorMsgNotANumber)
	}

	regex, err = regexp.Compile("-?0+")
	if err == nil {
		matching = regex.FindString(number)
		if matching == number {
			return zeroRepresent, nil
		}
	}

	var fullRepresent []string

	var isNegative bool
	if number[0] == '-' {
		isNegative = true
		number = number[1:]
	}

	var powOfThous int
	for len(number) > 0 {
		if powOfThous >= len(PowersOfThousand) {
			return "", fmt.Errorf(ErrorMsgTooBig)
		}

		idx := len(number) - 3
		if idx < 0 {
			idx = 0
		}
		lastThreeDigits := number[idx:]
		curNumber, err := strconv.Atoi(lastThreeDigits)
		if err != nil {
			return "", fmt.Errorf(ErrorMsgNotANumber)
		}

		if curNumber > 0 {
			curRepresent := hundredsToNumber(curNumber)
			correctGender(&curRepresent[len(curRepresent)-1], PowersOfThousand[powOfThous])
			curRepresent = append(curRepresent, powThousandSuffix(curNumber%100, PowersOfThousand[powOfThous]))
			fullRepresent = append(curRepresent, fullRepresent...)
		}

		number = number[:idx]
		powOfThous++
	}

	result := strings.Join(fullRepresent, delim)
	if isNegative {
		result = minusRepresent + delim + result
	}

	return strings.TrimSuffix(result, delim), nil
}
