package views

import (
	"courseScheduling/scheduling"
)

type EvaluatorList []struct {
	Key     string `json:"key"`
	Explain string `json:"explain"`
}

func ToEvaluatorList(keylist []*scheduling.EvaluatorInfo) (l EvaluatorList) {
	for _, info := range keylist {
		l = append(l, struct {
			Key     string `json:"key"`
			Explain string `json:"explain"`
		}{
			Key:     info.Key,
			Explain: info.Explain,
		})
	}
	return
}
