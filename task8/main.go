package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	delimiter    = ", "
	instructions = `You should specify both begin & end of Fibonacci sequence (both >= 0, end > begin) by passing them to function call
		task8 -b=1 -e=20	//will print 'Result is: 1, 1, 2, 3, 5, 8, 13'`
)

var (
	begin, end uint64
)

func main() {
	if !dataInput() {
		fmt.Println(instructions)
		return
	}
	fmt.Print("Result is: ")
	for ind, res := range getUnderlyingSequence() {
		if ind > 0 {
			fmt.Print(delimiter)
		}
		fmt.Printf("%d", res)
	}
	fmt.Println()
}

func getUnderlyingSequence() []uint64 {
	fibo := fibonacci()
	curFibo := fibo()
	resSequence := make([]uint64, 0, 10)
	for {
		if curFibo > end {
			break
		}
		if curFibo >= begin && curFibo <= end {
			resSequence = append(resSequence, curFibo)
		}
		curFibo = fibo()
	}
	return resSequence
}

func fibonacci() func() uint64 {
	var curr, next uint64 = 1, 0
	return func() uint64 {
		curr, next = next, curr+next
		return curr
	}
}

func init() {
	flag.Uint64Var(&begin, "b", 0, "Sequence begin")
	flag.Uint64Var(&end, "e", 0, "Sequence end")
}

func dataInput() bool {
	flag.Parse()
	successfulInput := (begin >= 0 && end >= 0 && begin < end) || (end > math.MaxUint64)
	return successfulInput
}
