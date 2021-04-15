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
	allInstructedClazz, err := models.AllInstructedClazzesForScheduling()
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
	result := scheduling.GenerateSchedule(&scheduling.Params{
		AllInstructedClazz: allInstructedClazz,
		AllClazzroom:       allClazzroom,
		AllTimespan:        allTimespan,
	})
	for _, item := range result {
		fmt.Printf("%v\n", item)
	}
}
