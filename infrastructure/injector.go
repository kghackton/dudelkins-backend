package infrastructure

import (
	"context"
	"dudelkins/internal/environment"

	"github.com/pkg/errors"
)

type IInjector interface {
}

type Kernel struct {
	env environment.Environment

	DbHandler *PostgresDatabaseClient
}

func Inject(ctx context.Context, env environment.Environment) (k *Kernel, err error) {
	k = &Kernel{}
	k.env = env

	postgresDatabaseClient, err := initPostgresClient(env.Postgres)
	if err != nil {
		return nil, errors.Wrap(err, "Inject")
	}
	k.DbHandler = postgresDatabaseClient

	return
}
