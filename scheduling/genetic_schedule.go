package scheduling

import (
	"math/rand"
	"sort"
	"strconv"

	"courseScheduling/models"
	"github.com/MaxHalford/eaopt"
)

// GeneticSchedule 以基因方法代表一个排课方案
// 修改字段时，记得同时修改MakeGeneticSchedule函数式和Clone方法，，
type GeneticSchedule struct {
	parent           *Generator
	items            GeneticScheduleItemCollection
	allTimespan      []*models.Timespan
	availableWeekday []int

	queryByTeacher         map[string][]*models.ScheduleItem
	queryByClazz           map[string][]*models.ScheduleItem
	queryByClazzroom       map[int][]*models.ScheduleItem
	queryByInstructedClazz map[string][]*models.ScheduleItem

	scores struct {
		invalidity float64
		h          [3]float64
		s          [3]float64
	}
}

func (s *GeneticSchedule) String() string {
	return GeneticSchedule2String(s)
}

// MakeGeneticSchedule creates a random GeneticSchedule
func MakeGeneticSchedule(g *Generator, rng *rand.Rand) *GeneticSchedule {
	var scheduleItems GeneticScheduleItemCollection
	// generate placements
	placeSeq := make([]int, len(g.allPlacement))
	for i := range g.allPlacement {
		placeSeq[i] = i
	}
	rng.Shuffle(len(placeSeq), func(i, j int) {
		placeSeq[i], placeSeq[j] = placeSeq[j], placeSeq[i]
	})
	// generate ScheduleItems
	var k = 0
	for _, instructedClazz := range g.params.AllInstructedClazz {
		for cnt := 0; cnt < instructedClazz.Instruct.Course.LessonsPerWeek; cnt++ {
			seq := placeSeq[k]
			item := &models.ScheduleItem{
				ScheduleItemId: 0,   // keep empty, filled by models package
				ScheduleId:     nil, // keep empty, filled by models package
				Instruct:       instructedClazz.Instruct,
				Clazz:          instructedClazz.Clazz,
				// filled by the algorithm
				ClazzroomId: g.allPlacement[seq].loc,
				TimespanId:  g.allPlacement[seq].timespan,
				DayOfWeek:   g.allPlacement[seq].dayOfWeek,
			}
			scheduleItems = append(scheduleItems, item)
			k++
		}
	}
	// generate index
	queryByTeacher := make(map[string][]*models.ScheduleItem, len(g.params.AllInstructedClazz)/2)
	queryByClazz := make(map[string][]*models.ScheduleItem, len(g.params.AllInstructedClazz)/3)
	queryByClazzroom := make(map[int][]*models.ScheduleItem, len(g.params.AllClazzroom))
	queryByInstructedClazz := make(map[string][]*models.ScheduleItem, len(g.params.AllInstructedClazz))
	for _, item := range scheduleItems {
		queryByTeacher[item.Instruct.Teacher.Id] = append(queryByTeacher[item.Instruct.Teacher.Id], item)
		queryByClazz[item.Clazz.ClazzId] = append(queryByClazz[item.Clazz.ClazzId], item)
		queryByClazzroom[item.ClazzroomId] = append(queryByClazzroom[item.ClazzroomId], item)
		key := strconv.Itoa(item.Instruct.InstructId) + "_" + item.Clazz.ClazzId
		queryByInstructedClazz[key] = append(queryByInstructedClazz[key], item)
	}
	for ic := range queryByInstructedClazz {
		sort.Slice(queryByInstructedClazz[ic], func(i, j int) bool {
			if queryByInstructedClazz[ic][i].DayOfWeek != queryByInstructedClazz[ic][j].DayOfWeek {
				return queryByInstructedClazz[ic][i].DayOfWeek < queryByInstructedClazz[ic][j].DayOfWeek
			}
			return queryByInstructedClazz[ic][i].TimespanId < queryByInstructedClazz[ic][j].TimespanId
		})
	}
	s := &GeneticSchedule{
		parent:                 g,
		items:                  scheduleItems,
		allTimespan:            g.params.AllTimespan,
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
		queryByTeacher         = make(map[string][]*models.ScheduleItem, len(X.queryByTeacher))
		queryByClazz           = make(map[string][]*models.ScheduleItem, len(X.queryByClazz))
		queryByClazzroom       = make(map[int][]*models.ScheduleItem, len(X.queryByClazzroom))
		queryByInstructedClazz = make(map[string][]*models.ScheduleItem, len(X.queryByInstructedClazz))
	)
	for i := range items {
		queryByTeacher[items[i].Instruct.Teacher.Id] = append(queryByTeacher[items[i].Instruct.Teacher.Id], items[i])
		queryByClazz[items[i].Clazz.ClazzId] = append(queryByClazz[items[i].Clazz.ClazzId], items[i])
		queryByClazzroom[items[i].ClazzroomId] = append(queryByClazzroom[items[i].ClazzroomId], items[i])
		key := strconv.Itoa(items[i].Instruct.InstructId) + "_" + items[i].Clazz.ClazzId
		queryByInstructedClazz[key] = append(queryByInstructedClazz[key], items[i])
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

// Mutate swap two ScheduleItem's placement at 3 times.
func (X *GeneticSchedule) Mutate(rng *rand.Rand) {
	eaopt.MutPermute(X.items, 2, rng)
}

// Crossover GeneticSchedule with another by applying 2-ScheduleItem crossover.
func (X *GeneticSchedule) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossGNX(X.items, Y.(*GeneticSchedule).items, 2, rng)
}