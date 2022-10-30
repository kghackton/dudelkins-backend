package services

import "dudelkins/internal/objects/bo"

type AnomalityService struct {
	AnomalyCheckers []bo.AnomalyClassCheck
}

func NewAnomalityService() *AnomalityService {
	return &AnomalityService{AnomalyCheckers: []bo.AnomalyClassCheck{
		bo.NewFastCloseAnomalyCheck(),
	}}
}

func (s *AnomalityService) CheckForAnomalies(application bo.Application) (anomalies map[string]bo.AnomalyClass) {
	anomalies = make(map[string]bo.AnomalyClass, len(s.AnomalyCheckers))
	for _, anomalyChecker := range s.AnomalyCheckers {
		verdict, class, description := anomalyChecker.CheckApplication(application)

		anomalies[class] = bo.AnomalyClass{
			Verdict:     verdict,
			Description: description,
		}
	}
	return
}
