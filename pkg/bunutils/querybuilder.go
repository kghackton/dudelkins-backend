package bunutils

import (
	"fmt"

	"github.com/uptrace/bun"
)

type QueryBuilderFunc func(q bun.QueryBuilder) bun.QueryBuilder // used for Where Filters

type SelectOption func(*bun.SelectQuery)

func WithLimit(limit int) SelectOption {
	return func(query *bun.SelectQuery) {
		query.Limit(limit)
	}
}

func WithOffset(offset int) SelectOption {
	return func(query *bun.SelectQuery) {
		query.Offset(offset)
	}
}

func WithOrder(column, order string) SelectOption {
	return func(query *bun.SelectQuery) {
		query.OrderExpr(fmt.Sprintf("%s %s", column, order))
	}
}
