package scheduling

import "time"

type ConfigType struct {
	MaxGenerations        uint
	Timeout               time.Duration
	StopWhenFitnessKeep   int
	FitnessJudgePrecision float64
}

var Config ConfigType

func init() {
	Config = ConfigType{
		MaxGenerations:        10e9,
		Timeout:               60 * time.Second,
		StopWhenFitnessKeep:   150,
		FitnessJudgePrecision: 0.01,
	}
}
