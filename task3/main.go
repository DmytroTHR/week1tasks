package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var trianglesArray []Triangle
	for {
		triangle, err := TriangleInitManually()
		if err == nil {
			trianglesArray = append(trianglesArray, triangle)
		} else {
			fmt.Println(err.Error())
		}
		if !weContinue() {
			break
		}
	}
	sort.Sort(ByArea(trianglesArray))
	outputResult(trianglesArray)
}

func weContinue() bool {
	fmt.Println(">> Wish to add another triangle Yes/No?")
	var ans string
	fmt.Scanln(&ans)
	ans = strings.TrimSpace(strings.ToLower(ans))
	if ans == "yes" || ans == "y" {
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
