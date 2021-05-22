package scheduling

import (
	"fmt"

	"github.com/MaxHalford/eaopt"
)

type GeneticScheduleItem struct {
	// primary keys
	InstructID int    // 授课ID，唯一确定教师开课
	ClassID    string // 课程ID，唯一确定班级
	// algorithm fill
	ClassroomID int // 教室ID，唯一确定教室
	TimespanID  int // 时间段ID
	DayOfWeek   int // 星期几
	// extra info
	TeacherID               string
	TendToObtainDayTimespan bool
}
type GeneticScheduleItemCollection []*GeneticScheduleItem

func (i *GeneticScheduleItem) Clone() *GeneticScheduleItem {
	baby := *i
	return &baby
}
func (i *GeneticScheduleItem) String() string {
	return fmt.Sprintf("课%v 班级%v 教室%v 老师%v", i.InstructID, i.ClassID, i.ClassroomID, i.TeacherID)
}

// implements the interface eaopt.Slice

func (s GeneticScheduleItemCollection) At(i int) interface{} {
	return s[i]
}

func (s GeneticScheduleItemCollection) Set(i int, v interface{}) {
	value := v.(*GeneticScheduleItem)
	s[i].TimespanID = value.TimespanID
	s[i].ClassroomID = value.ClassroomID
}

func (s GeneticScheduleItemCollection) Len() int {
	return len(s)
}

func (s GeneticScheduleItemCollection) Swap(i, j int) {
	s[i].TimespanID, s[j].TimespanID = s[j].TimespanID, s[i].TimespanID
	s[i].ClassroomID, s[j].ClassroomID = s[j].ClassroomID, s[i].ClassroomID
	s[i].DayOfWeek, s[j].DayOfWeek = s[j].DayOfWeek, s[i].DayOfWeek
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
	var t = make([]*GeneticScheduleItem, len(s))
	for i := range t {
		// and deep copy all elements in the slice
		t[i] = s[i].Clone()
	}
	return GeneticScheduleItemCollection(t)
}
