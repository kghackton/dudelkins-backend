package infrastructure

import (
	"dudelkins/internal/repositories"
	"dudelkins/internal/services"
)

func (k *Kernel) InjectApplicationService() *services.ApplicationService {
	return &services.ApplicationService{
		Db:                    k.DbHandler,
		AnomalityService:      k.InjectAnomalityService(),
		ApplicationRepository: &repositories.ApplicationRepository{},
	}
}

func (k *Kernel) InjectAnomalityService() *services.AnomalityService {
	return services.NewAnomalityService()
}
