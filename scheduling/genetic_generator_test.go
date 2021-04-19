package scheduling

import (
	"testing"

	"courseScheduling/models"
)

var testSemester = &models.Semester{StartDate: "2021/3/1"}

func TestGenerator_GenerateSchedule(t *testing.T) {
	allInstructedClazz, err := models.AllInstructedClazzesForScheduling(testSemester)
	if err != nil {
		t.Error(err)
		return
	}
	allClazzroom, err := models.AllClazzroom()
	if err != nil {
		t.Error(err)
		return
	}
	allTimespan, err := models.AllTimespan()
	if err != nil {
		t.Error(err)
		return
	}
	g := NewGenerator(&Params{
		AllInstructedClazz: allInstructedClazz,
		AllClazzroom:       allClazzroom,
		AllTimespan:        allTimespan,
	})
	g.GenerateSchedule()
}

func BenchmarkGenerator_GenerateSchedule(b *testing.B) {
	allInstructedClazz, err := models.AllInstructedClazzesForScheduling(testSemester)
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	allClazzroom, err := models.AllClazzroom()
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	allTimespan, err := models.AllTimespan()
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	g := NewGenerator(&Params{
		AllInstructedClazz: allInstructedClazz,
		AllClazzroom:       allClazzroom,
		AllTimespan:        allTimespan,
	})
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = g.GenerateSchedule()
	}
	b.StopTimer()
}
