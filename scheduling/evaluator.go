package scheduling

import (
	"log"
	"sort"
	"strings"

	"github.com/montanaflynn/stats"
)

type Evaluator func(X *GeneticSchedule) (fit, min, max float64, err error)
type EvaluatorInfo struct {
	Key     string
	Explain string
	Weight  float64
	F       Evaluator
}

var (
	EvaluatorsTable = map[string]*EvaluatorInfo{
		"AvoidUseNight": {
			Explain: "避免必修课和考试课使用晚上的时间段",
			Weight:  950.0,
			F:       AvoidUseNight,
		},
		"DisperseSameCourse": {
			Explain: "同一门课在一周内尽量分散",
			Weight:  450.0,
			F:       DisperseSameCourse,
		},
		"KeepAllLessonsDisperseEveryTimespan": {
			Explain: "以时间段为单位计算课程数标准差，确保课程分散性",
			Weight:  50.0,
			F:       KeepAllLessonsDisperseEveryTimespan,
		},
		"KeepAllLessonsDisperseEveryDay": {
			Explain: "以天为单位计算课程数标准差，确保课程分散性",
			Weight:  50.0,
			F:       KeepAllLessonsDisperseEveryDay,
		},
	}
	EvaluatorList []*EvaluatorInfo
)

func init() {
	EvaluatorList = make([]*EvaluatorInfo, 0, len(EvaluatorsTable))
	for key := range EvaluatorsTable {
		EvaluatorsTable[key].Key = key
		EvaluatorList = append(EvaluatorList, EvaluatorsTable[key])
	}
	sort.SliceStable(EvaluatorList, func(i, j int) bool {
		return strings.Compare(EvaluatorList[i].Key, EvaluatorList[j].Key) == -1
	})
}

// 1. 避免必修课和考试课使用晚上的时间段
func AvoidUseNight(X *GeneticSchedule) (fit, min, max float64, err error) {
	costTimespan := 0.0
	cntKaoshi := 0
	for i := range X.items {
		if X.items[i].TendToObtainDayTimespan {
			timespanID := X.items[i].TimespanID
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
	return costTimespan, minCostTimespan, maxCostTimespan, nil
}

// 2. 同一门课在一周内尽量分散
func DisperseSameCourse(X *GeneticSchedule) (fit, min, max float64, err error) {
	var (
		maxDayOfWeek = 5
		maxTimespan  = len(X.allTimespan)
	)
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
			theTime := (items[i].DayOfWeek-1)*maxTimespan + items[i].TimespanID
			prevTime := (items[i-1].DayOfWeek-1)*maxTimespan + items[i-1].TimespanID
			distance := theTime - prevTime
			oneDensity += 1.0 - float64(distance)/maxDistance
		}
		oneDensity /= float64(len(items))
		costSameICDensity += oneDensity
		cnt++
	}
	costSameICDensity /= float64(cnt)
	return costSameICDensity, 0.0, float64(maxDayOfWeek * maxTimespan), nil
}

// 3. 以时间段为单位计算课程数标准差，确保课程分散性。
func KeepAllLessonsDisperseEveryTimespan(X *GeneticSchedule) (fit, min, max float64, err error) {
	var (
		maxDayOfWeek = 5
		maxTimespan  = len(X.allTimespan)
	)
	var (
		minSdevCntInTD = 0.0
		maxSdevCntInTD = float64(maxDayOfWeek*maxTimespan+1) / 2
	)
	lenT := len(X.allTimespan)
	lenD := len(X.availableWeekday)
	cntInTD := make([]float64, lenT*lenD)
	for i := range X.items {
		cntInTD[(X.items[i].DayOfWeek-1)+lenD*(X.items[i].TimespanID-1)]++
	}
	sdevCntInTD, err := stats.StandardDeviation(cntInTD)
	if err != nil {
		log.Println(err)
		return 0.0, minSdevCntInTD, maxSdevCntInTD, err
	}
	return sdevCntInTD, minSdevCntInTD, maxSdevCntInTD, nil
}

// 4. 以天为单位计算课程数标准差，确保课程分散性。
func KeepAllLessonsDisperseEveryDay(X *GeneticSchedule) (fit, min, max float64, err error) {
	var (
		minSdevCntInD = 0.0
		maxSdevCntInD = float64(X.parent.cntLessons)
	)
	lenD := len(X.availableWeekday)
	cntInD := make([]float64, lenD)
	for i := range X.items {
		cntInD[X.items[i].DayOfWeek-1]++
	}
	sdevCntInD, err := stats.StandardDeviation(cntInD)
	if err != nil {
		log.Println(err)
		return 0.0, minSdevCntInD, maxSdevCntInD, err
	}
	return sdevCntInD, minSdevCntInD, maxSdevCntInD, nil
}
