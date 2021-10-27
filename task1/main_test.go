package main

import (
	"testing"
)

/*func Test_main(t *testing.T){
	//simple - not testing
	main()
}*/

func Test_drawChessBoard(t *testing.T) {
	var expectations = [][]string{
		{},
		{squareBlack},
		{squareBlack + squareWhite + squareBlack, squareWhite + squareBlack + squareWhite, squareBlack + squareWhite + squareBlack},
	}
	var gotResults = [][]string{
		drawChessBoard(0, 0),
		drawChessBoard(1, 1),
		drawChessBoard(3, 3),
	}
	for i := 0; i < len(expectations); i++ {
		expect := expectations[i]
		got := gotResults[i]

		if len(expect) != len(got) {
			t.Errorf("different expected %d & received %d sizes of slice", len(expect), len(got))
		}

		for j, val := range got {
			if expect[j] != val {
				t.Errorf("Wrong result in string %d expected %v, but got %v", i, expect[i], val)
			}
		}
	}
}

func Test_dataInput(t *testing.T) {
	if !(dataInput()) {
		t.Error("wrong data input")
	}
}
