package interfaces

import (
	"context"

	"dudelkins/internal/objects/bo"
)

type IApplicationService interface {
	Create(ctx context.Context, application bo.Application) (err error)
	Get(ctx context.Context, opts *bo.ApplicationRetrieveOpts) (applications bo.Applications, err error)
}

type IAnomalityService interface {
	CheckForAnomalies(application bo.Application) (anomalies map[string]bo.AnomalyClass)
}
