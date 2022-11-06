package interfaces

import (
	"context"

	"dudelkins/internal/objects/bo"
)

type IApplicationUploadService interface {
	Create(ctx context.Context, application bo.Application) (err error)
}

type IApplicationViewService interface {
	Get(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (applications bo.Applications, err error)
	GetSingle(ctx context.Context, id int) (application bo.Application, err error)
	CountAnomalyClasses(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (anomalyClassCountersMap bo.AnomalyClassCountersMap, err error)
	CountAnomalyClassesWithCreationHour(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (anomalyClassCountersWithCreationHourMap bo.AnomalyClassCountersWithCreationHourMap, err error)
	CountNormalAbnormal(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (normalAbnormalCountersMap bo.NormalAbnormalCountersMap, err error)
}

type IAnomalityService interface {
	CheckForAnomalies(application bo.Application) (anomalies map[string]bo.AnomalyClass, err error)
}

type IInsService interface {
	IsAbnormal(application bo.Application) (isAbnormal bool, confidence float64, err error)
}
