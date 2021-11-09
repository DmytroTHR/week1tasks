package main

import (
	"errors"
	"flag"
	"fmt"
)

const (
	delimiter    = ", "
	MaxInt64     = 1<<63 - 1
	instructions = `You should specify both begin & end of Fibonacci sequence (both >= 0, end > begin) by passing them to function call
		task8 -b=1 -e=20	//will print 'Result is: 1, 1, 2, 3, 5, 8, 13'`
)

var (
	begin, end int64
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

func getUnderlyingSequence() []int64 {
	fibo := fibonacci()
	curFibo, err := fibo()
	resSequence := make([]int64, 0, 10)
	for {
		if curFibo >= begin && curFibo <= end {
			resSequence = append(resSequence, curFibo)
		} else if curFibo > end || err != nil {
			break
		}
		curFibo, err = fibo()
	}
	return resSequence
}

func fibonacci() func() (int64, error) {
	var curr, next int64 = 1, 0
	return func() (int64, error) {
		if MaxInt64-curr < curr {
			return next, errors.New("overflow error")
		}
		curr, next = next, curr+next

		return curr, nil
	}
}

func init() {
	flag.Int64Var(&begin, "b", 0, "Sequence begin")
	flag.Int64Var(&end, "e", 0, "Sequence end")
}

func dataInput() bool {
	flag.Parse()
	successfulInput := (begin >= 0 && end >= 0 && begin < end)
	return successfulInput
}
