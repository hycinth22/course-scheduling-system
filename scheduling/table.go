package scheduling

import (
	"log"
	"strings"

	"courseScheduling/models"
	"github.com/modood/table"
)

type TableRow struct {
	Monday    string `table:"  周一"`
	Tuesday   string `table:"  周二"`
	Wednesday string `table:"  周三"`
	Thursday  string `table:"  周四"`
	Friday    string `table:"  周五"`
	Saturday  string `table:"  周六"`
	Sunday    string `table:"  周日"`
}

func GeneticSchedule2String(s *GeneticSchedule) string {
	// [timespan][week][index]
	t := make([][][]*models.ScheduleItem, len(s.allTimespan)+1) // 1-len(s.allTimespan), 0 is unused
	for i := range t {
		t[i] = make([][]*models.ScheduleItem, 8) // 1-7, 0 is unused
	}
	for _, item := range s.items {
		if item.TimespanId <= 0 || item.TimespanId > len(s.allTimespan) {
			log.Println("invalid TimespanId")
			continue
		}
		if item.DayOfWeek < 1 || item.DayOfWeek > 7 {
			log.Println("invalid DayOfWeek")
			continue
		}
		t[item.TimespanId][item.DayOfWeek] = append(t[item.TimespanId][item.DayOfWeek], item)
	}

	lenT := len(s.allTimespan)
	rows := make([]TableRow, 0, 2*lenT)
	iRow := 0
	for iT := 1; iT <= lenT; iT++ {
		rows = append(rows, TableRow{})
		if len(t[iT][1]) > 0 {
			rows[iRow].Monday = generateLessonsListString(t[iT][1])
		}
		if len(t[iT][2]) > 0 {
			rows[iRow].Tuesday = generateLessonsListString(t[iT][2])
		}
		if len(t[iT][3]) > 0 {
			rows[iRow].Wednesday = generateLessonsListString(t[iT][3])
		}
		if len(t[iT][4]) > 0 {
			rows[iRow].Thursday = generateLessonsListString(t[iT][4])
		}
		if len(t[iT][5]) > 0 {
			rows[iRow].Friday = generateLessonsListString(t[iT][5])
		}
		if len(t[iT][6]) > 0 {
			rows[iRow].Saturday = generateLessonsListString(t[iT][6])
		}
		if len(t[iT][7]) > 0 {
			rows[iRow].Sunday = generateLessonsListString(t[iT][7])
		}
		// rows = append(rows, splitRow)
		iRow++
	}
	return table.Table(rows)
}

func generateLessonsListString(a []*models.ScheduleItem) string {
	s := make([]string, len(a))
	for i := range a {
		s[i] = a[i].String()
	}
	return strings.Join(s, ", ")
}
