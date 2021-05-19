package scheduling

import "courseScheduling/models"

func GenerateSchedule(params *Params) ([]*models.ScheduleItem, float64) {
	var (
		sch   *GeneticSchedule
		score float64
	)
	for sch == nil {
		sch, score = NewGenerator(params, DefaultConfig).GenerateSchedule()
	}
	return sch.items, score
}
