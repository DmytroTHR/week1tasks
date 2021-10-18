package main

import (
	"fmt"
	"strings"
)

func main() {
	for{
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

		checkIfFits(envelope1, envelope2)

		if !weContinue(){
			break
		}
		fmt.Println("=============================")
	}
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
