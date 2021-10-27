package main

import (
	"testing"
)

/*func Test_main(t *testing.T){
	//simple - not testing
	main()
}*/

func Test_drawChessBoard(t *testing.T) {
	var expect = []string{"x_x", "_x_", "x_x"}
	got := drawChessBoard(3, 3)
	if len(expect) != len(got){
		t.Errorf("different expected %d & received %d sizes of slice", len(expect), len(got))
	}
	for i, val := range got {
		if expect[i] != val {
			t.Errorf("Wrong result in string %d expected %v, but got %v", i, expect[i], val)
		}
	}

}

func Test_dataInput(t *testing.T){
	if ! (dataInput()){
		t.Error("wrong data input")
	}
}