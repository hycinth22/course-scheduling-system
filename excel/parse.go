package excel

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"courseScheduling/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GetSheet1Rows(reader io.Reader) ([][]string, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sheet := f.GetSheetName(0)
	return f.GetRows(sheet)
}

func ParseCourseExcel(reader io.Reader) (r []*models.Course) {
	rows, err := GetSheet1Rows(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		var err error
		if i == 0 {
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			continue
		}
		lessons, err := strconv.Atoi(strings.TrimSpace(row[4]))
		if err != nil {
			log.Println(err)
			continue
		}
		lpw, err := strconv.Atoi(strings.TrimSpace(row[8]))
		if err != nil {
			log.Println(err)
			continue
		}
		r = append(r, &models.Course{
			Id:             row[2],
			Name:           row[3],
			Lessons:        lessons,
			LessonsPerWeek: lpw,
			ExamMode:       row[6],
			Founder:        row[7],
		})
	}
	return
}

func ParseTimespanExcel(reader io.Reader) (r []*models.Timespan) {
	rows, err := GetSheet1Rows(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		var err error
		if i == 0 {
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			continue
		}
		id, err := strconv.Atoi(strings.TrimSpace(row[0]))
		if err != nil {
			log.Println(err)
			continue
		}
		beginTime, err := time.Parse("15:04:05", strings.TrimSpace(row[1]))
		if err != nil {
			log.Println(err)
			continue
		}
		endTime, err := time.Parse("15:04:05", strings.TrimSpace(row[2]))
		if err != nil {
			log.Println(err)
			continue
		}
		length, err := time.Parse("15:04", strings.TrimSpace(row[3]))
		if err != nil {
			log.Println(err)
			continue
		}
		r = append(r, &models.Timespan{
			Id:          id,
			BeginHour:   beginTime.Hour(),
			BeginMinute: beginTime.Minute(),
			EndHour:     endTime.Hour(),
			EndMinute:   endTime.Minute(),
			Length:      time.Duration(length.Hour())*time.Hour + time.Duration(length.Minute())*time.Minute,
		})
	}
	return
}

func ParseTeacherExcel(reader io.Reader) (r []*models.Teacher) {
	rows, err := GetSheet1Rows(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			continue
		}
		r = append(r, &models.Teacher{
			Id:    strings.TrimSpace(row[0]),
			Name:  strings.TrimSpace(row[1]),
			Title: strings.TrimSpace(row[2]),
			Tel:   strings.TrimSpace(row[3]),
			Dept:  &models.Department{DeptId: strings.TrimSpace(row[4])},
		})
	}
	return
}

func ParseSemesterExcel(reader io.Reader) (r []*models.Semester) {
	rows, err := GetSheet1Rows(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		var err error
		if i == 0 {
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			continue
		}
		beginDate := strings.TrimSpace(row[0])
		name := strings.TrimSpace(row[1])
		weeks, err := strconv.Atoi(strings.TrimSpace(row[2]))
		if err != nil {
			log.Println(err)
			continue
		}
		r = append(r, &models.Semester{
			StartDate: beginDate,
			Name:      name,
			Weeks:     weeks,
		})
	}
	return
}

func ParseClazzExcel(reader io.Reader) (r []*models.Clazz) {
	rows, err := GetSheet1Rows(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			continue
		}
		id := strings.TrimSpace(row[0])
		name := strings.TrimSpace(row[1])
		majorid := strings.TrimSpace(row[2])
		r = append(r, &models.Clazz{
			ClazzId:   id,
			ClazzName: name,
			Major:     &models.Major{MajorId: majorid},
		})
	}
	return
}

func ParseClazzroomExcel(reader io.Reader) (r []*models.Clazzroom) {
	rows, err := GetSheet1Rows(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		var err error
		if i == 0 {
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			continue
		}
		id, err := strconv.Atoi(strings.TrimSpace(row[0]))
		if err != nil {
			log.Println(err)
			continue
		}
		building := strings.TrimSpace(row[1])
		room := strings.TrimSpace(row[2])
		r = append(r, &models.Clazzroom{
			Id:       id,
			Building: building,
			Room:     room,
		})
	}
	return
}

func ParseInstructExcel(reader io.Reader) (r []*models.Instruct) {
	rows, err := GetSheet1Rows(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		var err error
		if i == 0 {
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			continue
		}
		id, err := strconv.Atoi(strings.TrimSpace(row[0]))
		if err != nil {
			log.Println(err)
			continue
		}
		teacherID := strings.TrimSpace(row[1])
		// teacherName := strings.TrimSpace(row[2])
		courseID := strings.TrimSpace(row[3])
		// courseName := strings.TrimSpace(row[4])
		// semesterName := strings.TrimSpace(row[5])
		semesterDate := strings.TrimSpace(row[6])
		r = append(r, &models.Instruct{
			InstructId: id,
			Teacher:    &models.Teacher{Id: teacherID},
			Course:     &models.Course{Id: courseID},
			Semester:   &models.Semester{StartDate: semesterDate},
		})
	}
	return
}
func ParseInstructedClazzExcel(reader io.Reader) (r []*models.InstructedClazz) {
	rows, err := GetSheet1Rows(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i, row := range rows {
		var err error
		if i == 0 {
			continue
		}
		if strings.TrimSpace(row[0]) == "" {
			continue
		}
		id, err := strconv.Atoi(strings.TrimSpace(row[0]))
		if err != nil {
			log.Println(err)
			continue
		}
		clazzID := strings.TrimSpace(row[1])
		// clazzName := strings.TrimSpace(row[2])
		instructID, err := strconv.Atoi(strings.TrimSpace(row[0]))
		if err != nil {
			log.Println(err)
			continue
		}
		// teacherID := strings.TrimSpace(row[4])
		// teacherName := strings.TrimSpace(row[5])
		// courseID := strings.TrimSpace(row[6])
		// courseName := strings.TrimSpace(row[7])
		// semesterName := strings.TrimSpace(row[8])
		// semesterDate := strings.TrimSpace(row[9])
		r = append(r, &models.InstructedClazz{
			Id:       id,
			Clazz:    &models.Clazz{ClazzId: clazzID},
			Instruct: &models.Instruct{InstructId: instructID},
		})
	}
	return
}
