package scheduling

import (
	"math/rand"
)

func (X *GeneticSchedule) Invalidity() (total float64) {
	// 一个老师在同一时间只能安排一门课程
	var cnt0 = 0.0
	for _, items4T := range X.queryByTeacher {
		all := createHashSet(len(items4T))
		for _, item := range items4T {
			key := timespanInDay{item.TimespanId, item.DayOfWeek}
			//log.Printf("teacher%v %+v is allocated to Instruct%v", teacher, key, item.Instruct.InstructId)
			if all.Has(key) {
				//log.Println("detect against h1")
				cnt0++
			}
			all.Insert(key)
		}
		all.Free()
	}
	X.scores.h[0] = cnt0
	// 一个班级在同一时间只能安排一门课程
	var cnt1 = 0.0
	for _, items4C := range X.queryByClazz {
		all := createHashSet(len(items4C))
		for _, item := range items4C {
			key := timespanInDay{item.TimespanId, item.DayOfWeek}
			//log.Printf("clazz%v %+v is allocated to Instruct%v", clazz, key, item.Instruct.InstructId)
			if all.Has(key) {
				//log.Println("detect against h2")
				cnt1++
			}
			all.Insert(key)
		}
		all.Free()
	}
	X.scores.h[1] = cnt1
	// 一个教室在同一时间只能在一个班级安排一门课程
	var cnt2 = 0.0
	for _, items4CR := range X.queryByClazzroom {
		all := createHashSet(len(items4CR))
		for _, item := range items4CR {
			key := pair{timespanInDay{item.TimespanId, item.DayOfWeek}, item.Clazz.ClazzId}
			//log.Printf("clazzroom%v %+v is allocated to Instruct%v", clazzroom, key, item.Instruct.InstructId)
			if all.Has(key) {
				//log.Println("detect against h3")
				cnt2++
			}
			all.Insert(key)
		}
		all.Free()
	}
	X.scores.h[2] = cnt2
	return cnt0 + cnt1 + cnt2
}

func createShuffleIntSeq(n int, raw []int, rng *rand.Rand) (result []int) {
	copy(result, raw)
	rng.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	result = result[:n]
	return
}
