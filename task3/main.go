package main

import (
	"fmt"
	"sort"
	"strings"
)

var trianglesArray []Triangle

func main() {
	for {
		triangle, err := TriangleInitManually()
		if err == nil {
			trianglesArray = append(trianglesArray, triangle)
		} else {
			fmt.Println(err.Error())
		}
		if !weContinue(askIfContinue()) {
			break
		}
	}
	sort.Sort(ByArea(trianglesArray))
	outputResult(trianglesArray)
}

func askIfContinue() string {
	var ans string
	fmt.Println(">> Wish to add another triangle Yes/No?")
	fmt.Scanln(&ans)
	return ans
}

func weContinue(answer string) bool {
	answer = strings.TrimSpace(strings.ToLower(answer))
	if answer == "yes" || answer == "y" {
		return true
	}
	return false
}

func outputResult(trianglesArray []Triangle) {
	fmt.Println("=========Triangles list:=============")
	for i, trngl := range trianglesArray {
		fmt.Printf("%d. [%s]: %.3f cm2\n", i+1, trngl.Name, trngl.Area())
	}
}
