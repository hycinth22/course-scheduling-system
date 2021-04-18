package excel

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"courseScheduling/models"
	"courseScheduling/views"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/pkg/errors"
)

func formatCell1(r *models.ScheduleItem) string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("%v(%v)\n%v%v\n%v%v",
		r.Instruct.Course.Name, r.Instruct.Course.Kind,
		r.Instruct.Teacher.Name, r.Instruct.Teacher.Title,
		r.Clazzroom.Building, r.Clazzroom.Room)
}

func formatCell2(r *models.ScheduleItem) string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("%v(%v)\n%v%v\n%v%v",
		r.Instruct.Teacher.Name, r.Instruct.Teacher.Title,
		r.Instruct.Course.Name, r.Clazz.ClazzName,
		r.Clazzroom.Building, r.Clazzroom.Room)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GenStudentTables(writer io.Writer, view *views.ScheduleItemsTableView, clazzes []*models.Clazz) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.Errorf("%v", x)
		}
	}()
	f := excelize.NewFile()
	for _, clazz := range clazzes {
		sheetName := clazz.ClazzName
		f.NewSheet(sheetName)
		// 表头
		_ = f.SetCellStr(sheetName, "A1", "时间段")
		checkError(f.SetCellStr(sheetName, "B1", "周一"))
		checkError(f.SetCellStr(sheetName, "C1", "周二"))
		checkError(f.SetCellStr(sheetName, "D1", "周三"))
		checkError(f.SetCellStr(sheetName, "E1", "周四"))
		checkError(f.SetCellStr(sheetName, "F1", "周五"))
		checkError(f.SetCellStr(sheetName, "G1", "周六"))
		checkError(f.SetCellStr(sheetName, "H1", "周日"))
		for i, r := range view.ByClazz[clazz.ClazzId] {
			rowIndex := strconv.Itoa(i + 2)
			checkError(f.SetCellStr(sheetName, "A"+rowIndex, strconv.Itoa(i)))
			checkError(f.SetCellStr(sheetName, "B"+rowIndex, formatCell1(r[1])))
			checkError(f.SetCellStr(sheetName, "C"+rowIndex, formatCell1(r[2])))
			checkError(f.SetCellStr(sheetName, "D"+rowIndex, formatCell1(r[3])))
			checkError(f.SetCellStr(sheetName, "E"+rowIndex, formatCell1(r[4])))
			checkError(f.SetCellStr(sheetName, "F"+rowIndex, formatCell1(r[5])))
			checkError(f.SetCellStr(sheetName, "G"+rowIndex, formatCell1(r[6])))
			checkError(f.SetCellStr(sheetName, "H"+rowIndex, formatCell1(r[7])))
		}
	}
	f.DeleteSheet("Sheet1")
	err = f.Write(writer)
	return err
}

func formatJoin(items []*models.ScheduleItem, formatter func(r *models.ScheduleItem) string) string {
	var r []string
	for _, item := range items {
		r = append(r, formatter(item))
	}
	return strings.Join(r, "\n\n")
}

func GenTeacherTables(writer io.Writer, view *views.ScheduleItemsTableView, depts []*models.Department) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = errors.Errorf("%v", x)
		}
	}()
	f := excelize.NewFile()
	for _, dept := range depts {
		sheetName := dept.DeptName
		f.NewSheet(sheetName)
		// 表头
		checkError(f.SetCellStr(sheetName, "A1", "时间段"))
		checkError(f.SetCellStr(sheetName, "B1", "周一"))
		checkError(f.SetCellStr(sheetName, "C1", "周二"))
		checkError(f.SetCellStr(sheetName, "D1", "周三"))
		checkError(f.SetCellStr(sheetName, "E1", "周四"))
		checkError(f.SetCellStr(sheetName, "F1", "周五"))
		checkError(f.SetCellStr(sheetName, "G1", "周六"))
		checkError(f.SetCellStr(sheetName, "H1", "周日"))
		for i, r := range view.ByDept[dept.DeptId] {
			rowIndex := strconv.Itoa(i + 2)
			checkError(f.SetCellStr(sheetName, "A"+rowIndex, strconv.Itoa(i)))
			checkError(f.SetCellStr(sheetName, "B"+rowIndex, formatJoin(r[1], formatCell2)))
			checkError(f.SetCellStr(sheetName, "C"+rowIndex, formatJoin(r[2], formatCell2)))
			checkError(f.SetCellStr(sheetName, "D"+rowIndex, formatJoin(r[3], formatCell2)))
			checkError(f.SetCellStr(sheetName, "E"+rowIndex, formatJoin(r[4], formatCell2)))
			checkError(f.SetCellStr(sheetName, "F"+rowIndex, formatJoin(r[5], formatCell2)))
			checkError(f.SetCellStr(sheetName, "G"+rowIndex, formatJoin(r[6], formatCell2)))
			checkError(f.SetCellStr(sheetName, "H"+rowIndex, formatJoin(r[7], formatCell2)))
		}
	}
	f.DeleteSheet("Sheet1")
	err = f.Write(writer)
	return err
}
