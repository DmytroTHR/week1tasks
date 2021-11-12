package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	ErrorMsgTooBig     = "cannot represent - number is too big"
	ErrorMsgNotANumber = "provided string is not a number"
)

var (
	OneToTwenty      map[int]string
	TwentyToNinety   map[int]string
	Hundreds         map[int]string
	PowersOfThousand map[int]ThousandToPows
	ZeroRepresent    string
	MinusRepresent   string
	OneInFemale      string
	TwoInFemale      string
	CurrentLang      = ""
)

const (
	delim = " "
)

type ThousandToPows struct {
	IsFemale     bool
	One          string
	TwoThreeFour string
	Others       string
}

func initStructuresForLanguage(lang string) {
	OneToTwenty = make(map[int]string)
	for i := 1; i < 20; i++ {
		val, err := GetValueByKey(lang, strconv.Itoa(i))
		if err != nil {
			log.Fatalln(err.Error())
		}
		OneToTwenty[i] = val
	}

	TwentyToNinety = make(map[int]string)
	for i := 20; i < 100; i += 10 {
		val, err := GetValueByKey(lang, strconv.Itoa(i))
		if err != nil {
			log.Fatalln(err.Error())
		}
		TwentyToNinety[i] = val
	}

	Hundreds = make(map[int]string)
	for i := 100; i < 1000; i += 100 {
		val, err := GetValueByKey(lang, strconv.Itoa(i))
		if err != nil {
			log.Fatalln(err.Error())
		}
		Hundreds[i] = val
	}

	PowersOfThousand = make(map[int]ThousandToPows)
	for i := 0; i < 7; i++ {
		keyStart := "1000_" + strconv.Itoa(i)
		isFemale, err := GetValueByKey(lang, keyStart+"female")
		if err != nil {
			log.Fatalln(err.Error())
		}
		forOne, err := GetValueByKey(lang, keyStart+"one")
		if err != nil {
			log.Fatalln(err.Error())
		}
		forTwo, err := GetValueByKey(lang, keyStart+"two")
		if err != nil {
			log.Fatalln(err.Error())
		}
		forFive, err := GetValueByKey(lang, keyStart+"five")
		if err != nil {
			log.Fatalln(err.Error())
		}
		PowersOfThousand[i] = ThousandToPows{
			IsFemale:     strings.ToLower(isFemale) == "true",
			One:          forOne,
			TwoThreeFour: forTwo,
			Others:       forFive,
		}
	}

	valZero, err := GetValueByKey(lang, "0")
	if err != nil {
		log.Fatalln(err.Error())
	}
	valMinus, err := GetValueByKey(lang, "-")
	if err != nil {
		log.Fatalln(err.Error())
	}
	valOneFemale, err := GetValueByKey(lang, "1female")
	if err != nil {
		log.Fatalln(err.Error())
	}
	valTwoFemale, err := GetValueByKey(lang, "2female")
	if err != nil {
		log.Fatalln(err.Error())
	}
	ZeroRepresent = valZero
	MinusRepresent = valMinus
	OneInFemale = valOneFemale
	TwoInFemale = valTwoFemale
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
		case *lastDigit == OneToTwenty[1]:
			*lastDigit = OneInFemale
		case *lastDigit == OneToTwenty[2]:
			*lastDigit = TwoInFemale
		}
	}
}

func GetStringRepresentation(number, lang string) (string, error) {

	if CurrentLang != lang {
		initStructuresForLanguage(lang)
		CurrentLang = lang
	}

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
			return ZeroRepresent, nil
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
		result = MinusRepresent + delim + result
	}

	return strings.TrimSuffix(result, delim), nil
}
