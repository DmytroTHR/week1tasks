package main

import (
	"fmt"
	"strings"
)

const (
	FirstIntoSecond = "First fits into Second"
	SecondIntoFirst = "Second fits into First"
	Neither         = "Neither fits"
)

func main() {
	for {
		fmt.Println("FIRST envelope")
		envelope1, err := EnvelopeInitManually()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("----------------------------")

		fmt.Println("SECOND envelope")
		envelope2, err := EnvelopeInitManually()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf("\n%v\n", resultFits(envelope1, envelope2))

		if !weContinue(askIfContinue()) {
			break
		}
		fmt.Println("=============================")
	}
}

func askIfContinue() string {
	var ans string
	fmt.Println(">> Continue Yes/No")
	fmt.Scanln(&ans)
	return ans
}

func weContinue(answer string) bool {
	answer = strings.TrimSpace(strings.ToLower(answer))
	if answer == "yes" || answer == "y" {
		return true
	}
	return false
}

func resultFits(envelope1, envelope2 Envelope) string {
	var result string
	switch {
	case envelope1.FitsInto(envelope2):
		result = FirstIntoSecond
	case envelope2.FitsInto(envelope1):
		result = SecondIntoFirst
	default:
		result = Neither
	}

	return result
}
