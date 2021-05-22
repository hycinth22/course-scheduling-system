package scheduling

type ID1 struct {
	timespanInDay
	instructID int
	roomID     int
}

func (t *ID1) Equal(equal IEqual) bool {
	r, ok := equal.(*ID1)
	if !ok {
		return false
	}
	return t.timespan == r.timespan && t.dayOfWeek == r.dayOfWeek
}

type ID3 struct {
	timespanInDay
	instructID int
}

func (t *ID3) Equal(equal IEqual) bool {
	r, ok := equal.(*ID3)
	if !ok {
		return false
	}
	return t.timespan == r.timespan && t.dayOfWeek == r.dayOfWeek
}

func (X *GeneticSchedule) Invalidity() (total int) {
	pool := X.parent.p
	// 一个老师在同一时间只能安排一门课程，但可以同时为多个班级在同一教室上同一门课
	var cnt0 = 0
	for _, items4T := range X.queryByTeacher {
		//log.Printf(teacher)
		all := createSliceSet(pool)
		for _, item := range items4T {
			key := &ID1{timespanInDay: timespanInDay{item.TimespanID, item.DayOfWeek}, instructID: item.InstructID, roomID: item.ClassroomID}
			//log.Printf("teacher%v %+v is allocated to Instruct%v", item.Instruct.Teacher.Id, key, item.Instruct.InstructId)
			val, exist := all.Get(key)
			if exist {
				v := val.(*ID1)
				// log.Println("item", item, "v", v)
				if v.instructID != item.InstructID || v.roomID != item.ClassroomID {
					// log.Println("detect against h1", item, v)
					cnt0++
				}
			} else {
				all.Insert(key)
			}

		}
		all.Free()
	}
	X.scores.h[0] = cnt0
	// 一个班级在同一时间只能安排一门课程
	var cnt1 = 0
	for _, items4C := range X.queryByClazz {
		all := createSliceSet(pool)
		// var all = createSliceSet(len(items4C))
		for _, item := range items4C {
			key := &timespanInDay{item.TimespanID, item.DayOfWeek}
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
	// 一个教室在同一时间只能安排一门课程，但可以有多个班级在同个教室一起上课
	var cnt2 = 0
	for _, items4CR := range X.queryByClazzroom {
		all := createSliceSet(pool)
		// var all = createSliceSet(len(items4CR))
		for _, item := range items4CR {
			key := &ID3{timespanInDay{item.TimespanID, item.DayOfWeek}, item.InstructID}
			// log.Printf("clazzroom%v %+v is allocated to Instruct%v", item.Clazzroom.Id, key, item.Instruct.InstructId)
			val, exist := all.Get(key)
			if exist {
				if val.(*ID3).instructID != item.InstructID {
					if cnt0 == 0 && cnt1 == 0 && cnt2 == 0 {
						//log.Println("detect against h3", item, val.(*ID3).instructID)
					}
					cnt2++
				}
			} else {
				all.Insert(key)
			}
		}
		all.Free()
	}
	X.scores.h[2] = cnt2
	return cnt0 + cnt1 + cnt2
}
