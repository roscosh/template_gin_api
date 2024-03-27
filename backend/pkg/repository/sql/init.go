package sql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"template_gin_api/misc"
)

var logger = misc.GetLogger()

type SQL struct {
	Users *UsersSQL
}

func NewSQL(pool *DbPool) *SQL {
	return &SQL{
		Users: NewUsersSQL(pool, usersTable),
	}
}

type DbPool struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func (d *DbPool) Close() {
	d.pool.Close()
}

func NewDbPool(dsn string) (*DbPool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &DbPool{pool: pool, ctx: ctx}, nil
}
