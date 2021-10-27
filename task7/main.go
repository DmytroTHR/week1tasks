package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	num int
)

const (
	delimiter    = ", "
	instructions = `You should specify natural number(>=0) to analyze by passing it to function call
		task7 -n 225	//will print 'Result is: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14'`
)

func main() {
	if !dataInput() {
		fmt.Println(instructions)
		return
	}

	fmt.Print("Result is: ")
	for ind, res := range findNeededNumbers() {
		if ind > 0 {
			fmt.Print(delimiter)
		}
		fmt.Printf("%d", res)
	}
	fmt.Println()
}

func findNeededNumbers() []int {
	bound := int(math.Sqrt(float64(num)))
	if bound*bound >= num {
		bound--
	}
	result := make([]int, bound+1) //+1 is for 0 value
	for i := 0; i <= bound; i++ {
		result[i] = i
	}
	return result
}

func dataInput() bool {
	flag.IntVar(&num, "n", 0, "Number to analyze")
	flag.Parse()
	successfulInput := (num >= 0)
	return successfulInput
}

