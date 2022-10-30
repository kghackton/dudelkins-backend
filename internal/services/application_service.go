package services

import (
	"context"

	"github.com/pkg/errors"

	"dudelkins/internal/interfaces"
	"dudelkins/internal/objects/bo"
	"dudelkins/internal/objects/dao"
)

type ApplicationService struct {
	Db interfaces.IDBHandler

	ApplicationRepository interfaces.IApplicationRepository
}

func (s *ApplicationService) Create(ctx context.Context, application bo.Application) (err error) {
	conn, err := s.Db.AcquireConn(ctx)
	if err != nil {
		return errors.Wrap(err, "Create")
	}
	defer conn.Close()

	err = s.ApplicationRepository.Insert(ctx, conn, dao.NewApplication(application))

	return errors.Wrap(err, "Create")
}
