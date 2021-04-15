package scheduling

import "courseScheduling/models"

type timespanInDay struct {
	timespan  int
	dayOfWeek int
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
