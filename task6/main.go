package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	moscow = "moscow"
	piter  = "piter"

	NumDigitsInTicket = 6

	AlgMoscow = iota + 1
	AlgPiter
)

func main() {

	paramFile, ticketsFile, err := dataInput()
	if err != nil {
		fmt.Printf("Error:\n %s \n", err)
		return
	}

	strFromParams, err := readFileToString(paramFile)
	if err != nil {
		fmt.Printf("File with algo error:\n %s \n", err)
		return
	}

	numbersArray, err := getNumbersFromFile(ticketsFile)
	if err != nil {
		fmt.Printf("Cannot extract ticket numbers from file %s\nReason: %s\n", ticketsFile, err.Error())
		return
	}

	algoChoice := chooseAlgorithm(strFromParams)

	result := make([]string, 0, len(numbersArray))
	switch algoChoice {
	case AlgMoscow:
		result = findHappyMoscowNumbers(numbersArray)
		fmt.Printf("Happy Moscow tickets count in file %s:\t%d\n", ticketsFile, len(result))
	case AlgPiter:
		result = findHappyPiterNumbers(numbersArray)
		fmt.Printf("Happy Piter tickets count in file %s:\t%d\n", ticketsFile, len(result))
	}
	if len(result) > 0 {
		fmt.Println(result)
	}
}

func dataInput() (string, string, error) {
	if len(os.Args) != 3 {
		return "", "", fmt.Errorf("params <AlgNameFile> <TicketsFile> should be specified")
	}

	return os.Args[1], os.Args[2], nil
}

func readFileToString(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func getNumbersFromFile(fileName string) ([]int, error) {
	fileData, err := readFileToString(fileName)
	if err != nil {
		return nil, err
	}

	ticketPattern := "[0-9]{" + strconv.Itoa(NumDigitsInTicket) + "}"
	reg := regexp.MustCompile(ticketPattern)
	matchingStrings := reg.FindAllString(fileData, -1)
	result := make([]int, 0, len(matchingStrings))

	for _, strTicket := range matchingStrings {
		ticketNum, err := strconv.Atoi(strTicket)
		if err != nil {
			continue
		}
		result = append(result, int(ticketNum))
	}

	return result, nil
}

func chooseAlgorithm(algoName string) int {
	algoName = strings.ToLower(algoName)
	isMoscow := strings.Contains(algoName, moscow)
	isPiter := strings.Contains(algoName, piter)

	switch {
	case isMoscow:
		return AlgMoscow
	case isPiter:
		return AlgPiter
	}

	return 0
}

func findHappyMoscowNumbers(numArray []int) []string {
	result := make([]string, 0, len(numArray))
	for _, num := range numArray {
		digits, err := getDigitsArrayFromNumber(num)
		if err != nil {
			continue
		}

		sumFirst, sumLast := 0, 0
		for i := 0; i < NumDigitsInTicket; i++ {
			if i < NumDigitsInTicket/2 {
				sumFirst += digits[i]
			} else {
				sumLast += digits[i]
			}
		}

		if sumFirst == sumLast {
			result = append(result, addZerosIfNeeded(num))
		}
	}
	return result
}

func findHappyPiterNumbers(numArray []int) []string {
	result := make([]string, 0, len(numArray))
	for _, num := range numArray {
		digits, err := getDigitsArrayFromNumber(num)
		if err != nil {
			continue
		}

		sumOdd, sumEven := 0, 0
		for i := 0; i < NumDigitsInTicket; i++ {
			if digits[i]%2 == 0 {
				sumEven += digits[i]
			} else {
				sumOdd += digits[i]
			}
		}

		if sumOdd == sumEven {
			result = append(result, addZerosIfNeeded(num))
		}
	}
	return result
}

func getDigitsArrayFromNumber(number int) ([]int, error) {
	result := make([]int, NumDigitsInTicket)

	numWas := number

	for i:=NumDigitsInTicket-1; i>=0; i-- {
		digit := number % 10
		result[i] = digit
		number = number / 10
	}

	if number != 0 {
		return nil, fmt.Errorf("number of digits in %d shouldn't be greater than %d", numWas, NumDigitsInTicket)
	}

	return result, nil
}

func addZerosIfNeeded(num int) string {
	strNum := strconv.Itoa(num)
	for {
		if len(strNum) >= NumDigitsInTicket {
			break
		}
		strNum = "0" + strNum
	}

	return strNum
}