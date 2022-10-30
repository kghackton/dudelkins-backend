package interfaces

import (
	"context"

	"github.com/uptrace/bun"
)

type IDBHandler interface {
	AcquireConn(ctx context.Context) (conn bun.Conn, err error)
	StartTransaction(ctx context.Context) (tx bun.Tx, err error)
	FinishTransaction(tx bun.Tx, err error) error
	Close() (err error)
}
