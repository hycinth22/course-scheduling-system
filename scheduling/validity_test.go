package scheduling

import (
	"math/rand"
	"testing"
	"time"

	"courseScheduling/models"
)

func BenchmarkGeneticschedule_Invalidity(b *testing.B) {
	allInstructedClazz, err := models.AllInstructedClazzesForScheduling(testSemester)
	if err != nil {
		b.FailNow()
		return
	}
	allClazzroom, err := models.AllClazzroom()
	if err != nil {
		b.FailNow()
		return
	}
	allTimespan, err := models.AllTimespan()
	if err != nil {
		b.FailNow()
		return
	}
	g := NewGenerator(&Params{
		AllInstructedClazz: allInstructedClazz,
		AllClazzroom:       allClazzroom,
		AllTimespan:        allTimespan,
	}, DefaultConfig)
	s := MakeGeneticSchedule(g, rand.New(rand.NewSource(time.Now().UnixNano())))
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Invalidity()
	}
	b.StopTimer()
}
