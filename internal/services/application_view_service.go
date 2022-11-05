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

func (s *ApplicationViewService) CountAnomalyClasses(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (anomalyClassCountersMap bo.AnomalyClassCountersMap, err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return anomalyClassCountersMap, errors.Wrap(err, "CountAnomalyClasses")
	}
	defer conn.Close()

	anomalyClassCounters, err := s.ApplicationRepository.CountAnomalyClasses(ctx, conn, opts.QueryBuilderFuncs())
	if err != nil {
		return anomalyClassCountersMap, errors.Wrap(err, "CountAnomalyClasses")
	}

	return anomalyClassCounters.ToBo().ToMap(), errors.Wrap(err, "CountAnomalyClasses")
}

func (s *ApplicationViewService) CountNormalAbnormal(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (normalAbnormalCountersMap bo.NormalAbnormalCountersMap, err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return normalAbnormalCountersMap, errors.Wrap(err, "CountNormalAbnormal")
	}
	defer conn.Close()

	normalAbnormalCounters, err := s.ApplicationRepository.CountNormalAbnormal(ctx, conn, opts.QueryBuilderFuncs())
	if err != nil {
		return normalAbnormalCountersMap, errors.Wrap(err, "CountNormalAbnormal")
	}

	return normalAbnormalCounters.ToBo().ToMap(), errors.Wrap(err, "CountNormalAbnormal")
}
