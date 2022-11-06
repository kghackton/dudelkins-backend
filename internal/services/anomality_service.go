package services

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"dudelkins/internal/consts"
	"dudelkins/internal/interfaces"
	"dudelkins/internal/objects/bo"
	"dudelkins/pkg/utils"
)

type AnomalityService struct {
	AnomalyCheckers []AnomalyClassCheck
}

func NewAnomalityService(applicationService interfaces.IApplicationViewService, defectIdsDurationMap map[int]*time.Duration, defectIdsDeviationMap map[string]map[int]time.Duration, insService interfaces.IInsService) *AnomalityService {
	return &AnomalityService{AnomalyCheckers: []AnomalyClassCheck{
		NewFastCloseAnomalyCheck(),
		NewClosedWithoutCompletionCheck(applicationService, defectIdsDurationMap),
		NewClosedWithCompletionWithoutReturningsCheck(applicationService, defectIdsDurationMap),
		NewBadReviewCheck(),
		NewWithReturningsCheck(),
		NewClosedForLessThan10MinutesWithNoReturningsCheck(applicationService),
		NewDeviationCheck(defectIdsDeviationMap),
		NewDudelkINSCheckCheck(insService),
	}}
}

func (s *AnomalityService) CheckForAnomalies(application bo.Application) (anomalies map[string]bo.AnomalyClass, err error) {
	anomalies = make(map[string]bo.AnomalyClass, len(s.AnomalyCheckers))
	for _, anomalyChecker := range s.AnomalyCheckers {
		verdict, class, description, err := anomalyChecker.CheckApplication(application)
		if err != nil {
			return nil, errors.Wrap(errors.WithMessagef(err, "class: %s", class), "CheckForAnomalies")
		}

		if verdict == true {
			anomalies[class] = bo.AnomalyClass{
				Verdict:     verdict,
				Description: description,
			}
		}
	}
	return
}

type AnomalyClassCheck interface {
	CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error)
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

func (c FastCloseAnomalyCheck) CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error) {
	if application.ResultCode == consts.ResultCodeResolved && application.AmountOfReturnings == nil {
		if subDuration := application.ClosedAt.Sub(application.CreatedAt); subDuration < time.Minute*10 {
			if !utils.OneOf(application.CategoryId, c.ListOfCategoryIdsThatCanBeClosedFast) {
				return true, c.Class, fmt.Sprintf("closed too fast. categoryId: %d closed for: %s", application.CategoryId, subDuration), err
			}
		}
	}

	return false, c.Class, "", err
}

type ClosedWithoutCompletionCheck struct {
	Class                   string
	ExceptDefectIds         []int
	ExceptRenderedServiceId int

	DefectIdsDuration map[int]*time.Duration

	applicationService interfaces.IApplicationViewService
}

func NewClosedWithoutCompletionCheck(applicationService interfaces.IApplicationViewService, defectIdsDurationMap map[int]*time.Duration) ClosedWithoutCompletionCheck {
	return ClosedWithoutCompletionCheck{
		Class:                   "closed without completion for same applicant",
		ExceptDefectIds:         []int{7906, 7907},
		ExceptRenderedServiceId: 18268,
		DefectIdsDuration:       defectIdsDurationMap,
		applicationService:      applicationService,
	}
}

func (c ClosedWithoutCompletionCheck) CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error) {
	if application.ResultCode != consts.ResultCodeResolved &&
		!utils.OneOf(c.ExceptRenderedServiceId, application.RenderedServicesIds) &&
		!utils.OneOf(application.DefectId, c.ExceptDefectIds) {
		if defectDuration, exists := c.DefectIdsDuration[application.DefectId]; exists && defectDuration != nil {
			closedFrom := application.CreatedAt.Add(-*defectDuration)
			opts := &bo.ApplicationRetrieveOpts{
				ClosedFrom: &closedFrom,
				ClosedTo:   &application.CreatedAt,

				DefectIds: []int{application.DefectId},
				UNOM:      &application.UNOM,
				Entrance:  application.Entrance,
				Floor:     application.Floor,
				Flat:      application.Flat,
			}
			applications, err := c.applicationService.Get(context.Background(), opts)
			if err != nil {
				return false, c.Class, "", errors.Wrap(err, "CheckApplication")
			}
			if len(applications) > 0 {
				applicationIds := make([]int32, 0, len(applications))
				for _, application := range applications {
					applicationIds = append(applicationIds, application.RootId)
				}
				return true, c.Class, fmt.Sprintf("applicationIds: %+v", applicationIds), nil
			}
		}
	}
	return false, c.Class, "", err
}

