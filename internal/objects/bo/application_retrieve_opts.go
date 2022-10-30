package bo

import (
	"time"

	"github.com/uptrace/bun"

	"dudelkins/pkg/bunutils"
)

type ApplicationRetrieveOpts struct {
	ClosedFrom *time.Time
	ClosedTo   *time.Time

	IsAbnormal  *bool
	CategoryIds []int
	DefectIds   []int

	Region   *string
	District *string
	UNOM     *int

	Limit  *int
	Offset *int
}

func (a ApplicationRetrieveOpts) QueryBuilderFuncs() (funcs []bunutils.QueryBuilderFunc) {
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
			return q.Where("unom = ?", a.UNOM)
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
