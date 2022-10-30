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
}
