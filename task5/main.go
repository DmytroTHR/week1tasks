package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	translNumber, err := dataInput()
	if err != nil {
		fmt.Printf("Int parsing error \n %s \n", err)
		return
	}
	translString := GetStringRepresentation(translNumber)
	fmt.Println("Result:", translString)

}

func dataInput() (int64, error) {
	if len(os.Args) != 2 {
		return 0, fmt.Errorf("wrong number of parameters")
	}

	return strconv.ParseInt(os.Args[1], 10, 64)
}
