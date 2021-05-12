package scheduling

import "time"

type ConfigType struct {
	MaxGenerations        uint
	Timeout               time.Duration
	StopWhenFitnessKeep   int
	FitnessJudgePrecision float64

	NumOfPopulations, SizeOfPopulation, NumOfOffsprings uint

	MutationRate, CrossoverRate float64

	MigrateAfterNGenerations, NumberOfMigrants int
}

var DefaultConfig ConfigType

func init() {
	DefaultConfig = ConfigType{
		MaxGenerations:        10e9,
		Timeout:               60 * time.Second,
		StopWhenFitnessKeep:   100,
		FitnessJudgePrecision: 0.01,

		MutationRate:  0.95,
		CrossoverRate: 0.01,

		NumOfPopulations: 1,
		SizeOfPopulation: 100,
		NumOfOffsprings:  25,

		MigrateAfterNGenerations: 20,
		NumberOfMigrants:         30,
	}
}
