package main

import (
	"fmt"
	"log"
	"os"

	"courseScheduling/dev/dummy"
	"courseScheduling/models"
	"courseScheduling/scheduling"
)

var (
	courses   map[string]*models.Course
	instructs map[int]*models.Instruct
	clazzes   map[string]*models.Clazz
)

// dummy data includes only the primary key, the function fill all fields from the corresponding table
func fillInstructedClazz(allInstructedClazz []*models.InstructedClazz) {
	for _, item := range allInstructedClazz {
		ins, exist := instructs[item.Instruct.InstructId]
		if !exist {
			log.Println("cannot found Instruct", item.Instruct.InstructId)
		}
		item.Instruct = ins
		cou, exist := courses[item.Instruct.Course.Id]
		if !exist {
			log.Println("cannot found Course", item.Instruct.Course.Id)
		}
		item.Instruct.Course = cou
		clazz, exist := clazzes[item.Clazz.ClazzId]
		if !exist {
			log.Println("cannot found Clazz", item.Clazz.ClazzId)
		}
		item.Clazz = clazz
	}
}

func main() {
	allCourses := dummy.ParseCourse()
	courses = make(map[string]*models.Course, len(allCourses))
	for _, item := range allCourses {
		courses[item.Id] = item
	}
	allInstructs := dummy.ParseInstruct()
	instructs = make(map[int]*models.Instruct, len(allInstructs))
	for _, item := range allInstructs {
		instructs[item.InstructId] = item
	}
	allClazzes := dummy.ParseClazz()
	clazzes = make(map[string]*models.Clazz, len(allClazzes))
	for _, item := range allClazzes {
		clazzes[item.ClazzId] = item
	}
	allInstructedClazz := dummy.ParseInstructedClazz()
	fillInstructedClazz(allInstructedClazz)
	allClazzroom := dummy.ParseClazzroom()
	allTimespan := dummy.ParseTimespan()
	result := scheduling.GenerateSchedule(&scheduling.Params{
		AllInstructedClazz: allInstructedClazz,
		AllClazzroom:       allClazzroom,
		AllTimespan:        allTimespan,
	})
	f, err := os.Create("s.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range result {
		_, err := fmt.Fprintf(f, "%v\n", item)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
