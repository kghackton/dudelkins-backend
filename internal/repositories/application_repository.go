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

func (r *ApplicationRepository) CountNormalAbnormal(ctx context.Context, bunC bun.IDB, queryOpts []bunutils.QueryBuilderFunc) (normalAbnormalCounters dao.NormalAbnormalCounters, err error) {
	cte1 := bunC.NewSelect().Table("applications").
		Column("region", "district").
		ColumnExpr("count(*) as abnormal_counter").
		Where("is_abnormal = ?", true).
		Group("region", "district")
	cte2 := bunC.NewSelect().Table("applications").
		Column("region", "district").
		ColumnExpr("count(*) as normal_counter").
		Where("is_abnormal = ?", false).
		Group("region", "district")
	for _, builderFunc := range queryOpts {
		cte1.ApplyQueryBuilder(builderFunc)
		cte2.ApplyQueryBuilder(builderFunc)
	}

	err = bunC.NewSelect().
		With("abnormal", cte1).
		With("normal", cte2).
		TableExpr("abnormal as abn").
		Column("n.region", "n.district", "abn.abnormal_counter", "n.normal_counter").
		Join("JOIN normal as n").JoinOn("n.region = abn.region AND n.district = abn.district").
		Scan(ctx, &normalAbnormalCounters)

	return normalAbnormalCounters, errors.Wrap(err, "CountNormalAbnormal")
}

/*
WITH grouped_anomaly_class AS (
    SELECT region, district, management_company_title, jsonb_object_keys(anomaly_classes) as anomaly_class, date_trunc('hour', created_at) creation_hour FROM applications
)
SELECT creation_hour, region, district, management_company_title, anomaly_class, count(anomaly_class) FROM grouped_anomaly_class
GROUP BY creation_hour, region, district, management_company_title, anomaly_class;
*/

func (r *ApplicationRepository) CountAnomalyClassesByCreationHour(ctx context.Context, bunC bun.IDB, queryOpts []bunutils.QueryBuilderFunc) (anomalyClassCounters dao.AnomalyClassCountersWithCreationHour, err error) {
	cte := bunC.NewSelect().Table("applications").
		Column("region", "district", "management_company_title").
		ColumnExpr("jsonb_object_keys(anomaly_classes) as anomaly_class").
		ColumnExpr("date_trunc('hour', created_at) as creation_hour")
	for _, builderFunc := range queryOpts {
		cte.ApplyQueryBuilder(builderFunc)
	}

	err = bunC.NewSelect().
		With("grouped_anomaly_class", cte).
		Table("grouped_anomaly_class").
		Column("creation_hour", "region", "district", "management_company_title", "anomaly_class").
		ColumnExpr("count(anomaly_class) as counter").
		Group("creation_hour", "region", "district", "management_company_title", "anomaly_class").
		Scan(ctx, &anomalyClassCounters)

	return anomalyClassCounters, errors.Wrap(err, "CountAnomalyClassesByCreationHour")
}
