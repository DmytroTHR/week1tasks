package main

import (
	"flag"
	"fmt"
)

const (
	squareBlack = "x"
	squareWhite = "_"
)

var(
	height, width int
)

func main(){
	if dataInput() {
		drawChessBoard()
	} else {
		printInstructions()
	}	
}

func dataInput() bool{
	flag.IntVar(&height, "h", 0, "Board height")
	flag.IntVar(&width, "w", 0, "Board width")
	flag.Parse() 
	successfulInput := (height * width != 0)
	return successfulInput
}

func drawChessBoard() {
	for i:=0; i<height; i++{
		for j:=0; j<width; j++{			
			if j%2==0 && i%2==0 || j%2!=0 && i%2!=0{
				fmt.Printf("%s", squareBlack)
			} else {
				fmt.Printf("%s", squareWhite)
			}
		}
		fmt.Println()
	}
}

func printInstructions(){
	fmt.Println("You should specify both height & width params by passing them to function call")
	fmt.Println("\ttask1 -h=8 -w=8\t//will print 8x8 chessboard")
}