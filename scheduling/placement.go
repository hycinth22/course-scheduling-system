package scheduling

import "courseScheduling/models"

type timespanInDay struct {
	timespan  int
	dayOfWeek int
}

func (t *timespanInDay) Equal(equal IEqual) bool {
	r, ok := equal.(*timespanInDay)
	if !ok {
		return false
	}
	return t.timespan == r.timespan && t.dayOfWeek == r.dayOfWeek
}

type placement struct {
	timespan  int
	loc       int
	dayOfWeek int
}

func product2Placement(allClazzroom []*models.Clazzroom, allTimespan []*models.Timespan, availableWeekday []int) []placement {
	result := make([]placement, len(allTimespan)*len(availableWeekday)*len(allClazzroom))
	cnt := 0
	for i := range availableWeekday {
		for j := range allTimespan {
			for k := range allClazzroom {
				result[cnt] = placement{
					dayOfWeek: availableWeekday[i],
					timespan:  allTimespan[j].Id,
					loc:       allClazzroom[k].Id,
				}
				cnt++
			}
		}
	}
	return result
}
