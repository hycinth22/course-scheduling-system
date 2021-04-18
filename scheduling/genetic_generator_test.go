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
