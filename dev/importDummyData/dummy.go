package main

import (
	"courseScheduling/dev/dummy"
	"courseScheduling/models"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	checkError(models.ImportCourse(dummy.ParseCourse()))
	checkError(models.ImportTeacher(dummy.ParseTeacher()))
	checkError(models.ImportInstruct(dummy.ParseInstruct()))
	checkError(models.ImportInstructedClazz(dummy.ParseInstructedClazz()))
	checkError(models.ImportClazz(dummy.ParseClazz()))
	checkError(models.ImportClazzroom(dummy.ParseClazzroom()))
	checkError(models.ImportTimespan(dummy.ParseTimespan()))
	checkError(models.ImportSemester(dummy.ParseSemester()))
	checkError(models.ImportColleges(dummy.ParseCollege()))
}
