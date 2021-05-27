package scheduling

import (
	"fmt"
	"log"
	"strings"

	"github.com/montanaflynn/stats"
)

type ScoreHard [3]int

func (s ScoreHard) String() string {
	return fmt.Sprintf("[%d %d %d]", s[0], s[1], s[2])
}

type ScoreSoft struct {
	Items  []string
	Scores map[string]float64
}

func (s ScoreSoft) String() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	for i := range s.Items {
		itemName := s.Items[i]
		builder.WriteString(fmt.Sprintf("%.4f ", s.Scores[itemName]))
	}
	builder.WriteString("]")
	return builder.String()
}

// Evaluate a Schedule
func (X *GeneticSchedule) Evaluate() (fit float64, err error) {
	// ！！！硬约束条件检测开始！！！
	hardTotal := X.Invalidity()
	if hardTotal > 0.0 {
		fit += 100000.0 * float64(hardTotal)
	}
	// ！！！软约束条件检测开始！！！
	if len(X.parent.Params.UseEvaluator) == 0 {
		return fit, nil
	}
	softTotal := 0.0
	softScoreMaxTotal := 0.0
	for _, evaluatorName := range X.parent.Params.UseEvaluator {
		evaluator := EvaluatorsTable[evaluatorName]
		score, minScore, maxScore, err := evaluator.F(X)
		if err != nil {
			log.Println(err)
			return 0.0, err
		}
		// 归一化后加权
		uniScore := normalizeFloat64(score, minScore, maxScore)
		X.scores.s.Scores[evaluatorName] = uniScore
		weighted := uniScore * evaluator.Weight
		softTotal += weighted
		softScoreMaxTotal += evaluator.Weight
	}
	uniSoft := normalizeFloat64(softTotal, 0.0, softScoreMaxTotal) * 10000
	fit += uniSoft
	return fit, nil
}

func (X *GeneticSchedule) CalcSdevForTD() (sdev float64, err error) {
	lenT := len(X.allTimespan)
	lenD := len(X.availableWeekday)
	cntInTD := make([]float64, lenT*lenD)
	for i := range X.items {
		cntInTD[(X.items[i].DayOfWeek-1)+lenD*(X.items[i].TimespanID-1)]++
	}
	sdev, err = stats.StandardDeviation(cntInTD)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}
	return sdev, err
}
