package scheduling

import (
	"math/rand"
	"sort"
	"strconv"

	"courseScheduling/models"
	"github.com/MaxHalford/eaopt"
)

// GeneticSchedule 以基因方法代表一个排课方案
// 修改字段时，记得同时修改MakeGeneticSchedule函数式和Clone方法
type GeneticSchedule struct {
	parent           *Generator
	items            GeneticScheduleItemCollection
	allTimespan      []*models.Timespan
	availableWeekday []int

	queryByTeacher         map[string][]*GeneticScheduleItem
	queryByClazz           map[string][]*GeneticScheduleItem
	queryByClazzroom       map[int][]*GeneticScheduleItem
	queryByInstructedClazz map[string][]*GeneticScheduleItem

	scores struct {
		invalidity float64
		h          ScoreHard
		s          ScoreSoft
	}
}

func (s *GeneticSchedule) String() string {
	return GeneticSchedule2String(s)
}

// MakeGeneticSchedule creates a random GeneticSchedule
func MakeGeneticSchedule(g *Generator, rng *rand.Rand) *GeneticSchedule {
	cntNeedToAlloc := g.cntLessons
	var scheduleItems GeneticScheduleItemCollection
	// generate placements
	placeSeq := randomInts(uint(cntNeedToAlloc), 0, len(g.allPlacement)-1, rng)
	// generate ScheduleItems
	var k = 0
	for _, instructedClazz := range g.Params.AllInstructedClazz {
		for cnt := 0; cnt < instructedClazz.Instruct.Course.LessonsPerWeek; cnt++ {
			seq := placeSeq[k]
			item := GeneticScheduleItem{
				InstructID: instructedClazz.Instruct.InstructId,
				ClassID:    instructedClazz.Clazz.ClazzId,
				TeacherID:  instructedClazz.Instruct.Teacher.Id,
				// filled by the algorithm
				ClassroomID:             g.allPlacement[seq].loc,
				TimespanID:              g.allPlacement[seq].timespan,
				DayOfWeek:               g.allPlacement[seq].dayOfWeek,
				TendToObtainDayTimespan: instructedClazz.Instruct.Course.Kind == "必修" || instructedClazz.Instruct.Course.ExamMode == "考试",
			}
			scheduleItems = append(scheduleItems, &item)
			k++
		}
	}
	// generate index
	queryByTeacher := make(map[string][]*GeneticScheduleItem, len(g.Params.AllInstructedClazz))
	queryByClazz := make(map[string][]*GeneticScheduleItem, len(g.Params.AllInstructedClazz))
	queryByClazzroom := make(map[int][]*GeneticScheduleItem, len(g.Params.AllClazzroom))
	queryByInstructedClazz := make(map[string][]*GeneticScheduleItem, len(g.Params.AllInstructedClazz))
	for _, item := range scheduleItems {
		queryByTeacher[item.TeacherID] = append(queryByTeacher[item.TeacherID], item)
		queryByClazz[item.ClassID] = append(queryByClazz[item.ClassID], item)
		queryByClazzroom[item.ClassroomID] = append(queryByClazzroom[item.ClassroomID], item)
		key := strconv.Itoa(item.InstructID) + "_" + item.ClassID
		queryByInstructedClazz[key] = append(queryByInstructedClazz[key], item)
	}
	for _, items := range queryByInstructedClazz {
		sort.Slice(items, func(i, j int) bool {
			if items[i].DayOfWeek != items[j].DayOfWeek {
				return items[i].DayOfWeek < items[j].DayOfWeek
			}
			return items[i].TimespanID < items[j].TimespanID
		})
	}
	s := &GeneticSchedule{
		parent:                 g,
		items:                  scheduleItems,
		allTimespan:            g.Params.AllTimespan,
		availableWeekday:       availableWeekday,
		queryByTeacher:         queryByTeacher,
		queryByClazz:           queryByClazz,
		queryByClazzroom:       queryByClazzroom,
		queryByInstructedClazz: queryByInstructedClazz,
	}
	return s
}

// Clone a GeneticSchedule. Deep-copy ScheduleItems
func (X *GeneticSchedule) Clone() eaopt.Genome {
	var (
		items                  = X.items.Copy().(GeneticScheduleItemCollection)
		queryByTeacher         = make(map[string][]*GeneticScheduleItem, len(X.queryByTeacher))
		queryByClazz           = make(map[string][]*GeneticScheduleItem, len(X.queryByClazz))
		queryByClazzroom       = make(map[int][]*GeneticScheduleItem, len(X.queryByClazzroom))
		queryByInstructedClazz = make(map[string][]*GeneticScheduleItem, len(X.queryByInstructedClazz))
	)
	for i := range items {
		queryByTeacher[items[i].TeacherID] = append(queryByTeacher[items[i].TeacherID], items[i])
		queryByClazz[items[i].ClassID] = append(queryByClazz[items[i].ClassID], items[i])
		queryByClazzroom[items[i].ClassroomID] = append(queryByClazzroom[items[i].ClassroomID], items[i])
		key := strconv.Itoa(items[i].InstructID) + "_" + items[i].ClassID
		queryByInstructedClazz[key] = append(queryByInstructedClazz[key], items[i])
	}
	for _, items := range queryByInstructedClazz {
		sort.Slice(items, func(i, j int) bool {
			if items[i].DayOfWeek != items[j].DayOfWeek {
				return items[i].DayOfWeek < items[j].DayOfWeek
			}
			return items[i].TimespanID < items[j].TimespanID
		})
	}
	return &GeneticSchedule{
		parent:                 X.parent,
		items:                  items,
		allTimespan:            X.allTimespan,
		availableWeekday:       X.availableWeekday,
		queryByTeacher:         queryByTeacher,
		queryByClazz:           queryByClazz,
		queryByClazzroom:       queryByClazzroom,
		queryByInstructedClazz: queryByInstructedClazz,
		scores:                 X.scores,
	}
}

// Mutate two ScheduleItem's placement
func (X *GeneticSchedule) Mutate(rng *rand.Rand) {
	const times = 1
	if X.items.Len() <= 2 {
		return
	}
	for i := 0; i < times; i++ {
		item := X.items[rng.Intn(len(X.items))]
		item.TimespanID = X.parent.Params.AllTimespan[rng.Intn(len(X.parent.Params.AllTimespan))].Id
		item.DayOfWeek = availableWeekday[rng.Intn(len(availableWeekday))]
		item.ClassroomID = X.parent.Params.AllClazzroom[rng.Intn(len(X.parent.Params.AllClazzroom))].Id
	}
}

// Crossover GeneticSchedule with another by applying ScheduleItem crossover.
func (X *GeneticSchedule) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	if X.items.Len() <= 1 {
		return
	}
	const times = 2
	for i := 0; i < times; i++ {
		// Choose two items randomly
		var ids = randomInts(2, 0, X.items.Len(), rng)
		i, j := ids[0], ids[1]
		if rng.Float64() < 0.5 {
			X.items[i].TimespanID, Y.(*GeneticSchedule).items[j].TimespanID = Y.(*GeneticSchedule).items[j].TimespanID, X.items[i].TimespanID
			X.items[i].DayOfWeek, Y.(*GeneticSchedule).items[j].DayOfWeek = Y.(*GeneticSchedule).items[j].DayOfWeek, X.items[i].DayOfWeek
		} else {
			X.items[i].ClassroomID, Y.(*GeneticSchedule).items[j].ClassroomID = Y.(*GeneticSchedule).items[j].ClassroomID, X.items[i].ClassroomID
		}
	}
}
