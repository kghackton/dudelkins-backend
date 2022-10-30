package infrastructure

import (
	"dudelkins/internal/repositories"
	"dudelkins/internal/services"
)

func (k *Kernel) InjectApplicationService() *services.ApplicationService {
	return &services.ApplicationService{
		Db:                    k.DbHandler,
		ApplicationRepository: &repositories.ApplicationRepository{},
	}
}
