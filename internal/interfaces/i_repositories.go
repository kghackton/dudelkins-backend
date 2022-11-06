package interfaces

import (
	"context"

	"github.com/uptrace/bun"

	"dudelkins/internal/objects/dao"
	"dudelkins/pkg/bunutils"
)

type IApplicationRepository interface {
	Insert(ctx context.Context, bun bun.IDB, application dao.Application) (err error)
	Select(ctx context.Context, bun bun.IDB, queryOpts []bunutils.QueryBuilderFunc, selectOpts []bunutils.SelectOption) (applications dao.Applications, err error)
	SelectSingleWithUnomCoordinates(ctx context.Context, bun bun.IDB, id int) (application dao.Application, err error)
	SelectWithUnomCoordinates(ctx context.Context, bun bun.IDB, queryOpts []bunutils.QueryBuilderFunc, selectOpts []bunutils.SelectOption) (applications dao.Applications, err error)
	CountAnomalyClasses(ctx context.Context, bunC bun.IDB, queryOpts []bunutils.QueryBuilderFunc) (anomalyClassCounters dao.AnomalyClassCounters, err error)
	CountAnomalyClassesByCreationHour(ctx context.Context, bunC bun.IDB, queryOpts []bunutils.QueryBuilderFunc) (anomalyClassCounters dao.AnomalyClassCountersWithCreationHour, err error)
	CountNormalAbnormal(ctx context.Context, bunC bun.IDB, queryOpts []bunutils.QueryBuilderFunc, queryOptsForNormap []bunutils.QueryBuilderFunc) (normalAbnormalCounters dao.NormalAbnormalCounters, err error)
}
