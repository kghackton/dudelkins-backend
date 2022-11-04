package services

import (
	"context"

	"github.com/pkg/errors"

	"dudelkins/internal/interfaces"
	"dudelkins/internal/objects/bo"
)

type ApplicationViewService struct {
	Db interfaces.IDBHandler

	ApplicationRepository interfaces.IApplicationRepository
}

func (s *ApplicationViewService) Get(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (applications bo.Applications, err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return applications, errors.Wrap(err, "Get")
	}
	defer conn.Close()

	applicationsDao, err := s.ApplicationRepository.SelectWithUnomCoordinates(ctx, conn, opts.QueryBuilderFuncs(), opts.SelectOpts())
	if err != nil {
		return applications, errors.Wrap(err, "Get")
	}
	applications, err = applicationsDao.ToBo()
	if err != nil {
		return applications, errors.Wrap(err, "Get")
	}

	return
}
