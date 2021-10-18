package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for{
		fmt.Println("FIRST envelope")
		envelope1, err := initEnvelope()
		if err != nil {				
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("----------------------------")

		fmt.Println("SECOND envelope")
		envelope2, err := initEnvelope()
		if err != nil {				
			fmt.Println(err.Error())
			continue
		}		

		checkIfFits(envelope1, envelope2)

		if !weContinue(){
			break
		}
		fmt.Println("=============================")
	}
}

func initEnvelope() (Envelope, error) {
	var strErr []string

	height, err1 := askForParamFloat("Envelope height:")
	if err1 != nil {
		strErr = append(strErr, "INCORRECT input - height parsing error")
	}
	width, err2 := askForParamFloat("Envelope width:")
	if err2 != nil {
		strErr = append(strErr, "INCORRECT input - width parsing error")
	}

	if err1 == nil && err2 == nil && (height <= 0 || width <= 0) {
		strErr = append(strErr, "INCORRECT input - height & width must be greater than 0")
	}

	if len(strErr) > 0 {
		return Envelope{}, errors.New(strings.Join(strErr, "\n"))
	}

	return Envelope{height, width}, nil
}

func weContinue() bool{
	fmt.Println(">> Continue Yes/No")
	var ans string
	fmt.Scanln(&ans)
	ans = strings.TrimSpace(strings.ToLower(ans))
	if ans == "yes" || ans == "y"{
		return true
	}
	return false
}

func askForParamFloat(strToAsk string) (float64, error) {
	fmt.Println(strToAsk)
	var inp string
	fmt.Scanln(&inp)
	return strconv.ParseFloat(inp, 64)
}

func checkIfFits(envelope1, envelope2 Envelope) {
	fmt.Println()

	switch{
	case envelope1.FitsInto(envelope2):
		fmt.Println("First fits into Second")
	case envelope2.FitsInto(envelope1):
		fmt.Println("Second fits into First")
	default:
		fmt.Println("Neither fits")	
	}	
	
	fmt.Println()
}
