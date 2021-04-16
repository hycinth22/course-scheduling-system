package scheduling

import (
	"math/rand"
	"testing"
	"time"

	"courseScheduling/models"
)

var testParams = &Params{}

func init() {
	var err error
	testParams.AllInstructedClazz, err = models.AllInstructedClazzesForScheduling()
	if err != nil {
		panic(err)
		return
	}
	testParams.AllClazzroom, err = models.AllClazzroom()
	if err != nil {
		panic(err)
		return
	}
	testParams.AllTimespan, err = models.AllTimespan()
	if err != nil {
		panic(err)
		return
	}
}

func BenchmarkGeneticSchedule_Clone(b *testing.B) {
	g := NewGenerator(testParams)
	s := MakeGeneticSchedule(g, rand.New(rand.NewSource(time.Now().UnixNano())))
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		schedule := s.Clone().(*GeneticSchedule)
		schedule.scores.invalidity = 0
	}
	b.StopTimer()
}

func BenchmarkGeneticSchedule_String(b *testing.B) {
	g := NewGenerator(testParams)
	s := MakeGeneticSchedule(g, rand.New(rand.NewSource(time.Now().UnixNano())))
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
	b.StopTimer()
}

func BenchmarkMakeGeneticSchedule(b *testing.B) {
	g := NewGenerator(testParams)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = MakeGeneticSchedule(g, rng)
	}
	b.StopTimer()
}
