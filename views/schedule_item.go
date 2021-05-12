package views

import "courseScheduling/models"

type ScheduleItemsTableView struct {
	// [clazz/dept][timespan][week][]
	ByClazz           map[string][]map[int]*models.ScheduleItem   `json:"by_clazz"`
	ByDept            map[string][]map[int][]*models.ScheduleItem `json:"by_dept"`
	ByTeacherPersonal map[string][]map[int]*models.ScheduleItem   `json:"by_teacher_personal"`
}

func NewScheduleItemsTableView(cap int) *ScheduleItemsTableView {
	return &ScheduleItemsTableView{
		ByClazz:           make(map[string][]map[int]*models.ScheduleItem, cap),
		ByDept:            make(map[string][]map[int][]*models.ScheduleItem, cap),
		ByTeacherPersonal: make(map[string][]map[int]*models.ScheduleItem, cap),
	}
}
