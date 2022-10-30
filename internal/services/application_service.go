package services

import (
	"context"

	"github.com/pkg/errors"

	"dudelkins/internal/interfaces"
	"dudelkins/internal/objects/bo"
	"dudelkins/internal/objects/dao"
)

type ApplicationService struct {
	Db interfaces.IDBHandler

	AnomalityService      interfaces.IAnomalityService
	ApplicationRepository interfaces.IApplicationRepository
}

func (s *ApplicationService) Create(ctx context.Context, application bo.Application) (err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return errors.Wrap(err, "Create")
	}
	defer conn.Close()

	application.AnomalyClasses = s.AnomalityService.CheckForAnomalies(application)
	applicationDao, err := dao.NewApplication(application)
	if err != nil {
		return errors.Wrap(err, "Create")
	}

	err = s.ApplicationRepository.Insert(ctx, conn, applicationDao)

	return errors.Wrap(err, "Create")
}

func (s *ApplicationService) Get(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (applications bo.Applications, err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return applications, errors.Wrap(err, "Get")
	}
	defer conn.Close()

	applicationsDao, err := s.ApplicationRepository.Select(ctx, conn, opts.QueryBuilderFuncs(), opts.SelectOpts())
	if err != nil {
		return applications, errors.Wrap(err, "Get")
	}
	applications, err = applicationsDao.ToBo()
	if err != nil {
		return applications, errors.Wrap(err, "Get")
	}

	return
}
