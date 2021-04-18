package views

import "courseScheduling/models"

type ScheduleItemsTableView struct {
	// [clazz/dept][timespan][week][]
	ByClazz map[string][]map[int]*models.ScheduleItem   `json:"by_clazz"`
	ByDept  map[string][]map[int][]*models.ScheduleItem `json:"by_dept"`
}

func NewScheduleItemsTableView(cap int) *ScheduleItemsTableView {
	return &ScheduleItemsTableView{
		ByClazz: make(map[string][]map[int]*models.ScheduleItem, cap),
		ByDept:  make(map[string][]map[int][]*models.ScheduleItem, cap),
	}
}
