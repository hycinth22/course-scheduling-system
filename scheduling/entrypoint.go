package scheduling

import "courseScheduling/models"

type Params struct {
	AllInstructedClazz []*models.InstructedClazz
	AllClazzroom       []*models.Clazzroom
	AllTimespan        []*models.Timespan
}

func GenerateSchedule(params *Params) []*models.ScheduleItem {
	return NewGenerator(params).GenerateSchedule().items
}
