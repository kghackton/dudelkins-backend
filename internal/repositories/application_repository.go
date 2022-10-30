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