type ClosedWithCompletionWithoutReturningsCheck struct {
	Class           string
	ExceptDefectIds []int

	DefectIdsDuration map[int]*time.Duration

	applicationService interfaces.IApplicationViewService
}

func NewClosedWithCompletionWithoutReturningsCheck(applicationService interfaces.IApplicationViewService, defectIdsDurationMap map[int]*time.Duration) ClosedWithCompletionWithoutReturningsCheck {
	return ClosedWithCompletionWithoutReturningsCheck{
		Class:              "closed with completion but without returnings for same applicant",
		ExceptDefectIds:    []int{7906, 7907},
		DefectIdsDuration:  defectIdsDurationMap,
		applicationService: applicationService,
	}
}

func (c ClosedWithCompletionWithoutReturningsCheck) CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error) {
	if application.ResultCode == consts.ResultCodeResolved &&
		!utils.OneOf(application.DefectId, c.ExceptDefectIds) &&
		application.ClosedAt.Sub(application.CreatedAt) > time.Minute*10 &&
		application.AmountOfReturnings == nil {
		if defectDuration, exists := c.DefectIdsDuration[application.DefectId]; exists && defectDuration != nil {
			closedFrom := application.CreatedAt.Add(-*defectDuration)
			opts := &bo.ApplicationRetrieveOpts{
				ClosedFrom: &closedFrom,
				ClosedTo:   &application.CreatedAt,

				DefectIds: []int{application.DefectId},
				UNOM:      &application.UNOM,
				Entrance:  application.Entrance,
				Floor:     application.Floor,
				Flat:      application.Flat,
			}
			applications, err := c.applicationService.Get(context.Background(), opts)
			if err != nil {
				return false, c.Class, "", errors.Wrap(err, "CheckApplication")
			}
			if len(applications) > 0 {
				applicationIds := make([]int32, 0, len(applications))
				for _, application := range applications {
					applicationIds = append(applicationIds, application.RootId)
				}
				return true, c.Class, fmt.Sprintf("applicationIds: %+v", applicationIds), nil
			}
		}
	}
	return false, c.Class, "", err
}

type BadReviewCheck struct {
	Class string
}

func NewBadReviewCheck() BadReviewCheck {
	return BadReviewCheck{
		Class: "bad review",
	}
}

func (c BadReviewCheck) CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error) {
	if application.RatingCode != nil && *application.RatingCode == consts.BadReviewRatingCode {
		return true, c.Class, "", nil
	}

	return false, c.Class, "", err
}

type ClosedForLessThan10MinutesWithNoReturningsCheck struct {
	Class                                string
	ListOfCategoryIdsThatCanBeClosedFast []int
	applicationService                   interfaces.IApplicationViewService
}

func NewClosedForLessThan10MinutesWithNoReturningsCheck(applicationService interfaces.IApplicationViewService) ClosedForLessThan10MinutesWithNoReturningsCheck {
	return ClosedForLessThan10MinutesWithNoReturningsCheck{
		Class:              "closed for less than 10 minutes with no returnings",
		applicationService: applicationService,
		ListOfCategoryIdsThatCanBeClosedFast: []int{
			2303, 2245, 1903, 2396, 1922, 1771, 2096, 7907, 7906,
		},
	}
}

func (c ClosedForLessThan10MinutesWithNoReturningsCheck) CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error) {
	if utils.OneOf(application.DefectId, c.ListOfCategoryIdsThatCanBeClosedFast) &&
		application.AmountOfReturnings == nil {
		closedFrom := application.CreatedAt.Add(-time.Hour)
		amountOfReturningsLessThan := 1
		opts := &bo.ApplicationRetrieveOpts{
			ClosedFrom: &closedFrom,
			ClosedTo:   &application.CreatedAt,

			DefectIds:                 []int{application.DefectId},
			UNOM:                      &application.UNOM,
			Entrance:                  application.Entrance,
			Floor:                     application.Floor,
			Flat:                      application.Flat,
			AmoutOfReturningsLessThan: &amountOfReturningsLessThan,
		}
		applications, err := c.applicationService.Get(context.Background(), opts)
		if err != nil {
			return false, c.Class, "", errors.Wrap(err, "CheckApplication")
		}
		if len(applications) > 0 {
			applicationIds := make([]int32, 0, len(applications))
			for _, application := range applications {
				applicationIds = append(applicationIds, application.RootId)
			}
			return true, c.Class, fmt.Sprintf("applicationIds: %+v", applicationIds), nil
		}
	}
	return false, c.Class, "", err
}

