package services

import (
	"context"

	"github.com/pkg/errors"

	"dudelkins/internal/interfaces"
	"dudelkins/internal/objects/bo"
	"dudelkins/internal/objects/dao"
)

type ApplicationUploadService struct {
	Db interfaces.IDBHandler

	AnomalityService      interfaces.IAnomalityService
	ApplicationRepository interfaces.IApplicationRepository
}

func (s *ApplicationUploadService) Create(ctx context.Context, application bo.Application) (err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return errors.Wrap(err, "Create")
	}
	defer conn.Close()

	application.AnomalyClasses, err = s.AnomalityService.CheckForAnomalies(application)
	if err != nil {
		return errors.Wrap(err, "Create")
	}

	var isAbnormal bool
	for _, anomalyClass := range application.AnomalyClasses {
		if anomalyClass.Verdict == true {
			isAbnormal = true
			break
		}
	}
	application.IsAbnormal = &isAbnormal

	applicationDao, err := dao.NewApplication(application)
	if err != nil {
		return errors.Wrap(err, "Create")
	}

	err = s.ApplicationRepository.Insert(ctx, conn, applicationDao)

	return errors.Wrap(err, "Create")
}
