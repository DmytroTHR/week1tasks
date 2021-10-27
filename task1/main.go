package main

import (
	"flag"
	"fmt"
)

const (
	squareBlack  = "x"
	squareWhite  = "_"
	instructions = `You should specify both height & width params (>0) by passing them to function call")
	task1 -h=8 -w=8 //will print 8x8 chessboard`
)

var (
	height, width int
)

func main() {
	if !dataInput() {
		fmt.Println(instructions)
		return
	}

	board := drawChessBoard(height, width)
	for _, val := range board {
		fmt.Println(val)
	}
}

func dataInput() bool {
	flag.IntVar(&height, "h", 1, "Board height")
	flag.IntVar(&width, "w", 1, "Board width")
	flag.Parse()
	successfulInput := (height > 0 && width > 0)
	return successfulInput
}

func drawChessBoard(height, width int) []string {
	result := make([]string, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if j%2 == 0 && i%2 == 0 || j%2 != 0 && i%2 != 0 {
				result[i] = result[i] + squareBlack
			} else {
				result[i] = result[i] + squareWhite
			}
		}
	}
	return result
}
