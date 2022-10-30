package bo

import (
	"fmt"
	"time"

	"dudelkins/internal/consts"
	"dudelkins/pkg/utils"
)

type AnomalyClassCheck interface {
	CheckApplication(application Application) (isAbnormal bool, class string, description string)
}

type FastCloseAnomalyCheck struct {
	Class                                string
	ListOfCategoryIdsThatCanBeClosedFast []int
}

func NewFastCloseAnomalyCheck() FastCloseAnomalyCheck {
	return FastCloseAnomalyCheck{
		Class: "closed too fast",
		ListOfCategoryIdsThatCanBeClosedFast: []int{
			2303, 2245, 1903, 2396, 1922, 1771, 2096, 7907, 7906,
		},
	}
}

func (c FastCloseAnomalyCheck) CheckApplication(application Application) (isAbnormal bool, class string, description string) {
	if application.ResultCode == consts.ResultCodeResolved && application.AmountOfReturnings == nil {
		if subDuration := application.ClosedAt.Sub(application.CreatedAt); subDuration < time.Minute*10 {
			if !utils.OneOf(application.CategoryId, c.ListOfCategoryIdsThatCanBeClosedFast) {
				return true, c.Class, fmt.Sprintf("closed too fast. categoryId: %d closed for: %s", application.CategoryId, subDuration)
			}
		}
	}

	return false, c.Class, ""
}
