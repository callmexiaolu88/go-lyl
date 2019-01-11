package utils

import (
	"math/rand"
	"time"
)

var randSource rand.Source
var globalRand *rand.Rand

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
	globalRand = rand.New(randSource)
}

func RandFloat32Slice(max int, length int) []float32 {
	result := make([]float32, length)
	for k := range result {
		result[k] = RandFloat32(max)
	}
	return result
}

func RandFloat32(max int) float32 {
	return (float32(max) - 1) + globalRand.Float32()
}
