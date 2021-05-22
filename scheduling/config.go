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
		MaxGenerations:        10e6,
		Timeout:               15 * time.Second,
		StopWhenFitnessKeep:   50,
		FitnessJudgePrecision: 0.1,

		MutationRate:  0.6,
		CrossoverRate: 0.1,

		NumOfPopulations: 2,
		SizeOfPopulation: 100,
		NumOfOffsprings:  25,

		MigrateAfterNGenerations: 200,
		NumberOfMigrants:         10,
	}
}
