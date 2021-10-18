package main

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