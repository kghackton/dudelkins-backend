package interfaces

import (
	"context"

	"github.com/uptrace/bun"

	"dudelkins/internal/objects/dao"
)

type IApplicationRepository interface {
	Insert(ctx context.Context, bun bun.IDB, application dao.Application) (err error)
}
