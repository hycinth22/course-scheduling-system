package scheduling

import (
	"testing"

	"courseScheduling/models"
)

func TestGenerator_GenerateSchedule(t *testing.T) {
	allInstructedClazz, err := models.AllInstructedClazzesForScheduling()
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
