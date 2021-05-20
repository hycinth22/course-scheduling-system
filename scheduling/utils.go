package scheduling

import (
	"log"
	"math/rand"
)

func normalizeFloat64(x, min, max float64) float64 {
	if x < min {
		log.Println("normalizeFloat64 lower than min:", x, min)
		return 0.0
	}
	if x > max {
		log.Println("normalizeFloat64 upper than max:", x, max)
		return 1.0
	}
	return (x - min) / (max - min)
}

func randomInts(k uint, min, max int, rng *rand.Rand) []int {
	var ints = make([]int, k)
	for i := 0; i < int(k); i++ {
		ints[i] = i + min
	}
	for i := int(k); i < max-min; i++ {
		var j = rng.Intn(i + 1)
		if j < int(k) {
			ints[j] = i + min
		}
	}
	return ints
}
