package scheduling

import (
	"fmt"

	"courseScheduling/models"
)

func NewSchedule(allInstructs []*models.Instruct, allClazz []*models.Clazz, allClazzroom []*models.Clazzroom) []*models.ScheduleItem {
	fmt.Println("allInstructs:")
	for _, item := range allInstructs {
		fmt.Printf("%#v", item)
	}
	fmt.Println()

	fmt.Println("allClazz:")
	for _, item := range allClazz {
		fmt.Printf("%#v", item)
	}
	fmt.Println()

	fmt.Println("allClazzroom:")
	for _, item := range allClazzroom {
		fmt.Printf("%#v", item)
	}
	fmt.Println()

	return nil
}
