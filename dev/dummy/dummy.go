package dummy

import (
	"log"
	"os"

	"courseSchduling/excel"
	"courseSchduling/models"
)

const (
	dataDir        = "./dev/data/"
	excelCourse    = dataDir + "/course.xls"
	excelTeacher   = dataDir + "/teacher.xls"
	excelInstruct  = dataDir + "/instruct.xls"
	excelClazz     = dataDir + "/clazz.xls"
	excelClazzroom = dataDir + "/clazzroom.xls"
	excelTimespan  = dataDir + "/timespan.xls"
	excelSemester  = dataDir + "/semester.xls"
)

func ParseCourse() []*models.Course {
	f, err := os.Open(excelCourse)
	if err != nil {
		log.Println(err)
	}
	return excel.ParseCourseExcel(f)
}

func ParseTeacher() []*models.Teacher {
	f, err := os.Open(excelTeacher)
	if err != nil {
		log.Println(err)
	}
	return excel.ParseTeacherExcel(f)
}
func ParseInstruct() []*models.Instruct {
	f, err := os.Open(excelInstruct)
	if err != nil {
		log.Println(err)
	}
	return excel.ParseInstructExcel(f)
}
func ParseClazz() []*models.Clazz {
	f, err := os.Open(excelClazz)
	if err != nil {
		log.Println(err)
	}
	return excel.ParseClazzExcel(f)
}
func ParseClazzroom() []*models.Clazzroom {
	f, err := os.Open(excelClazzroom)
	if err != nil {
		log.Println(err)
	}
	return excel.ParseClazzroomExcel(f)
}
func ParseTimespan() []*models.Timespan {
	f, err := os.Open(excelTimespan)
	if err != nil {
		log.Println(err)
	}
	return excel.ParseTimespanExcel(f)
}
func ParseSemester() []*models.Semester {
	f, err := os.Open(excelSemester)
	if err != nil {
		log.Println(err)
	}
	return excel.ParseSemesterExcel(f)
}
