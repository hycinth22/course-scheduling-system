package scheduling

import "courseScheduling/models"

func GenerateSchedule(params *Params) []*models.ScheduleItem {
	return NewGenerator(params, DefaultConfig).GenerateSchedule().items
}
