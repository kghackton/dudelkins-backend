package infrastructure

import (
	"net/http"
	"time"

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
	return services.NewAnomalityService(k.InjectApplicationViewService(), k.DefectIdsDuration, k.DefectIdsDeviation, k.InjectInsService())
}

func (k *Kernel) InjectInsService() *services.InsService {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &services.InsService{
		HttpClient:  httpClient,
		InsEndPoint: k.env.InsEndpoint,
	}
}
