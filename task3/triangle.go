package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Triangle struct {
	sideA, sideB, sideC float64
	Name                string
}

type GeronArea interface {
	Area() float64
}

func (triangle *Triangle) Area() float64 {
	halfPerimeter := (triangle.sideA + triangle.sideB + triangle.sideC) / 2.0
	return math.Sqrt(halfPerimeter * (halfPerimeter - triangle.sideA) * (halfPerimeter - triangle.sideB) * (halfPerimeter - triangle.sideC))
}

func TriangleInitManually() (Triangle, error) {
	fmt.Println("Input triangle info: <name>, <sideA>, <sideB>, <sideC>")

	inputResult := readUserInput()

	name, sideA, sideB, sideC, err := parseTriangleInput(inputResult)
	if err != nil {
		return Triangle{}, errors.New("unable to parse triangle info from: " + inputResult + "\n" + err.Error())
	}

	if !canBeTriangle(sideA, sideB, sideC) {
		return Triangle{}, errors.New("Triangle cannot be built: " + inputResult)
	}

	return Triangle{sideA, sideB, sideC, name}, nil
}

func canBeTriangle(sideA, sideB, sideC float64) bool {
	return sideA+sideB > sideC && sideB+sideC > sideA && sideC+sideA > sideB
}

func readUserInput() string {
	var inputResult string
	bufScanner := bufio.NewScanner(os.Stdin)
	if bufScanner.Scan() {
		inputResult = bufScanner.Text()
	}
	return inputResult
}

func parseTriangleInput(userInput string) (string, float64, float64, float64, error) {
	foundParams := strings.Split(userInput, ",")
	if len(foundParams) != 4 {
		return "", 0, 0, 0, errors.New("number of received parameters doesn't equal to 4")
	}
	name := trimTabsAndSpaces(foundParams[0])
	sideA, okA := parseToFloat(foundParams[1])
	sideB, okB := parseToFloat(foundParams[2])
	sideC, okC := parseToFloat(foundParams[3])

	if !okA || !okB || !okC {
		return "", 0, 0, 0, errors.New("cannot parse one ore more side sizes")
	}

	return name, sideA, sideB, sideC, nil
}

func trimTabsAndSpaces(str string) string {
	res := strings.Trim(str, " ")
	res = strings.Trim(res, "\t")
	return res
}

func parseToFloat(str string) (float64, bool) {
	res, err := strconv.ParseFloat(trimTabsAndSpaces(str), 64)
	if err != nil {
		return 0, false
	}
	return res, true
}

type ByArea []Triangle

func (trArr ByArea) Len() int {
	return len(trArr)
}

func (trArr ByArea) Less(i, j int) bool {
	return trArr[i].Area() > trArr[j].Area() //instead of < to change sorting order
}

func (trArr ByArea) Swap(i, j int) {
	trArr[i], trArr[j] = trArr[j], trArr[i]
}
