package interfaces

import (
	"context"

	"dudelkins/internal/objects/bo"
)

type IApplicationService interface {
	Create(ctx context.Context, application bo.Application) (err error)
}
