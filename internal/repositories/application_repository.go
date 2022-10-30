package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"dudelkins/internal/objects/dao"
)

type ApplicationRepository struct{}

func (r *ApplicationRepository) Insert(ctx context.Context, bun bun.IDB, application dao.Application) (err error) {
	_, err = bun.NewInsert().
		Model(&application).
		Exec(ctx)
	return errors.Wrap(err, "Insert")
}

func (r *ApplicationRepository) Select(ctx context.Context, bun bun.IDB) (applications dao.Applications, err error) {
	err = bun.NewSelect().
		Model(&applications).
		Scan(ctx, &applications)
	return applications, errors.Wrap(err, "Select")
}
