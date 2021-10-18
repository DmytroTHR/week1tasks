package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Envelope struct {
	height, width float64
}

type Dimensions interface {
	maxDimension() float64
	minDimension() float64
	FitsInto(envelopeToFit Envelope) bool
}

func (env Envelope) maxDimension() float64 {
	if env.height > env.width {
		return env.height
	}
	return env.width
}

func (env Envelope) minDimension() float64 {
	if env.height < env.width {
		return env.height
	}
	return env.width
}

func (env Envelope) FitsInto(envelopeToFit Envelope) bool {
	if env.maxDimension() < envelopeToFit.maxDimension() && env.minDimension() < envelopeToFit.minDimension() {
		return true
	}
	return false
}

func EnvelopeInitManually() (Envelope, error) {
	var strErr []string

	height, err1 := askForParamFloat("Envelope height:")
	if err1 != nil {
		strErr = append(strErr, "INCORRECT input - height parsing error")
	}
	width, err2 := askForParamFloat("Envelope width:")
	if err2 != nil {
		strErr = append(strErr, "INCORRECT input - width parsing error")
	}

	if err1 == nil && err2 == nil && (height <= 0 || width <= 0) {
		strErr = append(strErr, "INCORRECT input - height & width must be greater than 0")
	}

	if len(strErr) > 0 {
		return Envelope{}, errors.New(strings.Join(strErr, "\n"))
	}

	return Envelope{height, width}, nil
}

func askForParamFloat(strToAsk string) (float64, error) {
	fmt.Println(strToAsk)
	var inp string
	fmt.Scanln(&inp)
	return strconv.ParseFloat(inp, 64)
}