package infrastructure

import (
	"context"
	"database/sql"
	"runtime"

	"dudelkins/internal/environment"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type PostgresDatabaseClient struct {
	*bun.DB
}

func initPostgresClient(env environment.Postgres) (*PostgresDatabaseClient, error) {
	dsn := env.FormConnStringPg()
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	maxOpenConns := runtime.NumCPU() * 8
	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetMaxIdleConns(maxOpenConns)

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(env.DebugQuery)))

	return &PostgresDatabaseClient{db}, nil
}

func (p *PostgresDatabaseClient) AcquireConn(ctx context.Context) (conn bun.Conn, err error) {
	return p.DB.Conn(ctx)
}

func (p *PostgresDatabaseClient) StartTransaction(ctx context.Context) (tx bun.Tx, err error) {
	return p.DB.BeginTx(ctx, &sql.TxOptions{})
}

func (p *PostgresDatabaseClient) FinishTransaction(tx bun.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Wrap(err, rollbackErr.Error())
		}

		return err
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return commitErr
	}

	return nil
}

func (p *PostgresDatabaseClient) Close() (err error) {
	return p.DB.Close()
}
