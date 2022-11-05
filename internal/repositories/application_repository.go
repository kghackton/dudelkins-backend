package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"dudelkins/internal/objects/dao"
	"dudelkins/pkg/bunutils"
)

type ApplicationRepository struct{}

func (r *ApplicationRepository) Insert(ctx context.Context, bun bun.IDB, application dao.Application) (err error) {
	_, err = bun.NewInsert().
		Model(&application).
		On("CONFLICT (root_id) DO NOTHING").
		Exec(ctx)
	return errors.Wrap(err, "Insert")
}

func (r *ApplicationRepository) Select(ctx context.Context, bun bun.IDB, queryOpts []bunutils.QueryBuilderFunc, selectOpts []bunutils.SelectOption) (applications dao.Applications, err error) {
	selectQuery := bun.NewSelect().Model(&applications)

	for _, builderFunc := range queryOpts {
		selectQuery.ApplyQueryBuilder(builderFunc)
	}
	for _, opt := range selectOpts {
		opt(selectQuery)
	}

	err = selectQuery.Scan(ctx, &applications)

	return applications, errors.Wrap(err, "Select")
}

func (r *ApplicationRepository) SelectWithUnomCoordinates(ctx context.Context, bun bun.IDB, queryOpts []bunutils.QueryBuilderFunc, selectOpts []bunutils.SelectOption) (applications dao.Applications, err error) {
	selectQuery := bun.NewSelect().Model(&applications).
		Relation("UnomCoordinate")

	for _, builderFunc := range queryOpts {
		selectQuery.ApplyQueryBuilder(builderFunc)
	}
	for _, opt := range selectOpts {
		opt(selectQuery)
	}

	err = selectQuery.Scan(ctx, &applications)

	return applications, errors.Wrap(err, "Select")
}

func (r *ApplicationRepository) CountAnomalyClasses(ctx context.Context, bunC bun.IDB, queryOpts []bunutils.QueryBuilderFunc) (anomalyClassCounters dao.AnomalyClassCounters, err error) {
	cte := bunC.NewSelect().Table("applications").
		Column("region", "district", "management_company_title").
		ColumnExpr("jsonb_object_keys(anomaly_classes) as anomaly_class")
	for _, builderFunc := range queryOpts {
		cte.ApplyQueryBuilder(builderFunc)
	}

	err = bunC.NewSelect().
		With("grouped_anomaly_class", cte).
		Table("grouped_anomaly_class").
		Column("region", "district", "management_company_title", "anomaly_class").
		ColumnExpr("count(anomaly_class) as counter").
		Group("region", "district", "management_company_title", "anomaly_class").
		Scan(ctx, &anomalyClassCounters)

	return anomalyClassCounters, errors.Wrap(err, "CountAnomalyClasses")
}
