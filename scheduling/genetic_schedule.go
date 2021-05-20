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
	for _, instructedClazz := range g.params.AllInstructedClazz {
		for cnt := 0; cnt < instructedClazz.Instruct.Course.LessonsPerWeek; cnt++ {
			seq := placeSeq[k]
			item := &models.ScheduleItem{
				ScheduleItemId: k, // keep empty, filled by models package
				ScheduleId:     0, // keep empty, filled by models package
				Instruct:       instructedClazz.Instruct,
				Clazz:          instructedClazz.Clazz,
				// filled by the algorithm
				Clazzroom:  &models.Clazzroom{Id: g.allPlacement[seq].loc},
				TimespanId: g.allPlacement[seq].timespan,
				DayOfWeek:  g.allPlacement[seq].dayOfWeek,
			}
			scheduleItems = append(scheduleItems, item)
			k++
		}
	}
	// generate index
	queryByTeacher := make(map[string][]*models.ScheduleItem, len(g.params.AllInstructedClazz))
	queryByClazz := make(map[string][]*models.ScheduleItem, len(g.params.AllInstructedClazz))
	queryByClazzroom := make(map[int][]*models.ScheduleItem, len(g.params.AllClazzroom))
	queryByInstructedClazz := make(map[string][]*models.ScheduleItem, len(g.params.AllInstructedClazz))
	for _, item := range scheduleItems {
		queryByTeacher[item.Instruct.Teacher.Id] = append(queryByTeacher[item.Instruct.Teacher.Id], item)
		queryByClazz[item.Clazz.ClazzId] = append(queryByClazz[item.Clazz.ClazzId], item)
		queryByClazzroom[item.Clazzroom.Id] = append(queryByClazzroom[item.Clazzroom.Id], item)
		key := strconv.Itoa(item.Instruct.InstructId) + "_" + item.Clazz.ClazzId
		queryByInstructedClazz[key] = append(queryByInstructedClazz[key], item)
	}
	for _, items := range queryByInstructedClazz {
		sort.Slice(items, func(i, j int) bool {
			if items[i].DayOfWeek != items[j].DayOfWeek {
				return items[i].DayOfWeek < items[j].DayOfWeek
			}
			return items[i].TimespanId < items[j].TimespanId
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
		queryByClazzroom[items[i].Clazzroom.Id] = append(queryByClazzroom[items[i].Clazzroom.Id], items[i])
		key := strconv.Itoa(items[i].Instruct.InstructId) + "_" + items[i].Clazz.ClazzId
		queryByInstructedClazz[key] = append(queryByInstructedClazz[key], items[i])
	}
	for _, items := range queryByInstructedClazz {
		sort.Slice(items, func(i, j int) bool {
			if items[i].DayOfWeek != items[j].DayOfWeek {
				return items[i].DayOfWeek < items[j].DayOfWeek
			}
			return items[i].TimespanId < items[j].TimespanId
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

// Mutate swap two ScheduleItem's placement at 2 times.
func (X *GeneticSchedule) Mutate(rng *rand.Rand) {
	const times = 1
	// eaopt.MutPermute(X.items, times, rng)
	if X.items.Len() <= 1 {
		return
	}
	for i := 0; i < times; i++ {
		// Choose two items randomly
		var ids = randomInts(2, 0, X.items.Len(), rng)
		X.items.Swap(ids[0], ids[1])
	}
}

// Crossover GeneticSchedule with another by applying 1-ScheduleItem crossover.
func (X *GeneticSchedule) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	// eaopt.CrossERX(X.items, Y.(*GeneticSchedule).items)
	eaopt.CrossGNX(X.items, Y.(*GeneticSchedule).items, 2, rng)
	// eaopt.CrossCX(X.items, Y.(*GeneticSchedule).items)
}
