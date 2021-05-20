package scheduling

import (
	"fmt"
	"log"

	"github.com/montanaflynn/stats"
)

type ScoreHard [3]int

func (s ScoreHard) String() string {
	return fmt.Sprintf("[%d %d %d]", s[0], s[1], s[2])
}

type ScoreSoft [4]float64

func (s ScoreSoft) String() string {
	return fmt.Sprintf("[%.4f %.4f %.4f %.4f]", s[0], s[1], s[2], s[3])
}

// Evaluate a Schedule
func (X *GeneticSchedule) Evaluate() (fit float64, err error) {
	var (
		maxDayOfWeek = 5
		maxTimespan  = 8
	)
	// ！！！硬约束条件检测开始！！！
	hardTotal := X.Invalidity()
	fit += 100000.0 * float64(hardTotal)
	//if hardTotal > 0.0 {
	//	return fit, nil
	//}
	// ！！！软约束条件检测开始！！！
	const (
		softScoreMax1     = 950.0
		softScoreMax2     = 450.0
		softScoreMax3     = 50.0
		softScoreMax4     = 50.0
		softScoreMaxTotal = softScoreMax1 + softScoreMax2 + softScoreMax3 + softScoreMax4
	)
	softTotal := 0.0
	// 1. 避免必修课和考试课使用晚上的时间段
	costTimespan := 0.0
	cntKaoshi := 0
	for i := range X.items {
		if X.items[i].Instruct.Course.Kind == "必修" || X.items[i].Instruct.Course.ExamMode == "考试" {
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
	var maxDistance = float64(maxDayOfWeek * maxTimespan)
	costSameICDensity := 0.0
	var cnt = 0
	for _, items := range X.queryByInstructedClazz {
		var oneDensity = 0.0
		if len(items) == 0 {
			continue
		}
		if len(items) == 1 {
			// 在长为totalDistance的线段上，任取两点的距离的数学期望为1/3*maxDistance
			// oneDensity += 1.0/3.0
			continue
		}
		for i := 1; i < len(items); i++ {
			theTime := (items[i].DayOfWeek-1)*maxTimespan + items[i].TimespanId
			prevTime := (items[i-1].DayOfWeek-1)*maxTimespan + items[i-1].TimespanId
			distance := theTime - prevTime
			oneDensity += 1.0 - float64(distance)/maxDistance
		}
		oneDensity /= float64(len(items))
		costSameICDensity += oneDensity
		cnt++
	}
	costSameICDensity /= float64(cnt)
	// 归一化
	var (
		minSameICDensity = 0.0
		maxSameICDensity = float64(maxDayOfWeek * maxTimespan)
	)
	uniSameICDensity := normalizeFloat64(costSameICDensity, minSameICDensity, maxSameICDensity)
	// log.Println("uniSameICDensity:", uniSameICDensity)
	softTotal += uniSameICDensity * softScoreMax2
	X.scores.s[1] = uniSameICDensity

	// 3. 以时间段为单位计算课程数标准差，确保课程分散性。
	lenT := len(X.allTimespan)
	lenD := len(X.availableWeekday)
	cntInTD := make([]float64, lenT*lenD)
	for i := range X.items {
		cntInTD[(X.items[i].DayOfWeek-1)+lenD*(X.items[i].TimespanId-1)]++
	}
	sdevCntInTD, err := stats.StandardDeviation(cntInTD)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}
	// 归一化
	var (
		minSdevCntInTD = 0.0
		maxSdevCntInTD = float64(maxDayOfWeek*maxTimespan+1) / 2
	)
	uniSdevCntInTD := normalizeFloat64(sdevCntInTD, minSdevCntInTD, maxSdevCntInTD)
	// log.Println("uniSdev:", uniSdev)
	softTotal += uniSdevCntInTD * softScoreMax3
	X.scores.s[2] = uniSdevCntInTD

	// 4. 以天为单位计算课程数标准差，确保课程分散性。
	cntInD := make([]float64, lenD)
	for i := range X.items {
		cntInD[X.items[i].DayOfWeek-1]++
	}
	sdevCntInD, err := stats.StandardDeviation(cntInD)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}
	// 归一化
	var (
		minSdevCntInD = 0.0
		maxSdevCntInD = float64(X.parent.cntLessons)
	)
	uniSdevCntInD := normalizeFloat64(sdevCntInD, minSdevCntInD, maxSdevCntInD)
	// log.Println("uniSdev:", uniSdev)
	softTotal += uniSdevCntInD * softScoreMax4
	X.scores.s[3] = uniSdevCntInD

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
