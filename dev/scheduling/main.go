package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"courseScheduling/models"
	"courseScheduling/scheduling"
)

var file *os.File

func init() {
	var err error
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err = ioutil.TempFile(wd, "example.*.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(file.Name())
}

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
		fmt.Fprintln(file, i, ":")
		begin := time.Now()
		result, score := scheduling.GenerateSchedule(&scheduling.Params{
			AllInstructedClazz: allInstructedClazz,
			AllClazzroom:       allClazzroom,
			AllTimespan:        allTimespan,
			UseEvaluator: []string{
				"AvoidUseNight", "DisperseSameCourse", "KeepAllLessonsDisperseEveryTimespan", "KeepAllLessonsDisperseEveryDay",
			},
		})
		_, err = models.AddNewSchedule(s, result, len(allTimespan), score)
		if err != nil {
			fmt.Println(err)
			return
		}
		end := time.Now()
		costTime := end.Sub(begin)
		fmt.Fprintln(file, i, score, costTime)
	}
}