type WithReturningsCheck struct {
	Class     string
	DefectIds map[int]struct{}
}

func NewWithReturningsCheck() WithReturningsCheck {
	return WithReturningsCheck{
		Class: "with returnings",
		DefectIds: map[int]struct{}{
			1853:  {},
			2259:  {},
			1677:  {},
			2517:  {},
			40150: {},
			40153: {},
			1705:  {},
			2011:  {},
			2333:  {},
			40138: {},
			2171:  {},
			1688:  {},
			1696:  {},
			2098:  {},
			39271: {},
			1695:  {},
			2184:  {},
			3963:  {},
			2233:  {},
			2244:  {},
			38008: {},
			38009: {},
			2222:  {},
			2162:  {},
			1939:  {},
			2196:  {},
			2053:  {},
			13161: {},
			1649:  {},
			38562: {},
			38567: {},
			1614:  {},
			38581: {},
			38583: {},
			2387:  {},
			38585: {},
			38578: {},
			2392:  {},
			1604:  {},
			2423:  {},
			2425:  {},
			2421:  {},
			1618:  {},
			1664:  {},
			18442: {},
			2207:  {},
			1630:  {},
			1622:  {},
			1701:  {},
			38216: {},
			38217: {},
			1665:  {},
			2220:  {},
			2159:  {},
			1620:  {},
			40141: {},
			1610:  {},
			2158:  {},
			2282:  {},
			2213:  {},
			2253:  {},
			4075:  {},
			2262:  {},
			2221:  {},
			2263:  {},
			40144: {},
			37735: {},
			2082:  {},
			39017: {},
			39015: {},
			2177:  {},
			2170:  {},
			38554: {},
			40135: {},
			1985:  {},
			1698:  {},
			1974:  {},
			1694:  {},
			1683:  {},
			2246:  {},
			38577: {},
			2150:  {},
			2175:  {},
		},
	}
}

func (c WithReturningsCheck) CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error) {
	if application.AmountOfReturnings != nil && *application.AmountOfReturnings > 0 {
		if _, exists := c.DefectIds[application.DefectId]; exists {
			return true, c.Class, "", nil
		}
	}

	return false, c.Class, "", err
}

type DeviationCheck struct {
	Class string

	DefectIdsDeviation map[string]map[int]time.Duration
}

func NewDeviationCheck(defectIdsDeviationMap map[string]map[int]time.Duration) DeviationCheck {
	return DeviationCheck{
		Class:              "deviation",
		DefectIdsDeviation: defectIdsDeviationMap,
	}
}

func (c DeviationCheck) CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error) {
	expectedTimeDuration, exists := c.DefectIdsDeviation[application.ResultCode][application.DefectId]
	if !exists {
		return false, c.Class, "", err
	}

	applicationDuration := application.ClosedAt.Sub(application.CreatedAt)

	if applicationDuration > expectedTimeDuration {
		return true, c.Class, fmt.Sprintf("applicationDuration: %s, expected: %s", applicationDuration, expectedTimeDuration), err
	}

	return false, c.Class, "", err
}

type DudelkINSCheck struct {
	Class string

	InsService interfaces.IInsService
}

func NewDudelkINSCheckCheck(insService interfaces.IInsService) DudelkINSCheck {
	return DudelkINSCheck{
		Class:      "DudelkINS",
		InsService: insService,
	}
}

func (c DudelkINSCheck) CheckApplication(application bo.Application) (isAbnormal bool, class string, description string, err error) {
	isAbnormal, confidence, err := c.InsService.IsAbnormal(application)
	if err != nil {
		return isAbnormal, c.Class, "", errors.Wrap(err, "CheckApplication DudelkINS")
	}

	if isAbnormal {
		return isAbnormal, c.Class, fmt.Sprintf("confidence: %f", confidence), err
	}

	return false, c.Class, "", err
}
