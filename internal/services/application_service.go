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

	ApplicationRepository interfaces.IApplicationRepository
}

func (s *ApplicationService) Create(ctx context.Context, application bo.Application) (err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return errors.Wrap(err, "Create")
	}
	defer conn.Close()

	err = s.ApplicationRepository.Insert(ctx, conn, dao.NewApplication(application))

	return errors.Wrap(err, "Create")
}

func (s *ApplicationService) Get(ctx context.Context) (applications bo.Applications, err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return applications, errors.Wrap(err, "Get")
	}
	defer conn.Close()

	applicationsDao, err := s.ApplicationRepository.Select(ctx, conn)
	if err != nil {
		return applications, errors.Wrap(err, "Get")
	}

	return applicationsDao.ToBo(), nil
}
