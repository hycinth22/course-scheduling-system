package scheduling

import (
	"courseScheduling/models"
)

func GenerateSchedule(params *Params) (sch []*models.ScheduleItem, score float64) {
	var gsch *GeneticSchedule
	for gsch == nil {
		gsch, score = NewGenerator(params, DefaultConfig).GenerateSchedule()
	}
	sch = make([]*models.ScheduleItem, gsch.items.Len())
	for i := range sch {
		sch[i] = &models.ScheduleItem{
			ScheduleItemId: 0, // keep empty, filled by models package
			ScheduleId:     0, // keep empty, filled by models package
			Instruct:       &models.Instruct{InstructId: gsch.items[i].InstructID},
			Clazz:          &models.Clazz{ClazzId: gsch.items[i].ClassID},
			Clazzroom:      &models.Clazzroom{Id: gsch.items[i].ClassroomID},
			TimespanId:     gsch.items[i].TimespanID,
			DayOfWeek:      gsch.items[i].DayOfWeek,
		}
	}
	return
}
