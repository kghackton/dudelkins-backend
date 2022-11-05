package infrastructure

import (
	"dudelkins/internal/repositories"
	"dudelkins/internal/services"
)

func (k *Kernel) InjectApplicationUploadService() *services.ApplicationUploadService {
	return &services.ApplicationUploadService{
		Db:                    k.DbHandler,
		AnomalityService:      k.InjectAnomalityService(),
		ApplicationRepository: &repositories.ApplicationRepository{},
	}
}

func (k *Kernel) InjectApplicationViewService() *services.ApplicationViewService {
	return &services.ApplicationViewService{
		Db:                    k.DbHandler,
		ApplicationRepository: &repositories.ApplicationRepository{},
	}
}

func (k *Kernel) InjectAnomalityService() *services.AnomalityService {
	return services.NewAnomalityService(k.InjectApplicationViewService(), k.DefectIdsDuration, k.DefectIdsDeviation)
}
