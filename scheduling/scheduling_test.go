package scheduling

import (
	"testing"

	"courseScheduling/dev/dummy"
	"courseScheduling/models"
)

func TestNewSchedule(t *testing.T) {
	type args struct {
		allInstructs []*models.Instruct
		allClazz     []*models.Clazz
		allClazzroom []*models.Clazzroom
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test dummy data", args: args{
			allInstructs: dummy.ParseInstruct(),
			allClazz:     dummy.ParseClazz(),
			allClazzroom: dummy.ParseClazzroom(),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewSchedule(tt.args.allInstructs, tt.args.allClazz, tt.args.allClazzroom)
			for _, item := range result {
				t.Log(item)
			}
		})
	}
}
