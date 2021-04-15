package scheduling

import (
	"log"
	"sort"

	"github.com/montanaflynn/stats"
)

// Evaluate a Schedule
func (X *GeneticSchedule) Evaluate() (fit float64, err error) {
	var (
		maxDayOfWeek = 5
		maxTimespan  = 8
	)
	// ！！！硬约束条件检测开始！！！
	hardTotal := X.Invalidity()
	if hardTotal > 0 {
		fit += 100000.0 * hardTotal
	}
	// ！！！软约束条件检测开始！！！
	// 共900分
	const (
		softScoreMaxTotal = 900.0
		softScoreMax1     = 350.0
		softScoreMax2     = 350.0
		softScoreMax3     = 200.0
	)
	softTotal := 0.0
	// 1. 避免考试课使用晚上的时间段
	costTimespan := 0.0
	cntKaoshi := 0
	for i := range X.items {
		if X.items[i].Instruct.Course.ExamMode == "考试" {
			timespanID := X.items[i].TimespanId
			pri := X.parent.timespanMap[timespanID].Priority
			costTimespan += float64(pri)
			cntKaoshi++
		}
	}
	// 归一化
	var (
		minCostTimespan = 0.0
		maxCostTimespan = 0.0
	)
	for i := range X.allTimespan {
		if X.allTimespan[i].Priority > 0 {
			maxCostTimespan += float64(X.allTimespan[i].Priority) * float64(cntKaoshi)
		}
	}
	uniCostTimespan := normalizeFloat64(costTimespan, minCostTimespan, maxCostTimespan)
	// log.Println("uniCostTimespan:", uniCostTimespan)
	softTotal += uniCostTimespan * softScoreMax1
	X.scores.s[0] = uniCostTimespan

	// 2. 同一门课在一周内尽量分散
	costSameICDensity := 0.0
	var (
		totalDistance = float64(maxDayOfWeek * maxTimespan)
	)
	for _, items := range X.queryByInstructedClazz {
		var oneDensity = 0.0
		if len(items) == 0 {
			continue
		}
		if len(items) == 1 {
			// 在长为totalDistance的线段上，任取两点的距离的数学期望为totalDistance/3
			oneDensity += totalDistance / 3.0
			continue
		}
		sort.Slice(items, func(i, j int) bool {
			if items[i].DayOfWeek != items[j].DayOfWeek {
				return items[i].DayOfWeek < items[j].DayOfWeek
			}
			return items[i].TimespanId < items[j].TimespanId
		})
		for i := 1; i < len(items); i++ {
			theTime := (items[i].DayOfWeek-1)*maxTimespan + items[i].TimespanId
			prevTime := (items[i-1].DayOfWeek-1)*maxTimespan + items[i-1].TimespanId
			distance := theTime - prevTime
			oneDensity += float64(distance)
		}
		oneDensity /= float64(len(items))
		costSameICDensity += oneDensity
	}
	costSameICDensity /= float64(len(X.queryByInstructedClazz))
	// 归一化
	var (
		minSameICDensity = 0.0
		maxSameICDensity = float64(maxDayOfWeek * maxTimespan)
	)
	uniSameICDensity := 1.0 - normalizeFloat64(costSameICDensity, minSameICDensity, maxSameICDensity)
	// log.Println("uniSameICDensity:", uniSameICDensity)
	softTotal += uniSameICDensity * softScoreMax2
	X.scores.s[1] = uniSameICDensity

	// 3. 计算标准差，确保课程分散性。
	lenT := len(X.allTimespan)
	lenD := len(X.availableWeekday)
	cntInTD := make([]float64, lenT*lenD)
	for i := range X.items {
		cntInTD[(X.items[i].DayOfWeek-1)+lenD*(X.items[i].TimespanId-1)]++
	}
	sdev, err := stats.StandardDeviation(cntInTD)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}
	// 归一化
	var (
		minSdev = 0.0
		maxSdev = float64(maxDayOfWeek*maxTimespan+1) / 2
	)
	uniSdev := normalizeFloat64(sdev, minSdev, maxSdev)
	// log.Println("uniSdev:", uniSdev)
	softTotal += uniSdev * softScoreMax3
	X.scores.s[2] = uniSdev

	uniSoftTotal := normalizeFloat64(softTotal, 0, softScoreMaxTotal)
	fit += uniSoftTotal * 10000
	return
}

func (X *GeneticSchedule) CalcSdevForTD() (sdev float64, err error) {
	lenT := len(X.allTimespan)
	lenD := len(X.availableWeekday)
	cntInTD := make([]float64, lenT*lenD)
	for i := range X.items {
		cntInTD[(X.items[i].DayOfWeek-1)+lenD*(X.items[i].TimespanId-1)]++
	}
	sdev, err = stats.StandardDeviation(cntInTD)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}
	return sdev, err
}
