package scheduling

import "courseScheduling/models"

func GenerateSchedule(params *Params) ([]*models.ScheduleItem, float64) {
	sch, score := NewGenerator(params, DefaultConfig).GenerateSchedule()
	return sch.items, score
}
