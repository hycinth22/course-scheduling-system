package main

import (
	"fmt"

	"courseScheduling/models"
	"courseScheduling/scheduling"
)

func main() {
	// allCourses, err := models.AllCourses()
	// allInstructs := dummy.ParseInstruct()
	// allClazzes := dummy.ParseClazz()
	s := &models.Semester{StartDate: "2021/3/1"}
	allInstructedClazz, err := models.AllInstructedClazzesForScheduling(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	allClazzroom, err := models.AllClazzroom()
	if err != nil {
		fmt.Println(err)
		return
	}
	allTimespan, err := models.AllTimespan()
	if err != nil {
		fmt.Println(err)
		return
	}
	//log.SetOutput(f)
	for i := 0; i < 100; i++ {
		result := scheduling.GenerateSchedule(&scheduling.Params{
			AllInstructedClazz: allInstructedClazz,
			AllClazzroom:       allClazzroom,
			AllTimespan:        allTimespan,
		})
		_, err = models.AddNewSchedule(s, result, len(allTimespan))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
