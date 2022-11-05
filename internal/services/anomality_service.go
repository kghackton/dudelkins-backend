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

func NewAnomalityService(applicationService interfaces.IApplicationViewService, defectIdsDurationMap map[int]*time.Duration) *AnomalityService {
	return &AnomalityService{AnomalyCheckers: []AnomalyClassCheck{
		NewFastCloseAnomalyCheck(),
		NewClosedWithoutCompletionCheck(applicationService, defectIdsDurationMap),
		NewClosedWithCompletionWithoutReturningsCheck(applicationService, defectIdsDurationMap),
		NewBadReviewCheck(),
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
