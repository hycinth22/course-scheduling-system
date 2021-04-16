package scheduling

import "log"

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

type pair struct {
	first, second interface{}
}
