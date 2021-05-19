package views

import "courseScheduling/models"

type ScheduleItemsTableView struct {
	// [clazz/dept/teacher][timespan][week][]
	ByClazz           map[string][]map[int]*models.ScheduleItem   `json:"by_clazz"`
	ByDept            map[string][]map[int][]*models.ScheduleItem `json:"by_dept"`
	ByTeacherPersonal map[string][]map[int]*models.ScheduleItem   `json:"by_teacher_personal"`
	Entire            []map[int][]*models.ScheduleItem            `json:"entire"`
}

func NewScheduleItemsTableView(cap int, useTimespan int) *ScheduleItemsTableView {
	s := &ScheduleItemsTableView{
		ByClazz:           make(map[string][]map[int]*models.ScheduleItem, cap),
		ByDept:            make(map[string][]map[int][]*models.ScheduleItem, cap),
		ByTeacherPersonal: make(map[string][]map[int]*models.ScheduleItem, cap),
		Entire:            make([]map[int][]*models.ScheduleItem, useTimespan),
	}
	// create
	const cntWeekday = 7
	for timespan := 0; timespan < useTimespan; timespan++ {
		s.Entire[timespan] = make(map[int][]*models.ScheduleItem, cntWeekday)
	}
	return s
}
