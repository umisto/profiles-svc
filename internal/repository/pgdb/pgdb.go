package pgdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/netbill/pgxtx"
	"github.com/netbill/profiles-svc/internal/repository"
)

type transaction struct {
	pool *pgxpool.Pool
}

func NewTransaction(pool *pgxpool.Pool) repository.Transactioner {
	return &transaction{
		pool: pool,
	}
}
func (q *transaction) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return pgxtx.Transaction(ctx, q.pool, fn)
}
