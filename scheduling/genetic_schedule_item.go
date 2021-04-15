package scheduling

import (
	"courseScheduling/models"
	"github.com/MaxHalford/eaopt"
)

type GeneticScheduleItemCollection []*models.ScheduleItem

// implements the interface eaopt.Slice

func (s GeneticScheduleItemCollection) At(i int) interface{} {
	return s[i]
}

func (s GeneticScheduleItemCollection) Set(i int, v interface{}) {
	value := v.(*models.ScheduleItem)
	s[i].TimespanId = value.TimespanId
	s[i].ClazzroomId = value.ClazzroomId
	// s[i] = v.(*models.ScheduleItem)
}

func (s GeneticScheduleItemCollection) Len() int {
	return len(s)
}

func (s GeneticScheduleItemCollection) Swap(i, j int) {
	s[i].TimespanId, s[j].TimespanId = s[j].TimespanId, s[i].TimespanId
	s[i].ClazzroomId, s[j].ClazzroomId = s[j].ClazzroomId, s[i].ClazzroomId
	// log.Println("swap ", i , j, s[i].TimespanId, s[j].TimespanId , s[j].TimespanId, s[i].TimespanId)
	//log.Println("swap Instruct", s[i].Instruct.InstructId, s[j].Instruct.InstructId)
	// s[i], s[j] = s[j], s[i]
}

func (s GeneticScheduleItemCollection) Slice(a, b int) eaopt.Slice {
	return s[a:b]
}

func (s GeneticScheduleItemCollection) Split(k int) (eaopt.Slice, eaopt.Slice) {
	return s[:k], s[k:]
}

func (s GeneticScheduleItemCollection) Append(slice eaopt.Slice) eaopt.Slice {
	return append(s, slice.(GeneticScheduleItemCollection)...)
}

func (s GeneticScheduleItemCollection) Replace(slice eaopt.Slice) {
	copy(s, slice.(GeneticScheduleItemCollection))
}

func (s GeneticScheduleItemCollection) Copy() eaopt.Slice {
	var t = make(GeneticScheduleItemCollection, len(s))
	copy(t, s) // generate new slice
	for i := range t {
		// and deep copy all elements in the slice
		t[i] = &models.ScheduleItem{
			ScheduleItemId: t[i].ScheduleItemId,
			ScheduleId:     t[i].ScheduleId,
			Instruct:       t[i].Instruct,
			Clazz:          t[i].Clazz,
			ClazzroomId:    t[i].ClazzroomId,
			TimespanId:     t[i].TimespanId,
			DayOfWeek:      t[i].DayOfWeek,
			CreatedAt:      t[i].CreatedAt,
			UpdatedAt:      t[i].UpdatedAt,
		}
	}
	return t
}
