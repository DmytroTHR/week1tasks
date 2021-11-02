package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	osArgsWere := os.Args
	defer func() { os.Args = osArgsWere }()

	os.Args = []string{"main.go", "1234567890"}
	main()

	os.Args = []string{"main.go", "abcd"}
	main()

	os.Args = []string{"main.go"}
	main()

}
