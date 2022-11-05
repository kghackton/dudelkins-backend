package bo

import (
	"fmt"
	"time"

	"github.com/uptrace/bun"

	"dudelkins/pkg/bunutils"
)

type ApplicationRetrieveOpts struct {
	CreatedFrom *time.Time
	CreatedTo   *time.Time

	ClosedFrom *time.Time
	ClosedTo   *time.Time

	IsAbnormal     *bool
	AnomalyClasses []string

	CategoryIds []int
	DefectIds   []int

	Region   *string
	District *string
	UNOM     *int64

	Entrance *string
	Floor    *string
	Flat     *string

	Limit  *int
	Offset *int
}

func (a ApplicationRetrieveOpts) QueryBuilderFuncs() (funcs []bunutils.QueryBuilderFunc) {
	if a.CreatedFrom != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("created_at >= ?", a.CreatedFrom)
		})
	}
	if a.CreatedTo != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("created_at <= ?", a.CreatedTo)
		})
	}

	if a.ClosedFrom != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("closed_at >= ?", a.ClosedFrom)
		})
	}
	if a.ClosedTo != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("closed_at <= ?", a.ClosedTo)
		})
	}

	if a.IsAbnormal != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("is_abnormal = ?", a.IsAbnormal)
		})
	}
	if len(a.AnomalyClasses) > 0 {
		anomalyClassFilter := "anomaly_classes @? '$ ? ("
		for idx, anomalyClass := range a.AnomalyClasses {
			if idx != 0 {
				anomalyClassFilter += "||"
			}
			anomalyClassFilter += fmt.Sprintf(`exists (@."%s")`, anomalyClass)
		}
		anomalyClassFilter += `)'`

		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where(anomalyClassFilter)
		})
	}

	if len(a.CategoryIds) > 0 {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("category_id IN (?)", bun.In(a.CategoryIds))
		})
	}
	if len(a.DefectIds) > 0 {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("defect_id IN (?)", bun.In(a.DefectIds))
		})
	}

	if a.Region != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("region = ?", a.Region)
		})
	}
	if a.District != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("district = ?", a.District)
		})
	}
	if a.UNOM != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("a.unom = ?", a.UNOM)
		})
	}
	if a.Entrance != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("entrance = ?", a.Entrance)
		})
	}
	if a.Floor != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("floor = ?", a.Floor)
		})
	}
	if a.Flat != nil {
		funcs = append(funcs, func(q bun.QueryBuilder) bun.QueryBuilder {
			return q.Where("flat = ?", a.Flat)
		})
	}

	return
}

func (a ApplicationRetrieveOpts) SelectOpts() (opts []bunutils.SelectOption) {
	if a.Limit != nil {
		opts = append(opts, bunutils.WithLimit(*a.Limit))
	}
	if a.Offset != nil {
		opts = append(opts, bunutils.WithOffset(*a.Offset))
	}

	return
}
